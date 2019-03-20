[ -z "$1" ] && echo "no remote supplied" && exit 1

env GOOS=linux GOARCH=arm GOARM=5 go build

if [ $? -eq 0 ]; then
  scp ./robocup-2019 $1:/home/robot/src/bin
  ssh -t $1 "/home/robot/src/bin/robocup-2019"
  rm ./robocup-2019
else
  echo "failed to build"
fi

exit 0