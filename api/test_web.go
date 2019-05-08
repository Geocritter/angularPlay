
package main

import (
	// "fmt"
	"io"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	// "github.com/gorilla/csrf"
)

type Oldman struct {
	Age 	uint8		`json:"age"`
	Hobbies	[]string	`json:"hobbies"`
}
type Oldmen []Oldman

var oldmen Oldmen

func main() {
	o1 := Oldman{Age: 77, Hobbies: []string{"music", "golf"}}
	o2 := Oldman{Age: 97, Hobbies: []string{"playdead",}}
	oldmen = Oldmen{o1, o2}
/*
	CSRF := csrf.Protect([]byte("ipcoredevops"),
		csrf.CookieName("Csrf-Token"),
		csrf.RequestHeader("X-Csrf-Token"),
		// Domain should be as default "/"
		//csrf.Domain("127.0.0.1"),
		csrf.HttpOnly(false),
		csrf.Secure(false),
	)
*/
	r := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{`http://localhost:4200`,}),
		//handlers.AllowedHeaders([]string{"content-type", "X-Csrf-Token", "withcredentials", "credentials",}),
		handlers.AllowedHeaders([]string{"content-type",}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET","POST","OPTIONS","HEAD",}),
	)
	r.HandleFunc("/index", list).Methods("GET")
	r.HandleFunc("/index", add).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", cors(r)))
}


func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Header().Set("X-Csrf-Token", csrf.Token(r))
    json.NewEncoder(w).Encode(oldmen)
}

func add(w http.ResponseWriter, r *http.Request) {

	var oldman Oldman
	// maximum 1M bytes
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
	}
	if err := json.Unmarshal(body, &oldman); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
 
    oldmen = append(oldmen, oldman)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Header().Set("X-Csrf-Token", csrf.Token(r))
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(oldmen); err != nil {
        panic(err)
	}
	
}