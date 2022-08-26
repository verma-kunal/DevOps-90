- KubeSlice Architecture
    
    ![](https://i.imgur.com/h5ufHNo.png)
    
    ---
    
    ### Components
    
    KubeSlice consists of the following main components deployed in one or more gateway nodes:
    
    1. **Kubeslice Controller**
        - Its installed in one of the cluster (the controller cluster)
        - Provides a central configuration management system for slices across multiple clusters.
        - Facilitates the creation of the slice overlay network between multiple worker clusters
        - Configuration and operational information in the cluster is stored using [CRDs](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) (Custom Resource Definitions)
        - The worker clusters connect to the **[Kubernetes API server](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/)** of the controller cluster to fetch configuration that is stored in the custom resource objects.
        - Slice Configuration includes:
            - slice VPN gateway
            - service discovery with service import/export
            - ingress/egress gateway related parameters
    2. **Slice Operator**
        - Also known as a **Worker Operator** is a Kubernetes Operator component that manages the life-cycle of the KubeSlice related CRDs
            - Kubernetes Operator → A [Kubernetes operator](https://www.redhat.com/en/resources/oreilly-kubernetes-operators-automation-ebook?intcmp=701f2000001OMH6AAO) is a method of packaging, deploying, and managing a Kubernetes application.
            - Function of Slice Operator - [Here](https://kubeslice.io/documentation/open-source/0.2.0/architecture#slice-operator)
    3. **Slice VPN Gateways**
        - A slice network service component
        - Provides a VPM tunnel between multiple clusters that are a part of the slice configuration
        - Slice Operator performs some life-cycle functions for Slice VPN Gateways - [Here](https://www.notion.so/KubeSlice-a7fcb42b4af5483baa9e009fc38457df)
        - **`NOTE:`** Kubeslice Controller manages the VPN gateway pairs for the attached clusters, and creates the keys & configurations required for the operation
    4. **Slice Router**
        - Its a virtual layer-3 device that sets up & performs the routing & forwarding rules in the slice overlay network.
        - One slice router pod → per slice (minimum requirement)
        - **Important:** Slice operator manages the life cycle of the slice router deployment and monitors its status periodically.
    5. **Slice Istio Components**
        - KubeSlice provides the option of setting up ingress and egress gateways for a slice using [Istio Service Mesh](https://istio.io/) resources
        - Slice Ingress/Egress Gateways are used for
            - **Internal East-West traffic** (inter-cluster, egress from one cluster, and ingress into another cluster)
            - **Slice North-South** Ingress Gateway for external traffic
        - **`NOTE:`**
            - `Ingress` is incoming traffic to the pod
            - `egress` is outgoing traffic from the pod
    6. **KubeSlice DNS**
        - DNS server that is used to resolve service names exposed on application slices
        - Slice Operator manages the DNS entries for all the K8s services exposed on the slice overlay network
        - A service is exported on the slice by installing a `ServiceExport` object
    7. **Network Service Mesh Control and Data Plane**
        - Network Service Mesh (NSM) component sets up the Kubeslice data plane and connects application pods to the slice overlay network
            - Control Plane & Date Plane
                
                **Control Plane**
                
                - The `brain` of K8s cluster
                - Here, Kubernetes carries out communications internally, and where all the connections from outside — via the API come into the cluster to tell it what to do
                - Here, all the cluster configurations, scheduling and communications happens
                - The main components:
                    - etcd
                    - API Server
                    - Scheduler
                
                **Data Plane**
                
                - The `body`of K8s cluster
                - Worker nodes (i.e. VMs) on the data plane carries out commands from the control plane and can communicates with each other via the `kubelet`, while the `kube-proxy` handles the networking layer
        - Difference between Kubeslice & Kubernetes Data Plane
            
            
    8. **NetOps**
        - Each slice in a cluster is associated with a `QoS` profile for bandwidth control across the clusters
            - QoS Profile → **rules that are used to control communication from a source sites to destination sites**
        - This is applied to the external interface of the VPN gateway nodes
        - `NetOps` pods → configure & enforce QoS profile for a slice on a cluster

---