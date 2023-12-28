package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops",http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n",data)

		fmt.Fprintf(w,"Hello %s", data)

	})	

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		fmt.Println("goodbye")
		log.Println("goodbye")
	})	
	http.ListenAndServe(":9090",nil)
}