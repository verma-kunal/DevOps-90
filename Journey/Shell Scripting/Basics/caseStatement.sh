#! /bin/bash

# Case Statements (Switch):

read -p "Are you 21 yrs or older (Y/N) " ANSWER

case "$ANSWER" in
    [yY] | [yY][eE][sS])
        echo "You are allowed to vote!"
        ;;
    [nN] | [nN][oO])
        echo "Voting NOT allowed"
        ;;
    *)
        echo "Invalid Input"
esac

