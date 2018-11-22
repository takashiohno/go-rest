package main

import (
    "net/http"

    //"github.com/takashiohno/go-rest/logs"
    _ "github.com/takashiohno/go-rest/resources"
)

func response(rw http.ResponseWriter, request *http.Request) {
    rw.Write([]byte("Hello world go"))
}

func main() {
    //http.HandleFunc("/", response)

    http.ListenAndServe(":3000", nil)

    //logs.Info.Println("Hello zap")
    //logs.ZapLogger.Info("Hello zap")
}
