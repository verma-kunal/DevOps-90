### List Of Commands

- `ls` → list all files & folder in a directory
- `mkdir <name>` → create a new folder
- `cd` → changing the directory/folder
- `cd ..` → to go back a folder
- `pwd` → print working directory or current directory
- `ls -a` → list all the hidden files as well
- `ls -l` → showing the file permissions as well
    - `rwxd` ⇒ read, write, execute, delete
- `ls -al` → combining the above two commands
    - hidden files + details
- Relative path:
    - `.` → the current directory
    - `..` → the previous directory
- `cd <the whole path>` → you can specify the whole path of your directory
- `ls -R` → listing all the sub-directories as well
- `cat <file_name>` → used to see whats inside a file quickly through the terminal  (concatinate)
- `cat > <filename>` → quickly create a file & add some text to it
    - `>` → **Redirection command:** redirect the output to some other file
- `cat file1 file2 > file3`  → merging the contents of file 1&2 into file 3
- `echo “something” > file.txt` → will add the text to the file mentioned
- `man <command>` → to learn more about a command (information)
- `|`  → **Pipe Command:** output for the 1st command, acts as input for the 2nd command
    - Example:
    
    ```bash
    cat file.txt | tr a-z A-Z > upper.txt
    ```
    
- `tr` → tranlate characters

---

### File System Commands

- `mkdir <name or path>` → make a new directory
- `mkdir -p dir1/middle/dir2` → creating a directory in the middle of `dir1` & `dir2`
- `touch <filename or path>` → create a new file
- `cp <filename> <new filename>` → to make a copy of a file
- `mv <filename> <location>` → to move a file to a new location
    - can also be used to rename a file
    - `mv <filename> <newfilename>`
- `mv <filename> <location + newfilename>` → to move as well as rename the file
- `rm <filename>` → to remove a file
    - Caution ⇒ removed permanently
- `cp -R dir1 dir2` → copy `dir1` to `dir2`
- `rmdir` → removes an empty directory
    - `rm -R` → to remove a directory completely along with its sub-directories
- `sudo` → super-user do (admin access command)
    - Need password for this
- `df` → check out the free disk space
    - it has flags:
        - `-h` → in human-readable format
- `du` → disk usage statistics
- `head <filename>` → by default will display the first 10 lines of a file
    - `head -n 3 <file>` → will only display the first 3 lines
- `tail <filename>` → by default displays the last 10
- `diff file1 file2` → compares both the files & would output the uncommon ones
- `locate "file extension"` → locates all the files with the same extension as mentioned
    - Example:
    
    ```bash
    locate "*.txt"
    ```
    
- `find <name>` → find things in a file or a directory
    - lot of useful flags here
    - check out `man find`

---

### File Permissions

- Three types:
    - read (r)
    - write (w)
    - execute (x)
- Three types of ppl that these permissions are set for:
    - user
    - group → an entity among a user group
    - guests or the other people
- `ls -l <filename>` → shows the file permissions for that particular file
    - Example:
    
    ```bash
    ls -l test2.txt
    
    -rw-r--r--  1 kunalverma  staff  19 Jun 16 11:44 test2.txt
    ```
    
    - `-rw` : for the user
    - `-r` : for a group
    - `--r` : for other ppl
- **Changing the file permissions**
    1. `chmod` command
        
        ```bash
        chmod u=rwx,g=rx,o=r <filename>
        ```
        
        - changes made here:
            - User: `rwx`
            - Group: `rx`
            - Other: `r`
    2. Using octal numbers with `chmod`
        - `chmod 777 <filename>`
            - `777` → full access to everyone
        - Categories:
            - `4` → read
            - `2` → write
            - `1` → execute
            - `0` → no permission
- **Changing the users of a file**
    - `chown` command
        
        ```bash
        sudo chown root <filename>
        ```
        
        - we have to use `sudo` as we are accessing `root` permissions
        - `root` user has the highest number of permissions in linux/unix based systems

---

- Performing action on multiple files
    - `-exec` command
        - does not create a new process it just replaces the bash with the command to be executed

---

- **grep** command
    - Global regular expression print
    - Allows to search some text in files
    - Case sensitive
        - Versions of grep:
        
        ```bash
        grep -V
        ```
        
        - **Mac** → BSD grep
        - **Linux** → GNU
    1. Simple search
        
        ```bash
        grep "Kunal" names.txt
        
        **Kunal**
        ```
        
        - If present, it will return the string itself!
    2. Expanding the whole word
        
        ```bash
        grep -w "Kunal" names.txt
        
        **Kunal** Verma
        ```
        
        - Using `-w` (means word-regexp), we searched for the “Kunal” & it returned the whole string associated with that word!
    3. To ignore the search case-sensitivity
        
        ```bash
        grep -i "kunal" names.txt
        
        **Kunal** Verma
        ```
        
    4. Print the line number of the word
        
        ```bash
        grep -n "Kunal" names.txt
        
        1:Kunal Verma
        ```
        
    5. Print lines after & before the word
        
        ```bash
        grep -B 5 "Rishab" names.txt
        
        Kunal Verma
        Arpit Aggrawal
        Parth Kumar
        Utsav Jhamb
        Malay Acharya
        **Rishab** Kesarwani
        ```
        
        - `-B 5` → means print 5 lines before the actual word we are searching
        
        ```bash
        grep -A 5 "Rishab" names.txt
        ```
        
        - `-A 5` → means print 5 lines after the actual word we are searching
    6. To search the entire directory for a word
        
        ```bash
        grep -win "Kunal" ./*.txt
        ```
        
        - This will check all the `.txt` files in the current directory for the word
    7. Which file contains a match
        
        ```bash
        grep -wirl "Kunal" .
        ```
        
        - using the `-l` tag here, we found the file that contained “Kunal” in the current directory
    8. Using `regex` to serch
        - Regular expression commands
        - **Come back to it later**

---

- **Useful terminal shortcuts**
    - `history` command
        - shows a history of all the commands you ran in that particular session
        - to directly use a command from history
        
        ```bash
        !<history-number>
        ```
        
    - `Cmd + K` → clear the entire terminal
    - `;` → to use multiple commands in the same line
    
    ```bash
    git add .;git commit -m "message";git push origin main
    ```
    

---

- `ping <URL>` → to check the connectivity status of a website
- `wget <URL>` → to download any file from the internet
    - you’ll have to install this using your package managers
        - `apt install`
        - `brew install`
        - `yum install`
- `top` → process running + CPU usage
    - `kill <process_id>` → to stop a running process
- `uname` → prints out your OS name
    - has many tags associated
- `zip & unzip`
- `hostname` → prints the host’s name
- `vmstat` → checking your virtual memory
- `id` → to check user & group id’s
- `getent` → if a particular user exists or not
- `lsof` → list all the open files

---

### Networking Commands

- `nslookup <URL>`
    - to checkout the ip address of a domain
- `cut`
    - very handy when handling log files
    - cuts the words/lines accordingly
    - Example:
    
    ```bash
    cut -c 1-2 company.txt
    
    Ci
    Ar
    Ch
    CN
    Ku
    ```
    

---

### Working with Operators

- Combining various commands together
1. `&` operator
2. `&&` operator
    
    ```bash
    echo "first" && echo "second"
    ```
    
    - only when the `first` command is completed, then the `second` will run
3. `;` operator
4. `||` operator (OR)
5. `|` operator (pipe)
6. `!()` operator (NOT)
7. `>` → over-write
8. `>>` → append
    
    ```bash
    echo "hey" >> names.txt
    ```
    
9. `{}` → combination operator

---