package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
  http.HandleFunc("/", Hello)
  http.HandleFunc("/", ConfigMap)
  http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
  name := os.Getenv("NAME")
  age := os.Getenv("AGE")

  fmt.Fprintf(w ,"Hello I'm %s. I'm %s", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
  
  data, err := io.ReadFile("myfamily/family.txt")

  fmt.Printf(w, "Hello, I'm %s. I'm %s", name, age)
}
