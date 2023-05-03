#!/bin/bash

# Set your Terraform Cloud organization name via the TF_ORG_NAME environment variable
ORG_NAME="$TF_ORG_NAME"

# Set the workspace name from the command line argument
WORKSPACE_NAME="$1"

# Check if the workspace name is provided as an argument
if [[ -z "$WORKSPACE_NAME" ]]; then
  echo "Please provide the workspace name as an argument"
  exit 1
fi

# Retrieve the latest run ID for the workspace
RUN_ID=$(curl -s -H "Authorization: Bearer $TF_API_TOKEN" "https://app.terraform.io/api/v2/workspaces/$ORG_NAME/$WORKSPACE_NAME/runs?filter%5Bstatus%5D=errored&filter%5Boperations%5D=false" | jq -r '.data[0].id')

# Check if there are any errored runs
if [[ -z "$RUN_ID" ]]; then
  echo "No errored runs found for workspace $WORKSPACE_NAME"
  exit 0
fi

# Retrieve the logs for the run
LOGS_URL=$(curl -s -H "Authorization: Bearer $TF_API_TOKEN" "https://app.terraform.io/api/v2/runs/$RUN_ID" | jq -r '.data.relationships."logs".links."related"')

curl -s -H "Authorization: Bearer $TF_API_TOKEN" "$LOGS_URL" | jq -r '.data.attributes."raw-content"'

