package main

import (
    "fmt"
    "net/http"
)

func main() {
    helloWorld()
}

func helloWorld() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "<h1>hello world %v</h1>", request.FormValue("name"))
    })
    http.ListenAndServe(":80", nil)
}

