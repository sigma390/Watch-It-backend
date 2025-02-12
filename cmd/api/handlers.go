package main

import (
	"fmt"
	"net/http"
)

//handler Function

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow")
}
