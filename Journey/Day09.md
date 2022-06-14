# Linux for DevOps

Number: 1
Type: pre-requisite

- Terminal Emulator
    - Application that helps us use the terminal in a graphical manner
    - Example: iterm2, linux terminal etc.
    - Interacting with the operating system

---

- Ever command has its own executable file
    - whether we run the command or the executable file, its the same thing
    - `where <command>` → to find the location of the executable file of that command

---

- **Environment Variables**
    - named values that are used to change how the process & commands are to be executed
        - **Process** → running instance of a command is called a Process
    - **PATH** environment variable:
        - denoted by `$PATH`
        - every command that we execute needs to be in the folder, that is present in the path variable
        Example: `/usr/bin`
        
        ```bash
        /usr/local/opt/ruby/bin:/Users/kunalverma/.rbenv/shims:/Users/kunalverma/.gem/ruby/3.0.0/bin:/usr/local/opt/openssl@3/bin:/Users/kunalverma/.gem/ruby/3.0.0/bin:/usr/local/mysql/bin:/opt/local/bin:/opt/local/sbin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/Applications/VMware Fusion.app/Contents/Public:/Users/kunalverma/.fig/bin:/Users/kunalverma/.local/bin:/Applications/apache-maven-3.8.4/bin
        ```
        
    - Creating your own env variables:
    
    ```bash
    export <VAR_NAME>="value"
    ```
    

---

- **Bash Files**
    - the conifigurable files that are executed when a terminal starts (these are hidden files)
        - what theme we are using?
        - what colour we use?
        - which type of shell we are using?
            - **Bash Shell**
                - `.bashrc`
                - `.bash_profile`
            - **Zsh Shell**
                - `.zshrc`
                - `/etc/zshrc`
                - `/etc/zprofile`

---

- **Aliases**
    - shortcuts for a command
    
    ```bash
    alias <custom_name>="YOUR COMMAND HERE"
    ```
    

---

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
- `cat <file_name>` → used to see whats inside a file quickly through the terminal