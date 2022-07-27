# Codefresh - GitOps Fundamentals

Link: https://codefresh.learnworlds.com/course/gitops-with-argo

### What is GitOps?

- GitOps is a set of best practices where:
    - the entire code delivery process is controlled via `Git`:
        - including the infrastructure
        - application definition as code
        - automation to complete updates and rollbacks (going back to previous version)

### GitOps Principles

1. **Declarative**
    - A system managed by GitOps must have its desired state (what we want the application to do) expressed declaratively
        - That means, we just tell the infrastructure that certain thing has to be done & it takes care of the rest of the things
2. **Versioned & Immutable**
    - The entire infrastructure is versioned in `Git`
    - The desired state is stored in a way that enforces immutability (not capable to change)
3. **Pulled Automatically**
    - Approved changes are automated & applied to the system
4. **Continuously Reconciled**
    - Software agents ensure that the actual system is being continuously observed & they **attempt to apply** the desired state to it

---

### GitOps Deployment in Kubernetes

- In **Kubernetes**, GitOps deployments happens as follows:
    1. A GitOps agent is deployed on the K8s cluster
    2. The GitOps agent is monitoring one or more Git repositories that define applications and contain Kubernetes manifests (or Helm charts or Kustomize files).
        - `Kubernetes Manifests` → plaintext config files that describe how a pod's containers should be run and connected to other objects, such as services etc.
    3. Once a Git commit happens the GitOps agent is instructing the cluster to reach the same state as what is described in Git.
    4. Developers, operators and other stakeholders perform all changes via Git operations and **never directly touch the cluster** (or perform manual `kubectl` commands)

---

### Traditional Deployment v/s GitOps Deployment

**Traditional Deployment**

![](https://i.imgur.com/yfMdeT7.png)

---

**GitOps Deployment**

- Nobody has direct access to the Kubernetes cluster
- There is a second Git repository that has all manifests that define the application.
- A **GitOps controller** that is running inside the cluster is monitoring the Git repository and as soon as a change is made, it changes the cluster state to match what is described in Git
    - The GitOps controller is running in a constant loop and always matches the Git state with the cluster state.
    
    ![](https://i.imgur.com/z8M52YO.png)
    

---

### GitOps Use Cases

- The most popular use-case is **Continous Deployment of Applications**
- **Continuous Deployment on Custom resources**
- **Continuous Deployment of Infrastructure**
- **Detecting/Avoiding Configuration Drifts**
    - You can also set up your agent so that instead of applying any change immediately, it will just notify you of the discrepancy.
    - As, **Ad-hoc** changes on infrastructure are a well known problem.
    - Configuration drift between what you think is on a server and what is actually on the server is one of the most common causes of failed deployments.
        - This power of detecting configuration drift early is one of the major differences between GitOps and traditional deployment solutions.
- **Multi-Cluster Deployment**

---
