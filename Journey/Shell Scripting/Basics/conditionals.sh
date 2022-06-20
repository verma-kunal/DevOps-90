#! /bin/bash

# read -p "Enter your name: " NAME

# Simple IF statment:
# if [ "$NAME" == "Kunal" ]
# then
#     echo "Hello, ${NAME}"
# fi

# IF-Else Statement:
# if [ "$NAME" == "Kunal" ]
# then
#     echo "Hello, ${NAME}"
# else
#     echo "Wrong name!"
# fi

# IF-Elif-Else-fi Statement:
# if [ "$NAME" == "Kunal" ]
# then
#     echo "Hello, ${NAME}"
# elif [ "$NAME" == "Jack" ]
# then
#     echo "Hello, ${NAME}"
# else
#     echo "Your name is neither Kunal nor Jack!"
# fi

# Comparisons (Using operators)
NUM1=10
NUM2=20

# if [ "$NUM1" -lt "$NUM2" ]
# then
#     echo "${NUM1} is less than ${NUM2}"
# else
#     echo "${NUM1} is not less than ${NUM2}"
# fi

# File Conditions:
FILE="main.txt"

if [ -f "$FILE" ]
then
    echo "${FILE} is a file!"
else
    echo "${FILE} is NOT file!"
fi


