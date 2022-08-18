# Network Layer

---

- **`NOTE`**
    - In Transport Layer, data is transferred as **Segments**
    - In Network Layer, data is transferred as **Packets**
    - In Data-Link Layer, data is transferred as **Frames**
- In this layer, we work with **routers**
- Every single router that basically connects together to transfer data, has its own **Network Layer Address**
    
    ![](https://user-images.githubusercontent.com/72245772/185274649-20c41c88-e77a-4d5f-b80e-958ad7500542.png)
    
- **Routing Table & Forwarding Table**
    - Every router consists of these both
    - Forwarding Table comes inside the Routing Table
        - Routing table consists of all the possible paths that can be taken to transfer the data packet
        - Forwarding Table consists only 1 path (therefore, faster)
    - Whenever a packet arrives to a particular router/network layer address, it would check for the packet’s desired destination in the **Forwarding Table** & then send/forward the packet to that desired router
        - this is called as **Hop-By-Hop Forwarding**
            - Forwarding the packet to routers, till it reaches its destination!

---

- **Breakdown of IP Address**
    
    ![](https://user-images.githubusercontent.com/72245772/185274651-2be15a4f-657e-4dfe-ba30-989344f33ef4.png)
    

---

- Question
    - Who creates these tables?
    - Answer: **Control Plane**

### Control Plane

- Used to build the **routing tables**
- If we look it as a Graph:
    - **Routers** → Nodes
    - **Links b/w routers** → Edges
- Two types of routing used to create the tables:
    1. **Static Routing**
        - Adding the addresses manually in the routing table
        - Not that much adpative
        - A bit complex
    2. **Dynamic Routing**
        - When there is a change in network, this evolves accordingly
        - Some of the algorithms that are used here: (Path-Finding Algorithms)
            - **[Bellman–Ford Algorithm](https://www.geeksforgeeks.org/bellman-ford-algorithm-dp-23/)**
            - **[Dijkstra's Algorithm](https://www.programiz.com/dsa/dijkstra-algorithm)**

---

### IP (Internet Protocol)

**IP Addresses** 

**IPv4**

- 32 bits
- 4-words
    
    ![](https://user-images.githubusercontent.com/72245772/185274654-0249195a-faea-4d1e-b3e8-5abc603ceff1.png)
    

**IPv6**

- 128 bits

---

![](https://user-images.githubusercontent.com/72245772/185274657-69d4d6c5-4074-4622-8513-54343391a397.png)

---

- **Class Of IP Addresses**
    - Class A → `0.0.0.0` to `127.255.255.255`
    - Class B → `128.0.0.0` to `191.255.255.255`
    - Class C → `192.0.0.0` to `223.255.255.255`
    - Class D → `224.0.0.0` to `239.255.255.255`
    - Class E → `240.0.0.0` to `255.255.255.255`

---

### Data Packets

- **Header** → 20 bytes
    - IP version
    - Total length
    - Identification Number
    - Flags (like SYN, ACK etc.)
    - CheckSum
    - **TTL (Time To Live)**
        - After a certain hops (in order to reach the destination) if the packet doesn’t reach the destination, it will be dropped
            
            ![](https://user-images.githubusercontent.com/72245772/185274660-ffb80754-5182-407f-acc5-6c5324a6c4ee.png)
            

---

### Resources Used

- [Computer Networking Course](https://youtu.be/IPvYjXCsTg8)