# Introduction to Containers - Docker

<aside>
📌 Resource → Civo DevOps BootCamp: [Session #2](https://youtu.be/CpybnXv_IeY)

</aside>

---

- What is a Container?
    - packing your application i.e the code, packages, modules, conf files etc & everything that your application needs to run, in a single box
    - Easy to ship your applications & run on different OS
    - A running instance of your Docker Image
- Containerisation v/s Virtualisation
    - **Virtualisation (VM)**
        - dividing the space of your hard-disk to create another operating system
        - is a little bit slow as compared to containers
    - **Containers**
        - creates an isolated environment of its own in our host OS only
        - really easy & fast way to run applications
        - uses less resources as compared to a VM

---

- A **container engine** is used to manage the containers running on our OS.
    - In this case, we are using **Docker**

---

- Docker Architecture
    - **Docker daemon**
        - listens to the API requests being made through the docker client
        - basically, this is the server that operates according to the command we give to the Docker Client
    - **Docker Client**
        - used to interact with Docker
        - the commands that a user gives to the server on what needs to be done
            - Example: `docker ps`, `docker run` etc.
    - **Docker Registeries**
        - Storage area for Docker Images
        - Most popular is **[DockerHub](https://hub.docker.com)**
        
        ![Screenshot 2022-04-22 at 6.01.22 PM.png](Introduction%20to%20Containers%20-%20Docker%20f8faf997e26a4601be86a1c33133c93f/Screenshot_2022-04-22_at_6.01.22_PM.png)
        

---

- **DockerFile**
    - The recipe of a Docker Image
    - Lists all the things required to run an application successfully
- **Docker Image**
    - Recipe for a docker container
    - File that defines a docker container
    - Whenever we run an image, a docker container is being created!
    - Images are built in layers

- Container deleted → data will be lost

---

- Common Docker Commands
    - Show the running containers
    
    ```docker
    docker ps
    ```
    
    - Run a docker image
    
    ```docker
    docker run imageName/ID
    ```
    
    - Start a container
    
    ```docker
    docker start imageName/ID
    ```
    
    - Kill a container
    
    ```docker
    docker kill ContainerName/ID
    ```
    
    - Remove a container
    
    ```docker
    docker rm ContainerName/ID
    ```
    
    - Inspecting a container image → showing some more info on that image
    
    ```docker
    docker inspect <image name>
    ```
    

---

### How does the workflow look like?

- When you run a command lets say `docker run`
- First, it will check whether that image is in your system or not!
- If NOT, it will head over to the Image Registry like **DockerHub** & pull that image from there (i.e download in your system)
- Now, you can RUN your image & your container will start!

---

- What if an image is not available in DockerHub
    - If we try to run the image, it would give an **`error`**
    - Generally, one is easily able to find an image of all the famous applications we use
    - We can create an image of our own application using **`Dockerfile`**
    - Organisations even have **private image registries** that they create for storing their own custom made images
- Difference between `docker run` & `docker start`
    - `docker run` is basically used to start a new fresh container
        - So, the state/data of the previous container will be lost
    - `docker start` is used to restart your container after stopping it (using `docker stop`)
        - here, we are able to preserve the state of the container & the data is not lost

---

## Container Runtimes

- Implementation of a container-runtime interface
- **Container-runtime Interface**
    - Bridge between the developer of the container runtime & the specifications of that
- **What should the container runtime be able to do?**
    - start/stop/delete the containers
    - push/pull the images
    - check the status of the containers
    - check the logs of a container & so on . . .

---

- Container Runtimes (in CNCF Landscape)
    - **containerD**
        - used to be part of docker
        - brought out of the main docker codebase
        - built as a separate project
        - got donated to the CNCF
    - **cri-o**
    - **Firecracker**
    - **gVisor**
    - **kata**
    - **lxd**
    - **SmartOS . . .**

---

- Parameters to consider while choosing a container-runtime
    - Performance
    - Stability
    - Hypervisor Isolation
    - Security Capabilities
    - Broad Usage
    - Multi-arch support

---

- **containerD**
    - container runtime with emphasis on:
        - simplicity
        - robustness
        - portability

---