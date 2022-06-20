# Shell Scripting

<aside>
⚙️ Traversy Media → [https://youtu.be/v-F3YLd6oMw](https://youtu.be/v-F3YLd6oMw)

</aside>

---

- What is a Shell script?
    - program designed to run on a UNIX-based command terminal/shell
- What is Bash?
    - Dialect of a shell syntax
    - `BASH` → Born Shell
    - Its a standard syntax
    - Used as a scripting language
    - `.sh` → Bash script extension

---

- By default, shell script files don’t have executable permissions
    - To make the file an executable:
        
        ```bash
        chmod +x <name>.sh
        ```
        
- To run a shell script:

```bash
./<name.sh>
```

---

- Writing the first script
    - Declaration of a bash script:
        
        ```bash
        #! /bin/bash
        ```
        
        - `#!` → Shebang
        - `/bin/bash` → location of the bash executable
    - To use comments → `#`
    - First Script → Printing “Hello World”
    
    ```bash
    #! /bin/bash
    
    # Echo - Hello World
    echo Hello World! 
    ```
    
- Variables in a script
    - By convention → `UPPERCASE`
    - Allowed in variables:
        - Letters
        - Numbers
        - `_`
    - Creating a variable:
    
    ```bash
    NAME="Kunal"
    ```
    
    - Printing the variable value:
    
    ```bash
    echo ${NAME}
    ```
    
- User Input
    
    ```bash
    read -p "Some string: " VARIABLE
    ```
    
    - Accessing the above variable:
    
    ```bash
    $VARIABLE_NAME
    ```
    
- Conditionals in a script
    - **IF statement**
        - Syntax:
            
            ```bash
            if [ expression ]
            then
               statement
            fi
            ```
            
    - **IF-Else statement**
        - Syntax:
            
            ```bash
            if [ expression ]
            then
               statement
            else
            	statement
            fi
            ```
            
    - **IF..elif..else..fi statement**
        - Syntax:
            
            ```bash
            if [ expression1 ]
            then
               statement1
               .
               .
            elif [ expression2 ]
            then
               statement2
               .
               .
            else
               statement3
            fi
            ```
            
- Comparisons (Operators)
    - `-eq` : check for equality
    - `-ne` : check for not equals
    - -`gt` : check for greater than
    - `-ge` : check for greater than or equals
    - `-lt` : check for less than
    - `-le` : check for less than or equals
- File Conditions
    - `-d <file>` : check for a directory
    - `-f <file>` : check for file existence or if the provided string is a file
    - `-g <file>` :  check for group id is set or not
    - `-r <file>` : check for if file is readable
    - `-s <file>` : check for if file has non-zero size
    - `-u` : user id set on file (true)
    - `-w` : file is writable (true)
    - `-x` : file is executable (true)
    - Example:
    
    ```bash
    FILE="main.txt"
    
    if [ -f "$FILE" ]
    then
        echo "${FILE} is a file!"
    else
        echo "${FILE} is NOT file!"
    fi
    ```
    
- Case Statements (Switch)
    - Syntax:
        
        ```bash
        case expression in 
        	pattern1) 
        		Statement(s) to be executed if pattern1 matches 
        		;; 
        	pattern2) 
        		Statement(s) to be executed if pattern2 matches 
        		;; 
        	pattern3) 
        		Statement(s) to be executed if pattern3 matches 
        		;; 
        	*) 
        		Default condition to be executed 
        		;; 
        esac
        ```
        
    - Example:
        
        ```bash
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
        ```
        
        - Output:
            
            ![](https://i.imgur.com/Rg0YglY.png)
            
- Loops
    - **FOR Loop:**
        - Syntax:
            
            ```bash
            for i in <data>
            do
              statement
            done
            ```
            
        - Example:
            
            ```bash
            NAMES="Kunal Jack Alice Mark"
            
            for NAME in $NAMES
                do
                    echo "Hello, ${NAME}"
            done
            ```
            
        - Example (Renaming some files):
            
            ```bash
            FILES=$(ls *.txt) # listing all the files in this directory
            NEW="new"
            
            for FILE in $FILES
                do
                    echo "Renaming $FILE to new-$FILE"
                    mv "$FILE" "$NEW-${FILE}"
            done
            
            # Output:
            Renaming 1.txt to new-1.txt
            Renaming 2.txt to new-2.txt
            Renaming 3.txt to new-3.txt
            ```
            
            - To use actual commands in your bash script, we can use `$(command)`
    
    ---
    
    - **WHILE Loop**
        - Syntax:
            
            ```bash
            while [ condition ]
            do
               command1
               command2
               command3
            done
            ```
            
        - Example: (Reading a file - line by line)
            
            ```bash
            LINE=1 # The iterator
            
            while read -r CURRENT_LINE 
                do
                    echo "$LINE: $CURRENT_LINE"
                    ((LINE++))
            done < "./new-1.txt"
            
            # Output:
            
            1: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut risus ante, vulputate id mauris sed, sodales elementum magna. Nullam nec est interdum, sodales tellus vel, molestie nibh. In venenatis urna at tellus imperdiet iaculis. Nulla condimentum aliquet maximus. Morbi pulvinar urna in pharetra blandit. Etiam vestibulum ipsum et tortor venenatis, eu consequat sapien maximus. Curabitur quis ultricies lectus. Vivamus cursus vestibulum odio. Cras sagittis orci in neque finibus ornare. Fusce in nisl sed turpis vehicula laoreet vitae vitae sem. Nulla justo leo, pretium at lobortis molestie, congue eget urna. Sed aliquet metus sit amet fermentum cursus. Donec sit amet dapibus elit, eget aliquet lorem. Phasellus tempus vulputate mauris, at tempus orci faucibus ac.
            2: 
            3: Etiam convallis turpis est, nec finibus neque dictum id. Quisque rutrum cursus vehicula. Maecenas est purus, vulputate non purus ac, iaculis sollicitudin felis. Donec sed malesuada sapien. Sed semper accumsan vulputate. Integer sit amet pulvinar nulla. In hac habitasse platea dictumst. Morbi iaculis tellus mattis neque tempor, in posuere enim luctus. Aliquam erat volutpat.
            4:
            ```
            
- Functions
    - Example:
        
        ```bash
        function myFunc() {
            echo "Hello world!"
        }
        
        # Calling the function:
        myFunc
        ```
        
    - Example (with params):
        
        ```bash
        function greet() {
            echo "Hello, I am $1 & my age is $2"
        }
        
        greet "Kunal" "19"
        
        # Output:
        
        Hello, I am Kunal & my age is 19
        ```
        
        - Here we actually used positional parameters i.e. `$1` & `$2`, which replaced the values given by us, according to the position of defining them!
- Write actual commands in a script
    - Example: Create a folder & write to a file in that
        
        ```bash
        #! /bin/bash
        
        # Create a directory & write a file in it:
        
        mkdir main
        touch "main/new.txt"
        echo "Hello World!" >> "main/new.txt"
        echo "Created a file: main/new.txt"
        ```