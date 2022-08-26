# Understanding KubeSlice

- Docs: https://kubeslice.io/documentation/open-source/0.2.0
    
- What is KubeSlice?
    - A tool that combines **network**, **application**, **Kubernetes**, and **deployment services** in a framework to accelerate application deployment in a **multi-cluster**, **multi-tenant** environment
    - How it achieves this?
        - By creating logical application **`Slice`** boundaries that allow pods and services to communicate efficiently across clusters, clouds, edges, and data centers
        - It enables creating multiple logical **slices** in a single cluster or group of clusters regardless of their physical location
        - Network traffic Isolation between clusters is done by creating an `overlay network` for **inter-cluster** communication (among the clusters)
            - A second interface is added to the pod wherein, the traffic for the external clusters route over the overlay network to its destination pod
            - **This makes KubeSlice CNI agnostic** (it has no clue of the pod’s CNI used for intra-communication in a cluster )
- Concept of Multi-tenancy
    - Single instance/cluster of an application serves multiple customers
    - Each customer is referred to as `tenant`
    - Tenants have the ability to customise some parts of the application, such as:
        - colour of the UI
        - policies of the application
    - Tenants cannot change the application code
    - Single instance of the software will run on one server and then serve multiple tenants
    - Cluster resources must be fairly allocated among tenants
- How Multi-tenancy is achieved in Kubernetes?
    - **NOTE:** Kubernetes cannot guarantee perfectly secure isolation between tenants
        1. Can separate each tenant and their Kubernetes resources into their own [namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
            - Namespaces - provides a mechanism for isolating groups of resources within a single cluster. Its intended for use in environments with many users spread across multiple teams, or projects
        2. **Tenant Isolation** is then acheived using [policies](https://kubernetes.io/docs/concepts/policy/)
            - Policies - Defines what end-users can do on the cluster
                - Different policies we can apply:
                    1. network
                    2. volume
                    3. resource usage
                    4. resource consumption
                    5. access control
                    6. security
                - Open Policy Agent’s [Gatekeeper](https://www.notion.so/Understanding-KubeSlice-5bd1e4dfa4434931a3ac211afca30fc6) (a CNCF project) is used to enforces policies on Kubernetes clusters
    - Intra-Cluster communication in a cluster is carried out by utlising each pod’s **CNI (Container Network Interface)**
    - Why then use KubeSlice? (A simple answer)
        - It automates things
        - It tackles the problem that Kubernetes have of achieving multi-tenancy on a multi-cluster architecture.
- How does Kubeslice makes things easy (in cluster communitcation)?
    - KubeSlice creates a flat **overlay network** to connect the clusters
    - Overlay Network → an application slice that provides a slice of connectivity between pods of an application running in multiple clusters
        - Can also be referred to as app-specific VPC (Virtual Private Cloud) across clusters
    - Pods can connect to this **slice overlay network** & comunicate with each other across clusters!
    - This connection is secured by creating encrypted VPN tunnels, for inter-cluster traffic
- Use-Case (Service Discovery & Reachability)
    - Enables Kubernetes service (running in a cluster) to be exported through the slice overlay network & be discovered and reached by pods running in other clusters.

---