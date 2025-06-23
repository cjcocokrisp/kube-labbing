#!/bin/bash
# This shell script builds a container image and then pushes it to 
# the specified repo.
#
# Usage: ./push_test_image.sh {TAG} {REPO}
# TAG argument is required however the REPO argument is optional
# If no repo is provided will default to specificed repo on line 31 

TAG="$1"
REPO="$2"
CONTAINER_NAME="webhook_test_dummy"
IMAGE_NAME="webhook_test_image"
 
set -u

if test "${TAG}" = ""; then
	echo "!!  Usage: ./push_test_image.sh {TAG}"
fi

# remove dummy container so you can create a new one
# container may not have existed so seeing an error is ok
sudo podman rm ${CONTAINER_NAME}

# build, run, and kill container
sudo podman build . -t ${IMAGE_NAME}
sudo podman run -d --name ${CONTAINER_NAME} -p 8000:8000 ${IMAGE_NAME}
sudo podman kill ${CONTAINER_NAME}

# commit and push container
if test "${REPO}" = ""; then
	REPO="quay.io/rh-ee-ccoco/webhooktest" 
fi

TARGET="${REPO}:${TAG}"
sudo podman commit ${CONTAINER_NAME} ${TARGET}
sudo podman push ${TARGET}
