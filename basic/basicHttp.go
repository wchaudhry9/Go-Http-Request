package main

import (
	"fmt"
	"log"
	"net/http"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Starting up server...")
	http.HandleFunc("/", customDisplayText) // Root url will try getting user input
	log.Fatal(http.ListenAndServe(":8080", nil)) // Listen on port 8080
}

func customDisplayText(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Enter in some phrase or characters: ")
	scan := bufio.NewReader(os.Stdin)
	phrase, _ := scan.ReadString('\n') 
	fmt.Printf(phrase) // print to command line
	fmt.Fprint(w, phrase) // print to HTML
}
