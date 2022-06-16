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