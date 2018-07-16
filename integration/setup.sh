#!/bin/bash

set -e

##define inputs
STACK=$1
TEST_STACK=$2
API_STACK=$3
COLLECTION=$4

##define variables
MACOS="Mac"

NEWURL=`aws cloudformation describe-stacks \
            --stack-name $STACK \
            --query "Stacks[0].Outputs[0].{OutputValue:OutputValue}" \
            --output text`
TESTURL=`aws cloudformation describe-stacks \
            --stack-name $TEST_STACK \
            --query "Stacks[0].Outputs[0].{OutputValue:OutputValue}" \
            --output text`
APIURL=`aws cloudformation describe-stacks \
            --stack-name $API_STACK \
            --query "Stacks[0].Outputs[0].{OutputValue:OutputValue}" \
            --output text`
echo $NEWURL
echo $TESTURL
echo $APIURL

##Check OS
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac
if [ "$machine" == "$MACOS" ]
then
    echo "You are on device with MAC OS, running sed with '' "
    sed  -i '' "s#{{URL}}#$NEWURL#g" ${COLLECTION}
    sed  -i '' "s#{{DI_API_URL}}#$APIURL#g" ${COLLECTION}
    sed  -i '' "s#{{TEST_URL}}#$TESTURL#g" ${COLLECTION}
else
    echo "You are not on a mac, running sed normally"
    sed  -i "s#{{URL}}#$NEWURL#g" ${COLLECTION}
    sed  -i "s#{{DI_API_URL}}#$APIURL#g" ${COLLECTION}
    sed  -i "s#{{TEST_URL}}#$TESTURL#g" ${COLLECTION}
fi