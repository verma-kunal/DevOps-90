# Build & Testing of Applications (CI/CD)

<aside>
📌 Resource → Civo DevOps BootCamp: [Session #3](https://youtu.be/ppdBp1GvGpw)
</aside>

---

- **Continuous Integration** part of the DevOps lifecycle

---

### Where do testing & building fit in DevOps?

- **Testing** is not just in one place

---

- Continous Integration v/s Continuous Delivery
    - **CI**
        - getting your changes merged into the deployable state frequently & continuously
        - Extreme version of this is → commiting after each change made
    - **CD**
        - As soon as you merge the code into your main branch or deployable state, it goes straight into the production environment
        - There can be a few stages in b/w like performancing testing, reviews etc. but the main idea is above ☝🏻
        - A very few orgs have this as it might be dangerous to the application for some cases
        

---

### Continuous Integration

- Integration of small changes
- Every time you commit → code is tested
- New build with every commit to `main` branch

---

- Tests done in **CI & CD**
    - **CI**
        - Unit Test
        - Integration Test
    - **CD**
        - Full system tests
        - Performance Tests
        - Conformance tests

---

### What do we mean by “Testing”?

- we actually mean **Automated Testing**
- giving us the confidence that our application is working
    - “If I add this feature, how do we know the product still works”
- We can write tests in every stage of the DevOps cycle
- Types:
    - **Unit Tests**
        - we are testing small parts of the code like functions or some lines of code
        - really quick to run
    - **Integration Tests**
        - whether several different functions are able to work together efficiently or not
    - **End-to-End Testing**
    - **Conformance Testing**

---

- Pyramid of Testing
    - Lower the rank of the test is, faster it would be ⇒ take less time to run
        1. Unit tests (in large numbers)
        2. Integration tests 
            
            .
            
            .
            
            .
            

---

- Faster the test, you should run them frequently in your development stage!
- Good tests == honest feedback about your application

- Test driven Development (`19:43`)
    - you start with a **failing test**
    - you try to reproduce an issue in your application by creating a test for each part of it
    - These test would run & if there is a failing test, you would know:
        - why its failing?
        - where its failing?
        - & what you need to do to fix it!
    

---

### CI Pipelines

- **Pipeline** → set of automated process that take your application from developing → to running tests → to deploying → to the production environment
- `⇒ x ⇒ y ⇒ z ⇒ . . .`  : At each step, some processes would automatically be executed & this sort of pathway is called as a Pipeline
- Tools:
    - Jenkins
    - circleCI
    - TravisCI
    - GitLab CI/CD
- **which tool to choose?**
    - depends on the budget, your requirements etc . . .
    - it should be able to continously integrate your code
- configuring CI Pipelines is in the **Operations** side of things in DevOps

---

- In an application, a function might depend on another function to work & **writing tests** can help us detect any sort of errors that might break the code
    - Also gives us confidence
- Every programming language has a separate testing framework to be used for that specific language
    - Example → Golang has `go test` etc.
- Configuring the CI Pipeline & adding tests there is a separate task & is done differently!

---

**A great feeling when all the test checks pass  🎉**

![](https://i.imgur.com/dNH9UMa.png)

---
