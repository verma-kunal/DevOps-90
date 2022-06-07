# Release & Deploy

---

- How we can make a configurable Release & deploy cycle using `CI/CD` & `IaC` tools

---

- Having the ability & agility to respond to business needs
    - plus the Developer & Operations team sanity

---

- **High-Mature DevOps Orgs:**
    - Automatic Deployments
        - scripts
    - Short-lived automated environments
    - Cloud/Hybrid cloud
    - Cross-functional teams
    - Multiple deploys a day
        - You become better at stuff when you do that frequently!
        Therefore, good quality deployments
- Code needs to be tested on an environment, similiar to the **production environment**

---

### CI/CD

![](https://i.imgur.com/NYBoRbd.png)

- Choosing a particular set of tools, depends on your organisation’s culture or needs

---

### Pipelines

- Set of steps executed in order
- Configured as code
    - Therefore, they can be peer reviewed
- The previous steps must be passed, for future steps to run
- Example of a pipeline:
    1. Linting
    2. Build
    3. Unit Test
    4. Deploy
    5. Scan
    6. Static Code Analysis
    7. Notification

---

### Infrastructure As Code (IaC)

- **What is an Infrastructure?**
    - The cloud resources that we use!
- Defining **in the form of code**, what sort of & how much resources we need in our K8s clusters, VMs etc. to run our application
- Some of the tools are:
    
    ![](https://i.imgur.com/6LbFBgG.png)
    

---
