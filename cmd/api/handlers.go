package main

import (
	"fmt"
	"log"
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

func (app *application) Statuss(w http.ResponseWriter, r *http.Request) {
	//kya send krna hai Frontend ko
	var payload = struct {
		Status  string `json:"status"`
		Version string `json:"version"`
		Message string `json:"msg"`
	}{
		Status:  "success",
		Version: "1.0.0",
		Message: "API is running fine",
	}
	//convert to json using Marshal

	_ = app.writeJSON(w, http.StatusOK, payload)
}

//movie Interface

func (app *application) Movies(w http.ResponseWriter, r *http.Request) {

	//get all movies
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err) //use errorJSON function
		return
	}

	// //convert to json
	// out, err := json.Marshal(movies)
	// if err != nil {
	// 	log.Println(err)
	// }

	// //set headers
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// //send json
	// w.Write(out)
	_ = app.writeJSON(w, http.StatusOK, movies)

}

func (app *application) Movie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {

	//read Request Payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//decode json
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	//get user By Email
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	//check Password and User exists Oor not

	//USER
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	log.Println(tokens)
	//get Refresh Cookie
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	w.Write([]byte(tokens.Token))

}
