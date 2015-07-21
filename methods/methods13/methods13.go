package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

/* Hello implements http.Handler
   package http

   type Handler interface {
       ServeHTTP(w ResponseWriter, r *Request)
   }
*/
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
