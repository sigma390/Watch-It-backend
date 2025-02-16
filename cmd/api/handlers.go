package main

import (
	"fmt"
	"net/http"
)

//handler Function

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow")
}

// new Route Handler
func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page , hello from %s", app.Domain)

}
