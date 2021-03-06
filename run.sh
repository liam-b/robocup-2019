[ -z "$1" ] && echo "no remote supplied" && exit 1

env GOOS=linux GOARCH=arm GOARM=5 go build

if [ $? -eq 0 ]; then
  sleep 1
  scp ./robocup-2019 $1:/home/robot/src/bin
  sleep 1
  ssh -t $1 "/home/robot/src/bin/robocup-2019"
  sleep 1
  ssh -t $1 "/home/robot/src/bin/stop.sh"
  rm ./robocup-2019
else
  echo "failed to build"
fi

exit 0