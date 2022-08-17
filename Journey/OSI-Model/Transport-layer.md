# Transport Layer

---

- In a situation where one message has to be transferred from 1 device to another, this transportation is done by **Network Layer**
- **Transport Layer is present in our devices**
    - Its job is to take the message received from the network layer & take that to the desired application and vice versa
        
        ![](https://i.imgur.com/AuJBbuv.png)
        
- Protocol followed:
    - **TCP**
    - **UDP**

---

- **Multiplexing**
    - a technique used to combine and send the multiple data signals over a single medium
    - Another name - ***muxing***
- **De-Multiplexing**
    - Delivering received signals/messages at the receiver side to the correct apps on another system.

![](https://i.imgur.com/yaoS8VT.png)

---

- Application are tracked using `PORT` Numbers
- **Remember:**
    - Transport Layer also takes care of **congestion control** (traffic control)
    - Congestion Control algorithms are built-in **TCP**

---

### CheckSum

- **CheckSum** → A string of certain numbers that is associated with the data being sent or received
- If the `checkSum sent ≠ checkSum received`  , then something has go wrong while data was being transferred.
- This ensures that the same data that was sent, is being received by the end-user

---

### Timers

- Question we tackle here
    
    How would we know that the data packets have been received to the end-user device?
    
- When the packet is sent, a timer is started at the sender’s side.
- When the packet is received by the receiver & they comfirm it, the timer on the sender’s side end → **which indicates that the data packet has been received!**
- In a case where the data packet is lost in the middle of transmission & not received by the receiver → **timer gets expired**
- This is called as **re-transmission timer**

---

### UDP (User-Datagram Protocol)

- In this case, data may or may not be delivered
- Data may change on the way
- Data may not be in order, when received
- Its a **connection-less protocol**
    - No connection is established between the devices/computers before sending the data
- UDP uses `checkSums`
    - If the data is corrupted, UDP would show that, but won’t do anything about it
- **Components of a UDP Packet**
    - Source PORT & Destination PORT numbers
    - Length of datagram
    - CheckSum
    - The data 
        
        ![](https://i.imgur.com/QfS1bza.png)
        
    
    ---
    
- Why do we use this? (Use-cases)
    - Its faster as compared to other protocols
    - Used in Video Conferencing apps, Gaming etc.
    - DNS → uses UDP
- A command to check the packets that are being transferred from your computer - 
`sudo tcpdump`
    
    ![](https://i.imgur.com/4epdibn.png)
    

---

### TCP (Transmission Control Protocol)

**(Important)**

- Its a **transport layer** protocol
- The raw data sent by the Application layer is segmented (divided into chunks, add header, etc.) by TCP
- Congestion Control or Traffic Control
- Takes care of the following operations:
    - When the data is not received or arrived
    - Maintains the order of data being transferred
        - Using **Sequence Numbers**
    
    ---
    
- Features
    - Its **Connection-Oriented**
        - First a connection is established between the devices & then data is transferred
    - Error Control
    - Congestion Control
    - Its **Full Duplex**
        1. A can send files to B
        2. B can send some other files to A
        3. Both A & B can send files/data simultaneously 
            
            ![](https://i.imgur.com/plxTjFg.png)
            
    - One TCP connection can only be established between 2 computers or endpoints **(Important)**
- Components of TCP packet
    - The only difference we have here from the UDP packet is, the addition of two components
        - Sequence Number (random numbers)
        - Acknowledgement Number
    - **3-Way Handshake**
        
        ![](https://i.imgur.com/BG61DOF.png)
        
        - Here:
            - `SYN` - Synchronisation Flag
            - `ACK` - Acknowledgement Flag
        - To be continued . . .

---

### Resources Used

- [Computer Networking Course](https://youtu.be/IPvYjXCsTg8)