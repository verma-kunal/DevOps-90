# Module-3: Using ArgoCD

---

### The Reconciliation Loop

- ArgoCD has a **3-minutes sync period**. That means every 3 minutes, ArgoCD will:
    1. Gather all applications that are marked to auto-sync.
    2. Fetch the latest Git state from the respective repositories.
    3. Compare the Git state with the cluster state.
    4. Take an action:
        - If both states are the same, do nothing and mark the application as `synced`
        - If states are different mark the application as `out-of-sync`
    5. For the applications marked as `out-of-sync`, you can take the decision of discarding the cluster or doing nothing.
- You can change the default sync period by editing the `argocd-cm configmap` found on the same namespace as Argo CD.
    
    Example:
    
    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: argocd-cm
      namespace: argocd
      labels:
        app.kubernetes.io/name: argocd-cm
        app.kubernetes.io/part-of: argocd
    data:
      timeout.reconciliation: 240s
    ```
    
    - Here, we setting the reconciliation/sync period to `240s` i.e. 4mins
    - To disable the sync functionality completely:
    
    ```yaml
    timeout.reconciliation: 0
    ```
    

---

- Using Git Webhooks
    - As mentioned before, ArgoCD take around 3 minutes (by default) to sync (poll) changes from Git manifests to your deployed application.
    - To eliminate this delay from polling, the API server can be configured to receive webhook events.
    - With Webhooks, as soon as there is any change in Git, Argo CD will run the sync process.
    - ArgoCD supports Git webhook notifications from GitHub, GitLab, Bitbucket, Bitbucket Server and Gogs.
    - More info can be found [here](https://argo-cd.readthedocs.io/en/stable/operator-manual/webhook/)

---

### Hands-On

- Created a new app - `demo`
    - with the default manifest files
- Deploy a new version
    - Commited a new change in `deployment.yml` file
    - Refresh the deployed app on the UI
    - Notice the yellow arrows on the deployment diagram - means the status is `OutOfSync`
        
        ![](https://i.imgur.com/hdTpixh.png)
        
    - To view the changes done on the UI, we can use `App Diff`
        
        ![](https://i.imgur.com/PH9AB7u.png)
        
        - 
    
    ---
    
    ### Sync Both Ways
    
- ArgoCD can also work in the opposite direction, when we are talking about syncing the changes
    - That is, If you make any change in the cluster then Argo CD will detect it and again tell you that something is different between Git and your cluster.
    - Lets change something in our cluster: (scaling replicas to 3)
        
        ```yaml
        kubectl scale --replicas=3 deployment simple-deployment
        ```
        
        ![](https://i.imgur.com/f0Kgni9.png)
        
- This capability is very powerful and you can easily detect configuration drift between your environments.

---

### Application Health

- ArgoCD keeps track of the service health for each deployed application.
- The health status values are:
    - `Healthy` **-** Resource is 100% healthy
    - `Progressing` - Resource is not healthy but still has a chance to reach healthy state
    - `Suspended` - Resource is suspended or paused. The typical example is a cron job
    - `Missing` - Resource is not present in the cluster
    - `Degraded` - Resource status indicates failure or resource could not reach healthy state in time
    - `Unknown` - Health assessment failed and actual health status is unknown
- Here, resources refer to the Kubernetes resources & whether the application is healthy or not, it all depends on these resources.
- Rules for built-in Kubernetes resources
    1. `Deployments`, `ReplicaSets`, `StatefulSets`, and `Daemon sets` are considered **“healthy”** if:
    **observed generation = desired generation**
    2. Number of updated replicas = number of desired replicas
    3. For a service of type `Loadbalancer` or `Ingress`, the resource is **healthy** if the `status.loadBalancer.ingress` list is **non-empty**, & at least has one value for `hostname` or `IP`
- For custom Kubernetes resources, health is defined in **[Lua scripts](https://www.lua.org/docs.html)**
    - You can add your own health check for a custom resource by implementing the check in **Lua** (a self-contained scripting language)
    - Example:
        
        ```lua
        hs = {}
        if obj.status ~= nil then
            if obj.status.phase == "Pending" then
                hs.status = "Progressing"
                hs.message = "Experiment is pending"
            end
            if obj.status.phase == "Running" then
                hs.status = "Progressing"
                hs.message = "Experiment is running"
            end
            if obj.status.phase == "Successful" then
                hs.status = "Healthy"
                hs.message = "Experiment is successful"
            end
            if obj.status.phase == "Failed" then
                hs.status = "Degraded"
                hs.message = "Experiment has failed"
            end
            if obj.status.phase == "Error" then
                hs.status = "Degraded"
                hs.message = "Experiment had an error"
            end
            return hs
        end
        
        hs.status = "Progressing"
        hs.message = "Waiting for experiment to finish: status has not been reconciled."
        return hs
        ```
        

---

### Sync Strategies

- 3 parameters that we can change while defining a sync strategy
    1. **Manual or automatic sync**
    2. **Auto-pruning of resources** - only applicable to automatic sync
    3. **Self-heal of cluster** - only applicable to automatic sync
- Manual or Automatic Sync
    - Defines what ArgoCD does when it finds a new version of your application in Git
        - If set to `automatic`:
            - will apply the changes then update/create new resources in the cluster
        - If set to `manual`:
            - will detect the change but will not change anything in the cluster.
    - To configure automatic sync run:
        1. CLI:
            
            ```bash
            argocd app set <APPNAME> --sync-policy automated
            ```
            
        2. In the application manifest:
            - Specify in the `syncPolicy`
                
                ```yaml
                spec:
                  syncPolicy:
                    automated: {}
                ```
                
- Auto-pruning
    - Defines what Argo CD does when you remove/delete files from Git
    - If `enabled`:
        - Argo CD will also remove the respective resources in the cluster as well
    - If `disabled`:
        - Argo CD will never delete anything from the cluster
    - To configure pruning as a part of automatic sync:
        1. CLI:
            
            ```bash
            argocd app set <APPNAME> --auto-prune
            ```
            
        2. In the app manifest:
            
            ```yaml
            spec:
              syncPolicy:
                automated:
                  prune: true
            ```
            
- Self-Heal
    - Defines what Argo CD does when you make changes directly to the cluster (via kubectl or any other way)
    - By default, changes that are made to the live cluster will not trigger automated sync.
    - If `enabled`:
        - Argo CD will discard the extra changes and bring the cluster back to the state described in Git.
    - To configure:
        1. CLI:
            
            ```bash
            argocd app set <APPNAME> --self-heal
            ```
            
        2. In the app manifest:
            
            ```bash
            spec:
              syncPolicy:
                automated:
                  selfHeal: true
            ```
            
- Various combinations of Sync Strategies
    
    ![](https://i.imgur.com/LpnDzrU.png)
    
    - **Policy A** (nothing is done by Argo CD) is how you should start adopting Argo CD , especially if you want to apply GitOps on an existing project (without deploying changes)
    - **Policy E** is the ideal one & at which one should look up to
    
**Hands-On**

1. **Deploying a new version with AutoSync**
    - Creating a new app - `demo`
        - with all the default configurations
        - repo: https://github.com/verma-kunal/gitops-certification-exercise
    -