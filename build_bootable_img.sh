#!/bin/bash

set -e  # Exit on any error
# Configuration
NOW=$(date +'%d.%m.%y')
IMAGE_NAME="my_image_${NOW}.img"
IMAGE_SIZE_IN_MB=10
MOUNT_POINT="/mnt/my_mount_point"
LOOP_DEVICE=""

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit 1
fi

# Clean up on exit or error
cleanup() {
  echo "Cleaning up..."
  # if mountpoint -q "$MOUNT_POINT"; then
  #   umount "$MOUNT_POINT"
  # fi
  
  # if [ -n "$LOOP_DEVICE" ]; then
  #   losetup -d "$MOUNT_POINT"
  # fi
  
  echo "Done."
}

trap cleanup EXIT

echo "Creating empty disk image ($IMAGE_SIZE_IN_MB MB)..."
dd if=/dev/zero of="$IMAGE_NAME" bs=1M count="$IMAGE_SIZE_IN_MB" status=progress

# Set up loop device
echo "Setting up loop device..."

# -f: Find the first unused loop device
LOOP_DEVICE=$(losetup -f)

echo "****Found first unused loop device $LOOP_DEVICE****"

# -P: --partscan, scan the partition table on a newly created loop device. default is sector size is 512 bytes
losetup -P "$LOOP_DEVICE" "$IMAGE_NAME"

# Create partition
echo "Creating partition..."
fdisk "$LOOP_DEVICE" <<EOF
o
n
p
1


w
EOF

# echo "Creating ext2 filesystem..."
# mkfs.ext2 "${LOOP_DEVICE}p1"

# echo "Mounting filesystem..."
# mkdir -p "$MOUNT_POINT"
# mount "${LOOP_DEVICE}p1" "$MOUNT_POINT"


# Make bootable
echo "Making partition bootable..."
fdisk "$LOOP_DEVICE" <<EOF
a
w
EOF

fdisk -l

# Sync and finish
echo "Syncing files..."
sync

echo "Image creation complete: $IMAGE_NAME"
echo "You can now flash this image to a CF card with:"
echo "  dd if=$IMAGE_NAME of=/dev/hdaX bs=4M status=progress"
