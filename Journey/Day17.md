# Building Blocks For OSI Model

---

### How Data is transferred?

- Data is transferred in small chunks called **Packets**

**How do we know which server to connect to?**

- Through IP address
- Every server connects with one another through IP addresses
- **Format:**
    - `**X.X.X.X**`
        - `(0-255).(0-255).(0-255).(0-255)`

---

- Types of IP addresses
    - **Global IP** addresses
    - **Local IP** addresses
- **DHCP (Dynamic Host Configuration Protocol)** → used to dynamically assign an IP address to any device, or node, on a network so they can communicate using that IP address
- Example:
    
    ![](https://i.imgur.com/txXVYvU.png)
    
    - Here **ISP** → Internet Service Provider
    - Example: Airtel, Jio etc.
- **NAT (Network Access Translator)** → helps the router determines which Device sent the request (through which local IP address) ⇒ refer the example above

---

### PORT Numbers

**Q. Which application in a particular device (local IP address) is requesting the data?**

- This is determined using **`PORTS`**
- a way to identify a specific process in a particular device
- PORTS are `**16-bit**` numbers
    - Total PORT Numbers possible are = 2^16
        
        ![](https://i.imgur.com/loYtbmi.png)
        
- PORT numbers defined for various networking protocols:
    
    ![](https://i.imgur.com/rRXhfb4.png)
    

---

- PORT Numbers from `0-1023` are reserved
    - You’ll not be able to use those for personel usage
- PORT Numbers from `1024-4915` are reserved for some applications like MongoDB, mySQL etc.
    - Example → every SQL server you run, has a PORT of `**1433**`
- PORT Numbers **after 4916** are not reserved & can be used publicly

---

### How do we measure the internet speed?

- **`1 mbps`**
    - mbps → Mega-bits per second
    - 1 bit → 0 or 1
    - Mega → 6 0’s
    - Therefore:
        - **1mbps = 10,00,000 bits/sec**
- Some common terms used here:
    - **Upload** → You are sending some data to another computer
    - **Download** → You are receiving some data in your computer

---

### How are countries connected with each other?

- Through cables that are laid out under the ocean
- To know more → [https://www.submarinecablemap.com](https://www.submarinecablemap.com/)

![](https://i.imgur.com/4xfQGNH.png)

---

- Different mediums to connect:
    - **Physical** → Optical fibre cables etc.
    - **Wireless →** Bluetooth, Wifi, etc.

---

### LAN, WAN, MAN

- **LAN** (Local Area Network)→ Small houses/offices
    - via ethernet cables, wifi, etc.
- **MAN** (Metropolitan Area Network) → Across a city
- **WAN** (Wide Area Network) → Across countries
    - via Optical fibre cables
    - Types:
        1. **SONET** (Synchronous Optical Networking) → carries data using optical fibre cable (covers large distances)
        2. **Frame Relay →** way to connect LAN to a wider area (the internet)

**⇒ Internet is a collection of all these 3 types of networks!**

---
### Resources Used

- [Computer Networking Course](https://youtu.be/IPvYjXCsTg8)
