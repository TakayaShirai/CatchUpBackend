package main

import (
	"fmt"
	"net/http"
)

func main() {
	// This line registers a handler function for the path / (i.e. the root URL).
	// When someone accesses http://localhost:8080/, this function is triggered.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, ...) writes to the response being sent to a browser or HTTP client.
		// Go automatically:
		// 1. Opens a connection to the client
		// 2. Writes "Hello, client!" into the response body
		// 3. Sends the full HTTP response
		// 4. Closes the connection (unless Keep-Alive is on)
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Number of bytes written:", n)
	})

	http.ListenAndServe(":8080", nil)
}
