#! /bin/bash


name=$1
compliment=$2

user=$(whoami)
date=$(date)
whereami=$(pwd)

echo "Good Morning, $name"
sleep 1
echo "You're looking good today, $name"
sleep 1
echo "You have a nice $compliment face, $name"

sleep 2
echo "You are logged in as $user & in the directory: $whereami"

sleep 1
echo "BTW, today is: $date"
