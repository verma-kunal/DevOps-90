# The OSI Model

---

- Open System’s Inter-connection model
- A standard for how people & machines communicate with each other

---

- There are 7 layers in the OSI Model
  1. **Application**
  2. **Presentation**
  3. **Session**
  4. **Transport**
  5. **Network**
  6. **Data Link**
  7. **Physical**
  ***

---

### An overview

- Application Layer
  - Implemented in Software
  - Users are able to interact here
  - Example: Browsers etc.
  -
- Presentation Layer
  - Get data from the Application Layer
  - Converts the data into machine-representable binary format
    - ASCII to EBCDIC
    - process is called as **Translation**
  - Encoding + Encryption
  - Provides abstraction
  - Data is compressed (lossy or lossless)
    - Makes it easy to transfer
- Session Layer
  - Helps in setting up & managing connections
    - - the termination of the connected sessions
  - Authentication & Authorisation is involved here
- Transport Layer
  - Making sure that the data is transferred effectively
  - Protocols involved → `TCP, UDP`
  - Transportation is done in 3 ways:
    1. Segmentation
       - Data received from the Session layer is divided into small packets called **Segments**
       - Each segment will have:
         1. Destination & Source PORT number
         2. Sequence Number
            - Helps in re-assembling the segments in the correct order
    2. Flow Control
       - Controls the amount of data that is being transferred
       - Error control is managed here
       - Adds a `checkSum` to every data segment
         - Quality of data check
- Network Layer
  - Mainly works on the transmission of the segments from 1 computer to another, that is located in different networks
  - **Router** lives over here
  - **Logical Addressing** → assingns the sender’s & receiver’s IP address to every segment `(IP packet)`
  - Moving a packet from source to destination → called **Routing**
  - Load-balancing happens over here
- Data Link Layer
  - A data packet consists: (Logical Addressing)
    ![](https://i.imgur.com/ws9Q0Pk.png)
  - This layer does **Physical Addressing**
    - `MAC Addresses` of sender & receiver
      - Forms a **Frame** → data unit of the Data Link Layer
    - `MAC Address` → 12-digit alpha-numeric number of the network interface of your computer
      ![](https://i.imgur.com/knFBEEz.png)
  ***
- Physical Layer
  - Consists of hardware like wires etc.
  - Transmits using `bits`
  - In the form of signals

---

![](https://i.imgur.com/ijG1zkd.png)

---

### TCP/IP Model

- known as the **Internet-protocol** suite
- Mostly similar to the OSI Model
- Was developed by **ARPA**
- 5 Layers:
  1. Application Layer
  2. Transport Layer
  3. Network Layer
  4. Data Link layer
  5. Physical Layer
- This is used practically!

---

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

### Resources Used

- [Computer Networking Course](https://youtu.be/IPvYjXCsTg8)
