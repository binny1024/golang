package main

import (
	"io"
	"log"
	"net/http"
)

/*
	类似一个回调函数的实现,func(ResponseWriter, *Request)
 */
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "C语言中文网\n")
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}