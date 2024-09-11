package main

import (
	"fmt"
	"log"
	"net/http"
)

func flightHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("flightHandler Method: ", r.Method)

	raw := `{
			"name": "phumiphat",
			"age": 23
		}`

	w.Write([]byte(raw))

	return
}

func main() {
	http.HandleFunc("GET /flights", flightHandler)

	log.Println("Start server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
