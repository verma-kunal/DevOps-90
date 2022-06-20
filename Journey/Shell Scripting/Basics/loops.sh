#! /bin/bash

# FOR Loop:
# NAMES="Kunal Jack Alice Mark"

# for NAME in $NAMES
#     do
#         echo "Hello, ${NAME}"
# done

# --------------------------------------------------------

# FOR Loop to rename files:

# FILES=$(ls *.txt) # listing all the files in this directory
# NEW="new"

# for FILE in $FILES
#     do
#         echo "Renaming $FILE to new-$FILE"
#         mv "$FILE" "$NEW-${FILE}"
# done

# --------------------------------------------------------

# WHILE Loop:

# Reading through a file (line by line):

LINE=1 # The iterator

while read -r CURRENT_LINE 
    do
        echo "$LINE: $CURRENT_LINE"
        ((LINE++))
done < "./new-1.txt"


