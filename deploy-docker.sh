#!/bin/bash

# This script deploys a Golang Docker image to a VPS
# Exit on error
set -e

# Configuration variables - replace these with your specific values
VPS_USER="cesc"               # Your VPS username
VPS_IP="10.6.1.83"      # Your VPS IP address
VPS_PORT="322"                     # SSH port on VPS
SSH_KEY="~/.ssh/telegram-bot_rsa" # Path to your SSH private key
IMAGE_NAME="english-with-me-bot"  # Your Docker image name
IMAGE_TAG="latest"                # Your Docker image tag
REMOTE_DIR="/home/$VPS_USER"      # Directory on VPS to deploy to

# Docker run options - customize as needed
DOCKER_RUN_OPTS="-d --name $IMAGE_NAME --restart unless-stopped"

# Function to clean up temporary files
cleanup() {
    echo "Cleaning up temporary files..."
    rm -f deploy.sh "$IMAGE_NAME.tar.gz" "$IMAGE_NAME.tar" 2>/dev/null || true
}

# Set up trap to call cleanup function on exit due to error
trap 'cleanup' ERR

echo "Starting deployment process for $IMAGE_NAME:$IMAGE_TAG to $VPS_USER@$VPS_IP..."

# Step 1: Package the Docker image
echo "Packaging Docker image $IMAGE_NAME:$IMAGE_TAG..."
docker build -t english-with-me-bot .

if ! docker save -o "$IMAGE_NAME.tar" "$IMAGE_NAME:$IMAGE_TAG"; then
    echo "Error: Failed to package Docker image. Is Docker running? Does the image exist?"
    exit 1
fi

# Step 2: Compress the tar file
echo "Compressing Docker image..."
if ! gzip -f "$IMAGE_NAME.tar"; then
    echo "Error: Failed to compress Docker image tar file."
    exit 1
fi

# Step 3: Create a deployment script to run on the VPS
echo "Creating deployment script..."
cat >deploy.sh <<EOF
#!/bin/bash

# Exit on error
set -e

echo "Starting deployment on VPS..."

# Check if Docker is installed and running
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed on the VPS."
    exit 1
fi

if ! docker info &> /dev/null; then
    echo "Error: Docker daemon is not running on the VPS."
    exit 1
fi

# Decompress the tar file
echo "Decompressing Docker image..."
gunzip -f $IMAGE_NAME.tar.gz

# Load the Docker image
echo "Loading Docker image..."
docker load -i $IMAGE_NAME.tar

# Stop existing container if it exists
echo "Stopping and removing existing container (if any)..."
docker stop $IMAGE_NAME 2>/dev/null || true
docker rm $IMAGE_NAME 2>/dev/null || true

# Run the new container
echo "Starting new container..."
docker run $DOCKER_RUN_OPTS $IMAGE_NAME:$IMAGE_TAG prod

# Cleanup
echo "Cleaning up temporary files..."
rm $IMAGE_NAME.tar

echo "Deployment completed successfully!"
EOF

chmod +x deploy.sh

# Step 4: Transfer the packaged image and deployment script to the VPS
echo "Transferring files to VPS..."
if ! scp -P "$VPS_PORT" -i "$SSH_KEY" "$IMAGE_NAME.tar.gz" deploy.sh "$VPS_USER@$VPS_IP:$REMOTE_DIR/"; then
    echo "Error: Failed to transfer files to VPS."
    cleanup
    exit 1
fi

# Step 5: Execute the deployment script on the VPS
echo "Executing deployment script on VPS..."
if ! ssh -p "$VPS_PORT" -i "$SSH_KEY" "$VPS_USER@$VPS_IP" "cd $REMOTE_DIR && ./deploy.sh"; then
    echo "Error: Failed to execute deployment script on VPS."
    cleanup
    exit 1
fi

# Step 6: Clean up local files
cleanup

echo "Script execution completed successfully!"
