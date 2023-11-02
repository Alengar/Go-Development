package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("http://localhost:8081/data")
		if err != nil {
			http.Error(w, "Error requesting data from Service 2", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		data := make([]byte, 100)
		_, err = response.Body.Read(data)
		if err != nil {
			http.Error(w, "Error reading response from Service 2", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(data))
	})

	http.ListenAndServe(":8080", nil)
}
