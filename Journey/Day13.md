# Web Servers - Nginx & Apache

---

## Nginx

NGINX is open source software for **web serving, reverse proxying, caching, load balancing, media streaming etc.**

- Solves a very specific problem of creating multiple instances of an application (on different ports) & thus confusing the user on which to choose!
- Nginx decides which PORT to serve the application on.
    - It manages the load in the all the ports automatically, so that we don’t have to worry about that

---

### Nginx on a remote server (Linode)

Now, let us learn how we can install & configure Nginx server in a remote **Ubuntu 22.04** server by **`Linode`**

- Nginx Installation
    - 2 ways to install Nginx:
        1. [via Ubuntu’s default repositories](https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-22-04)
        2. [Mainline installation via official Nginx packages](https://nginx.org/en/linux_packages.html?_ga=2.31476086.934049599.1656297289-952190285.1656297289#Ubuntu)
    
    ---
    
    - Installing Nginx from **Ubuntu’s default repositories** using apt
        
        ```bash
        sudo apt update
        sudo apt install nginx
        ```
        
    - To check if Nginx is running:
        
        ```bash
        systemctl status nginx
        
        #Output:
        
        ● nginx.service - A high performance web server and a reverse proxy server
             Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
             Active: active (running) since Mon 2022-06-27 02:57:08 UTC; 3min 47s ago
               Docs: man:nginx(8)
            Process: 1795 ExecStartPre=/usr/sbin/nginx -t -q -g daemon on; master_process on; (code=exited, status=0/SUCCESS)
            Process: 1796 ExecStart=/usr/sbin/nginx -g daemon on; master_process on; (code=exited, status=0/SUCCESS)
           Main PID: 1889 (nginx)
              Tasks: 2 (limit: 1034)
             Memory: 4.7M
                CPU: 19ms
             CGroup: /system.slice/nginx.service
                     ├─1889 "nginx: master process /usr/sbin/nginx -g daemon on; master_process on;"
                     └─1892 "nginx: worker process" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" "" ""
        ```
        
    - You can now go ahead open up your IP Address in a web browser  find the default Nginx Page
        - In my case, that would be [http://172.105.39.205](http://172.105.39.205/)
            
            ![](https://i.imgur.com/zSblDji.png)
            
        
- Configuration files location of Nginx → `**~/etc/nginx**`
- Overview of `nginx.conf` file
    - Default lines/contexts in every Nginx configuration (not server-specific)
        - These are config lines for Nginx itself
        
        ```bash
        user www-data;
        worker_processes auto;
        
        error_log /var/log/nginx/error.log warn;
        pid /run/nginx.pid;
        include /etc/nginx/modules-enabled/*.conf;
        ```
        
    - `www-data` → is the user that web servers on Ubuntu (Apache, nginx, for example) use by default for normal operations.
        - Two types of users can found here (online sources)
            1. user `nginx`
            2. user `www-data`
        - There is no real difference between the two users!
    - `worker_processes`
        - determines how many worker processes of Nginx would be running in the background → how many to **spawn simultaneously!**
        - Recommended to set this value equal to the number of CPUs or cores.
        - `worker_processes auto` → to automatically adjust the number of Nginx worker processes based on available cores.
        - To check the CPU cores:
            
            ```bash
            cat /proc/cpuinfo
            ```
            
            ![](https://i.imgur.com/QIK66bZ.png)
            
    - `error_log`
        - choosing the location of the log file for Nginx → `/var/log/nginx/error.log`
        - log level → `warn`
        - Two types of files at the location:
            - `access.log` → show the logs everytime we access our server
            - `error.log` → show any errors while server is running
    - `pid /run/nginx.pid`
        - shows the location of the `pid` file for Nginx
    - `include /etc/nginx/modules-enabled/*.conf`
        - List of all the enabled modules in Nginx → by default
        - importing them to be used in the configuration file
    
    ---
    
    - `events`
        
        ```bash
        events {
                worker_connections 768;
                # multi_accept on;
        }
        ```
        
        - No. of connections allowed per worker
        - Total number of connections your worker is allowed to make on a server
    - `http {}`
        - We’ll be spending most of our time in this block, to customise our configuration
- Overview of `default.conf` file
    - Location → `/etc/nginx/conf.d`
    - The default `nginx.conf` file includes a line that will load additional configurations files into the http { } context.
        
        ![](https://i.imgur.com/ue30Wkd.png)
        
        - In most cases, options specified in the primary `nginx.conf` file can be overridden by creating a new file at the above location i.e. `/etc/nginx/conf.d` or `/etc/nginx/sites-enabled`
    

---

- Nginx paths/routes are case-sensitive
    - Example:
        
        ```bash
        location /blog/ {
        
        		root /usr/share/nginx/blog;
        		index index.html index.htm
        }
        ```
        
        - The endpoints `/blog` & `/Blog` would not be the same

---

### Restarting Nginx

- Restarting the whole server
- Thus, making interrupts/drops in your connections

```bash
sudo systemctl restart nginx
```

### Reload Nginx

- Reloading the server, to save the config changes & keeps on running, without restarting the entire server!
- To prevent dropping the connections

```bash
sudo systemctl reload nginx
```

### Stopping Nginx

```bash
sudo systemctl stop nginx
```

---

### Nginx Reverse Proxy

- **Proxying** in simple terms is used to distribute the load among several servers, seamlessly show content from different websites.
- **Nginx reverse proxy** acts as an **intermediate** server that takes the client requests & forwards them to the appropriate backend server and subsequently delivers the server's response back to the client.
- **Hands-On**
    - Installing Nginx via Mainline installation method (through official Nginx packages)
    - Setting up Nginx Proxy
        - Creating a new config file in `/etc/nginx/conf.d/` named → `nodeapp.conf`
            
            ```bash
            server {
            
                listen 80;
                listen [::]:80;
            
                server_name 194.195.117.167;
            
                location / {
                    proxy_pass http://localhost:3000/;
                }
            }
            ```
            
            - `proxy_pass` → To pass a request to an HTTP proxied server, the [proxy_pass](https://nginx.org/en/docs/http/ngx_http_proxy_module.html?&_ga=2.56722146.934049599.1656297289-952190285.1656297289#proxy_pass) directive is specified
                
                ```bash
                proxy_pass <URL>
                ```
                
                - This can be used to proxy requests to an `HTTP` server (another NGINX server or any other server) or a `non-HTTP` server (running your own applications developed with a specific framework)
        - Testing the configuration for synxtax error n stuff like that:
            
            ```bash
            nginx -t
            
            #Output:
            nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
            nginx: configuration file /etc/nginx/nginx.conf test is successful 
            ```
            
    - Running the NodeJs app
        - You can now simple run your node js app by `npm start` & it should be serving in your remote IP address, in this case: [http://194.195.117.167](http://194.195.117.167/)
        - The standard Nginx `port-80` is **proxied** to `port-3000`
            - Therefore, we can direcltly open our application through the URL [http://194.195.117.167](http://194.195.117.167/) without the need to mention the specific port here!
                
                ![](https://i.imgur.com/9GG4Ly2.png)
                
        

---