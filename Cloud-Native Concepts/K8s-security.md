# Kubernetes Security - By Saiyam Pathak

Date: June 1, 2022
Number: 1

- A typical DevOps pipeline looks like this:
`Git → Build → Deploy → Release → Monitor`
- What is DevSecOps?
    - Developer, operations & the security are collaborating with each other
    - Its also a culture, similar to DevOps
    - Using best practices for security at each layer
    - Fail fast `==` fix fast
        
        ![](https://i.imgur.com/CbvmRkd.png)
        
    
- Why DevSecOps - Why care about Security now?
    - Due to the revolution from traditional architecture to micro-service architecture, there are lot of layers that have been formed in building an application.
    Example: k8s layer, container layer, cloud layer etc.
    - The no. of ways in which attacks can occur now with each layer, have increased
    - In the automated DevOps pipeline, **we need security at each layer/stage of the pipeline**, to safeguard our deployment
    - We just don’t want security at the bottleneck or at the end, when we have the end product ready
        - **Security at EVERY STAGE**
- 4Cs of Security in Cloud-native ecosystem
    1. Cloud - the infrastructure
    2. Cluster - running on top the cloud
    3. Container - running inside the clusters
    4. Code - which is in those containers
    
    ---
    
    - Aim: To provide security & scanning each layer for vulnerabilities
        
        ![](https://i.imgur.com/ExqP9Vy.png)
        
    - `AAA` (in a Cluster) - Authorisation, Authentication & Admission
    - **K8s Hardening guide by NSA/CSS (National Security Agency/Central Security Service)** - guidelines & best practices to follow
    - DockerFile best practices
    - OPA (Open Policy Agent) - was used previously for policy tooling
        - Kyverno is doing great work here
        - GateKeeper is a project, used for K8s policies specifically - by OPA

---

## Tools in CNCF Ecosystem

![](https://i.imgur.com/Usuf9x6.png)

- Cloud-native security landscape:
    - Prevention Layer
        - changing some behaviour of the application
    - Detection Layer
        - monitoring the process & alert
- Falco
    - CNCF Incubated project
    - Comes under the detection layer of cloud-native security landscape
    - Run-time security project
    - Working:
        
        ![](https://i.imgur.com/sJV0I7K.png)
        
        - Set of rules → YAML config file
        - Internal working:
        
        ![](https://i.imgur.com/BTJJE5L.png)
        
    
- Terrascan
    - CLI tool
    - detect violation across IAAC tom reduce the risk in the infrastructure
    - Used with:
        - Terraform
        - DockerFiles
        - YAML files → K8s
        - Helm
    - VS code extension → Rego editor
- Kubescape - By ARMO
    - Came after NSA Hardening guide - aheres to the guide
    - Ensure security from dev to production for:
        - K8s clusters
        - Microservices
        - Pods
    - Built for developers
        - Nice approach & makes things easy for developers to work on
    - Includes:
        
        ![](https://i.imgur.com/3cpq04q.png)


        
        - `ARMO | Best` → policies that ARMO thinks are the best to use
        - RBAC Visualizer

---