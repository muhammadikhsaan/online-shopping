#!/bin/bash

tmp="./tmp"
output="tmp"
input="pensiel.com/cmd"

if [[ "$OSTYPE" == "msys" ]];
  then
    output+="\\main.exe"
    tmp+="/main.exe"
else
  output+="main"
  tmp+="/main"
fi

if [[ $ENVIRONTMENT == "PRERELEASE" ]];
  then
    export $(echo $(cat .env.prerelease | xargs)) &&
    air --build.cmd "go build -o $tmp $input" --build.bin "$output" --build.exclude_dir "migrations,bash,.vscode,logs,tmp"
elif [[ $ENVIRONTMENT == "PRODUCTION" ]];
  then
    export $(echo $(cat .env.production | xargs)) &&
    air --build.cmd "go build -o $tmp $input" --build.bin "$output" --build.exclude_dir "migrations,bash,.vscode,logs,tmp"
else
    export ENVIRONTMENT=LOCAL
    export $(echo $(cat .env.local | xargs)) &&
    air --build.cmd "go build -o $tmp $input" --build.bin "$output" --build.exclude_dir "migrations,bash,.vscode,logs,tmp"
fi