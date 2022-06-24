# SSH - Secure Shell

---

- What is SSH?
    - Full form → secure shell
    - Its works as a communication protocol & helps to connect with a remote computer
    - We can do anything on remote computer
        - File changes
        - Data changes
        - Setting up web servers
    - Main feature → Traffic is encrypted (thats why called secure shell)
    - Used mostly in the terminal/command line
- Client/Server Communication
    - `SSH` → client
    - `SSHD` → server (OpenSSH Daemon)
        - listens for SSH connections
        - servers have `SSHD` config files
    - Servers have  `SSHD` installed & running → required to make a connection using `SSH`
- Authentication Methods
    - Authenticating servers using `ssh`
    - Ways:
        1. **Password**
            - default method
            - creating a user on a remote server & logging in using ssh & the password for that use
            - Example:
                
                ```bash
                ssh kunal@192.168.1.29
                ```
                
        2. **Public/Private Key Pair**
            - recommended way to logging in to the remote server
            - we can generate public & private keys
            - this is a safer approach as compared to having passwords
        3. **Host Based**
            - Pre-defining the hosts that are allowed to connect to a machine
            - Defined in a file called as **`known_hosts`**
                - client file  containing all remotely connected known hosts
                , and the ssh client uses this file. This file authenticates for the client to the server they are connecting to. The known_hosts file contains the host public key for all known hosts.
- Generating `SSH` keys
    - Command used:
        
        ```bash
        ssh-keygen
        ```
        
        - `~/.ssh/id_rsa` (Private Key)
        - `~/.ssh/id_rsa.pub` (Public Key)
    - Public key goes into the server → in `authorized_keys` file
- For Windows users
    - Windows 10 supports native SSH
    - Git Bash & other terminal programs include `ssh` command & other UNIX tools

---

## Hands-On & Cheat Sheet

### ****Login via SSH with password (LOCAL SERVER)****

- **Logging in to the local server**
    - `ssh kunal@192.168.1.29`
        - `kunal` → name of the user that we wanna login as (already created in the local server)
        - `192.168.1.29` → IP address of the local server
- **Generating SSH Keys**
    - `ssh-keygen`
- **Copying the public key to the server**
    - Using the command:
        
        ```bash
        ssh-copy-id username@remote_host
        ```
        
        - The `ssh-copy-id` tool is included by default in many operating systems, so you may have it available on your local system.
    - Alternative method:
        
        ```bash
        cat ~/.ssh/id_rsa.pub | ssh username@remote_host "mkdir -p ~/.ssh && touch ~/.ssh/authorized_keys && chmod -R go= ~/.ssh && cat >> ~/.ssh/authorized_keys"
        ```
        

---

### Using SSH to log in to a remote server (created using service providers)

- Various cloud-service providers:
    - Digital Ocean
    - Linode
    - AWS
    - Hostwinds
- **Using `Linode` today!**
    - Generating new keys:
        
        ```bash
        ssh-keygen -t rsa
        ```
        
        - You can give custom names to each pair of key files you generate
            - By default, the name would be `id_rsa` & `id_rsa.pub`
        - In our case, lets name them with `id_rsa_lo`
        
        ![](https://i.imgur.com/jTztrzx.png)
        
        - Therefore, the 2 keys generated:
            - `id_rsa_lo`
            - `id_rsa_lo.pub`
    - Adding the public SSH key to Linode:
        
        ![](https://i.imgur.com/GKRuBly.png)
        
    - **Remote server is created successfully 🎉**
        
        ![](https://i.imgur.com/XhvIpK5.png)
        
    - Its time to log in to our server using ssh
        - Using the IP address
        
        ```bash
        ssh root@170.187.250.90
        ```
        
        - We are in our `Ubuntu` Remote Server
            
            ![](https://i.imgur.com/yOUtYOr.png)
            
    
    ---
    
    ### Using Ubuntu in a remote server
    
    - Updating & Upgrading all the packages is a good practice once you create your server
        
        ```bash
        sudo apt update
        ```
        
        ```bash
        sudo apt upgrade
        ```
        
    - Creating a `new user` → again, always a good practice to not use the `root` user for working
        
        ```bash
        adduser <name>
        ```
        
        ![](https://i.imgur.com/QBmAFAZ.png)
        
    - To check the info of the new user → Example: `kunal`
        
        ```bash
        id kunal
        
        # Ouput:
        uid=1000(kunal) gid=1000(kunal) groups=1000(kunal)
        ```
        
    - Adding sudo priviledges to a new user
        
        ```bash
        usermod -aG sudo kunal
        ```
        
        - Checking the above:
            
            ```bash
            id kunal
            
            # Output:
            uid=1000(kunal) gid=1000(kunal) groups=1000(kunal),27(sudo)
            ```
            
            - `27(sudo)` → `sudo` permissions have been added
    - Logging in to the server (again) as a new user (kunal)
        
        ```bash
        ssh kunal@170.187.250.90
        ```
        
    - Disabling the root user login now:
        1. Heading over the config file location
        
        ```bash
        sudo nano /etc/ssh/sshd_config
        ```
        
        1. Setting the value of `PermitRootLogin` as → `no`
        2. Reloading SSHD service
            
            ```bash
            sudo systemctl reload sshd
            ```
            
    
    ---
    
    ### Creating an Apache server on the remote server
    
    - Installing Apache
        
        ```bash
        apt install apache2 -y
        ```
        
        - `-y` → yes to everything (assumes the answers to all the prompts as YES)
    - We can now open up our IP address on the browser & will have Apache server running
        
        ![](https://i.imgur.com/uJJBxRL.png)
        

---

### Using Git & GitHub on the remote Ubuntu server

- Setting up SSH keys for GitHub
    
    ```bash
    ssh-keygen -t rsa 
    ```
    
    - We can name these as:
        - `id_rsa_gh`
        - `id_rsa_gh.pub`
- Adding our public key to GitHub (in our host server)
    
    ![](https://i.imgur.com/1d9fpYP.png)
    
- Using `ssh-add` command to add **SSH private keys** into the SSH authentication agent
    - To add the ssh-agent, for using ssh-add
        
        ```bash
        eval `ssh-agent -s`
        ```
        
    - Using ssh-add:
        
        ```bash
        ssh-add /home/kunal/.ssh/id_rsa_gh
        ```
        
    - We now have our GitHub identity added
- We can now clone the repository (on which we have added the ssh public key)
    - Example:
        
        ```bash
        git clone git@github.com:verma-kunal/PingIO
        ```
        
        ![](https://i.imgur.com/QShuP6F.png)
        

---

- Installing `Node Js` on Ubuntu:
    - Used the documentation [here](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-20-04)

---