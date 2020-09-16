package main

import(
	"fmt"
	"log"
	"net/http"
)
func main(){
	fmt.Println("Hello go")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello go")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}