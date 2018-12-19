#!/bin/sh

export GO_BUILD_COUNT=$(($GO_BUILD_COUNT+1))
FILES=`find ./src -name "*.go" -print`

rm -rf ./pkg
mkdir -p ./pkg

for f in $FILES ; do
  FILE=${f##*/}
  FILE_WITHOUT_SUFFIX=${FILE%%.*}
  cp $f ./pkg/${FILE_WITHOUT_SUFFIX}-${RANDOM}.go
done
# perl -i -pe 's/_main/main/g' ./pkg/* 
echo "> \033[0;32mbuilding\033[0;0m \033[0;30m"$GO_BUILD_COUNT"\033[0;0m"
# env GOOS=linux GOARCH=arm GOARM=5 go build -o robocup pkg/*
env go build -o robocup pkg/*.go
if [[ $? != 0 ]]; then
  echo "< \033[0;31mbuild failed\033[0;0m"
else
  echo "| \033[0;32mbuild finished\033[0;0m"
  rm bin/* 2> /dev/null
  mv robocup bin/
  echo "| \033[0;32msending executable\033[0;0m"
  # scp -q bin/robocup $1:/home/robot/src/bin/
  if [[ $? != 0 ]]; then
    echo "< \033[0;31msend failed\033[0;0m"
  else
    echo "< \033[0;32mdone\033[0;0m"
  fi
fi
