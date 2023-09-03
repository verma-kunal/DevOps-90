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
    3. Error Control
       - Some data will be lost / corrupted
       - It adds checkSum to every data segment
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

### Resources Used

- [Computer Networking Course](https://youtu.be/IPvYjXCsTg8)
