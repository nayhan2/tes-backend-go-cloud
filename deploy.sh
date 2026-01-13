#!/bin/bash

# Deploy script untuk Google Cloud Run
# Usage: ./deploy.sh

set -e

PROJECT_ID="nehan-416303"
SERVICE_NAME="backend-go"
REGION="asia-southeast1"
IMAGE="gcr.io/${PROJECT_ID}/${SERVICE_NAME}"

echo "🔨 Building Docker image..."
gcloud builds submit --tag ${IMAGE}

echo "🚀 Deploying to Cloud Run..."
gcloud run deploy ${SERVICE_NAME} \
  --image ${IMAGE} \
  --region ${REGION} \
  --platform managed \
  --allow-unauthenticated \
  --set-env-vars "DB_HOST=34.126.66.238,DB_PORT=5432,DB_USER=postgres,DB_PASSWORD=\$Nehangans3,DB_NAME=postgres,SWAGGER_HOST=backend-go-122276103213.asia-southeast1.run.app"

echo "✅ Deploy selesai!"
echo "🌐 URL: https://${SERVICE_NAME}-122276103213.${REGION}.run.app"
