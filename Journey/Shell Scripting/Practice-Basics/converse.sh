#! /bin/bash

echo "Welcome to JARVIS!"
read -p "Enter your name please: " NAME

sleep 3

echo "Hello ${NAME}! I am JARVIS, your personal assistent!"
echo
sleep 3

read -p "how r u today? " MOOD
echo
if [ $MOOD == "good" ]
    then
        sleep 3
        echo "Thats great, hope you have a nice day!"
elif [ $MOOD == "bad" ]
    then 
        sleep 3   
        echo "Well, everyone has bad days sometimes! Don't worry at all"
else
    sleep 3
    echo "Alright, no worries!"
fi

echo
sleep 3
echo "Thank you for using me, ${NAME}!"
echo
sleep 3
echo "HAVE A NICE DAY!!!"

