package main

import (
    "fmt"
    "log"
    "net/http"
)
 
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
 
func main() {
	fmt.Println("Started server at http://localhost:8000")
    handleRequests()
}