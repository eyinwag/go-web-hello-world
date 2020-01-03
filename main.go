package main

import (
     "net/http"
     "fmt"
     "log"
)

func main() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
           fmt.Fprint(writer, "Hello, world")
     })
     log.Fatal(http.ListenAndServe(":3100", nil))
}
