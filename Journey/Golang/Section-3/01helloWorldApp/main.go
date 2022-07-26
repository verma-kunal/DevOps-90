package main

import (
	"fmt"
	"net/http"
)

func main() {

	// creating a server using 'http' package:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")

		// as the variables in Go when declared cannot be un-used, therefore we'll use them somewhere:
		if err != nil {
			fmt.Println(err.Error())
		}

		// "Sprintf" -> converts any statement into a string type:
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// listening to the server: (without handling any errors (used '_' for that))
	_ = http.ListenAndServe("localhost:8080", nil)

}
