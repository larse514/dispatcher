#!/bin/bash

set -e

STACK_NAME=$1
DEPLOYMENT_BUCKET=$2
API_STACK=$3

APIURL=`aws cloudformation describe-stacks \
            --stack-name $API_STACK \
            --query "Stacks[0].Outputs[0].{OutputValue:OutputValue}" \
            --output text`

echo "retrieved ${APIURL}"
make
make package-aws DEPLOYMENT_BUCKET=${DEPLOYMENT_BUCKET} SOURCE_URL=${APIURL}
make deploy STACK_NAME=${STACK_NAME} ENV=ci 