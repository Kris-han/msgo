package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%s 欢迎来到马神之路web教程", "krishan")
	})
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		log.Fatal(err)
	}
}
