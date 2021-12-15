#!bin/bash

APP_NAME=gore

if [ "$APP_NAME" = "" ]
then
  echo "Please set the APP_NAME variable in re.sh to use this script."
  exit 1
fi

GOOS=linux GOARCH=amd64 go build
mv $APP_NAME ~/release/$APP_NAME
