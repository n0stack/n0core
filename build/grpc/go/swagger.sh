#!/bin/bash
set -x

dirs=`find /src -type d | grep -v .git | grep -v test`
echo > n0stack.swagger.json

for d in $dirs
do
  # 複数のファイルを指定できない
  ls -1 $d/*.proto > /dev/null 2>&1
  if [ "$?" = "0" ]; then
    protoc \
      -I/src \
      -I/tmp/include \
      -I${GOPATH}/src \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=n0stack.swagger.json:$* \
      $d/*.proto
  fi
done