# Application Layer

---

- Main layer where the users interact!
- Consists of applications like web browsers, whatsapp, youtube etc.
- Where does this lie? → On our devices

---

### Client-Server Architecture

- Question
  How does these applications talk to each other & sends requests from the users?

![](https://i.imgur.com/pIizpvE.png)

- Processes:
  - Client
  - Server-side
- These processes are communicating with each other through
  `IP addresses`
- There are multiple servers, if we talk about big projects/apps
  - Collection of servers is called a **Data Center**
  - Have **static** IP addresses
- We can use the `ping` command to see how requests are sent to a server:
  ![](https://i.imgur.com/oPl3D3c.png)

  - `icmp_seq` → is the sequence number of that packet
- `ping time` → measures the round-trip time for messages sent from origin host to destination & then comes back.

---

### Peer-to-Peer Architecture (P 2 P)

- Applications running on various devices are connected & can communicate with each other
- No single server or a datacenter
- Scaling is fast
- Every single system can be termed as a Client or a server
- Example → **BitTorrent**
  ![](https://i.imgur.com/LrWnHLQ.png)

---

### Protocols

**Web Protocols**

1. `**TCP/IP**`
   - HTTP → Hypertext Transfer Protocol
     - defines how data is transferred
   - DHCP → Dynamic Host Control Protocol
     - allocates IP adresses to the devices connected to your network
   - FTP → File Transferred Protocol
     - how files are transferred across the internet
     - not used much now as this is done through HTTP
   - SMTP → Simple Mail Transfer Protocol
     - how emails work
     - use to send the emails
   - POP3 & IMAC
     - use to receive emails
   - SSH → Secure Shell
     - to connect with remote servers
   - VNC → Virtual Network Computing
     - for graphical control
2. **`Telnet`**
   - A terminal emulation that enables a host or device to connect to a remote server using the **telnet client**
   - PORT → 23
   - low-level protocol
3. **`UDP`**
   - Stateless connection i.e does not maintain the state
   - Date maybe lost here

---

- A process is the running instance of a program
- A thread is the lighter version of a process
  - Thread only does 1 single-job

---

- **How do these devices communicate with each other over the internet?**
  - IP Addresses
  - PORT Numbers

---

### Sockets

- Sending messages from 1 system to the other
- Interface between the process & the internet

---

### PORTS

- Ports tell us which application we are working within a device

**Ephemeral Ports**

- When multiple application instances are running, automatic PORTs will be assigned to processes.
- This can exist on the **client-side** but on the server, we need structured & defined PORT numbers.

---

### HTTP - Hypertext Transfer Protocol

- Referring to the client-server architecture here!
  ![Screenshot 2022-07-21 at 6.55.15 PM.png](Application%20Layer%20908cb1b88f12400284cfb4dcfe4a1ab5/Screenshot_2022-07-21_at_6.55.15_PM.png)
- This is a **Client-Server** protocol
- Basically defines:
  - How the data is requested to the server by the client →
    **HTTP request**
  - How the server will respond to that particular data → **HTTP Response**
- **Every Application Layer protocol (HTTP) require some Transport Layer protocols as well**
- Some of the methods associated with HTTP:
  - `GET`
  - `POST`
- HTTP uses `TCP` (Transmission Control Protocol)
- HTTP is a **stateless** protocol
  - not save or store the client’s info

---

### HTTP Methods

1. **GET**
   - Requesting some data or page
2. **POST**
   - Giving something to the server as a client
3. **PUT**
   - Puts data at a certain location
4. **DELETE**
   - Delete data from a server

---

### Error/Status Codes

To know about the status of a particular request (whether it was sent, failed etc.)

- 200 → Successful
- 404 → Not found
- 400 → Bad request
- 500 → Internal Server Errors

**Classes associated:**

- `1 X X` → Informational Category
- `2 X X` → Success Codes
- `3 X X` → Redirecting
- `4 X X` → Client-side error
- `5 X X` → Server-side error

---

### Cookies

- A unique string
- Stored on the client-side browser (file)
- After visiting the browser once, a cookie will be created → storing your info
  - The next time when you’ll request, a cookie will be sent to the server (which has your saved info) & the browser would retrieve the previous session through that
- **Third-party cookies**
  - These are set for the URL that we don’t
  -

---

### How Email works?

- Application Layer protocol used here is **SMTP (Simple Mail Transfer Protocol)** → to send emails
    - To receive emails, we an use **POP3** protocol
- Transport layer protocol used here is TCP
    - Because we don’t want any data to be lost (as it gets in UDP)

![](https://i.imgur.com/PkkS4Uf.png)

- A command to find more info about the SMTP serves for various domains - `nslookup`
    
    ![](https://i.imgur.com/ER0z1Yn.png)
    
    - Here `-type=mx` means “Mail Exchange” i.e. SMTP servers
- **Downloading Emails**
    - POP → Post Office Protocol (PORT 110)
    - POP connects with the client-server & then transact all the emails.

---

### Domain Name System (DNS)

- Question
    
    How does the internet finds a particular server for the domain name & the associated **IP address** ?
    
- Domain names are mapped to IP addresses
- When we write `google.com`, that uses DNS to find the IP address of Google’s server
- `HTTP` → converts the URL we type into an IP address (using DNS) to find the desired server
- DNS → Directory or database that has all the IP addresses of servers

---

- **Classes of Domains:**
    
    ![](https://i.imgur.com/kfJgTcd.png)
    
    - Instead of storing data in at a single place, there are multiple places for these 3 categories
        1. **Root DNS Servers**
            - For top-level domains
            - 1st point to contact when you search a domain
            - Examples: `.io`, `.org`, `.com` etc.
            - Who is maintaining the root DNS servers
                
                [https://root-servers.org](https://root-servers.org/)
                
            - Managed by **ICAN** (Internet Corporation for Assigned Names & Numbers)

---

- **Lifecycle of connecting to a server via a domain name:**
    
    ![](https://i.imgur.com/BANRI8L.png)
    

---

- Some organisation have their own top-level domain names.
    - Example: Google has `.com`
- **FACT** → You cannot buy your own domain name, can only rent it

---

- `dig` command (Domain Info Groper) is used for retrieving information about DNS name servers
    
    ![](https://i.imgur.com/gU0jomK.png)
    

---

### Resources Used

- [Computer Networking Course](https://youtu.be/IPvYjXCsTg8)
