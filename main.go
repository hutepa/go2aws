package main

import (
  "log"
  "net/http"
  "os"
)

func main() {
  
  port := os.Getenv("PORT")
  if port == "" {
    port = "80"
  }
  const index = "index.html"
  http.HnadleFunc("/", fun(w http.ResponseWriter, r *http.Request"){
      http.ServeFile(w,r,index)
  })
  log.Printf("started webserver on port %s\n\n", port)
  http.ListenAndServe(":"+port, nil)

}                  
