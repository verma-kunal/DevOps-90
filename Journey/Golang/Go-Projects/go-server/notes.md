# Go: Simple Web Server

Status: Done

<aside>
🔗 Link: [https://www.youtube.com/watch?v=jFfo23yIWac&t=0s](https://www.youtube.com/watch?v=jFfo23yIWac&t=0s)

</aside>

---

### Pre-requisites

- Golang installed in your system
    - Here: [https://go.dev/doc/install](https://go.dev/doc/install)
- Basic knowledge of Golang

---

- Packages we need to import
    - `fmt` - [https://pkg.go.dev/fmt](https://pkg.go.dev/fmt)
    - `log` - [https://pkg.go.dev/log](https://pkg.go.dev/log)
    - `net/http` - [https://pkg.go.dev/net/http](https://pkg.go.dev/net/http)
    
- An overview of the project
    
    ![](https://i.imgur.com/zLkNocg.png)
    
- Creating the HTML fIles (to serve)
    - `index.html`
        
        ```html
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>Main Page</title>
        </head>
        <body>
            <h2>Main Page!</h2>
        </body>
        </html>
        ```
        
    - `form.html`
        
        ```html
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>Form</title>
        </head>
        <body>
            <div>
                <form action="/form" method="post">
                    <label for="fname">First name:</label><br>
                    <input type="text" id="fname" name="fname"><br>
        
                    <label for="lname">Last name:</label><br>
                    <input type="text" id="lname" name="lname">
        
                    <input type="submit" value="Submit">
                </form>
            </div>
        </body>
        </html>
        ```
        
- Main Logic (in `main.go`)
    - Setting the path for the static files folder:
        
        ```go
        fileServer := http.FileServer(http.Dir("./static"))
        ```
        
        - `http.FileServer` → it serves up an entire directory path and all its children
    - We’ll define a Handle which is for our index.html file from “static” folder:
        
        ```go
        http.Handle("/", fileServer)
        ```
        
    - We’ll define a `HandleFunc()` that has the arguments:
        - the path/route
        - the handler function that we define (which has the information of what to serve)
        
        ```go
        http.HandleFunc("/hello", helloHandler)
        http.HandleFunc("/form", formHandler)
        ```
        
        - Here `helloHandler` & `formHandler` are the two functions that we’ll be creating to serve something (we’ll define that shortly)
    - Creating the `helloHandler()`
        
        ```go
        func helloHandler(w http.ResponseWriter, req *http.Request) {
        	// writing some basic checks:
        
        	// 1. If the path is not "/hello":
        	if req.URL.Path != "/hello" {
        		http.Error(w, "404 Not Found", http.StatusNotFound)
        	}
        	// 2. If the method of our request is not "GET"
        	if req.Method != "GET" {
        		http.Error(w, "Method is not supported", http.StatusNotFound)
        	}
        
        	// In an ideal case, do this:
        	fmt.fprintf(w, "Hello World!")
        }
        ```
        
    - Creating the `formHandler()`
        
        ```go
        func formHandler(w http.ResponseWriter, req *http.Request) {
        	err := req.ParseForm()
        	if err != nil {
        		fmt.Fprintf(w, "ParseForm() error: %v", err)
        		return
        	}
        	// print when our request is successful
        	fmt.Fprintf(w, "POST request successfull")
        
        	// getting the values from our form (on-submit):
        	firstName := req.FormValue("fname")
        	lastName := req.FormValue("lname")
        
        	// printing the values to the screen:
        	fmt.Fprintf(w, "First Name: %d\n", firstName)
        	fmt.Fprintf(w, "Last Name: %d\n", lastName)
        
        }
        ```
        
        - We are using the `ParseForm()` to parse our form request data (POST)
            - Its always used on the **request object** (req)
- Output
    - [http://localhost:8080/](http://localhost:8080/)
        
        ![](https://i.imgur.com/IxQSYfH.png)
        
    - [http://localhost:8080/hello](http://localhost:8080/hello)
        
        ![](https://i.imgur.com/ZITudCr.png)
        
    - [http://localhost:8080/form.html](http://localhost:8080/form.html)
        
        ![](https://i.imgur.com/ZnJELE3.png)
        
    - [http://localhost:8080/form](http://localhost:8080/form)
        
        ![](https://i.imgur.com/CthQt3e.png)