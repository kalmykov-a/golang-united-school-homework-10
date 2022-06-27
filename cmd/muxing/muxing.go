package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", HandleParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", HandleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", HandleData).Methods(http.MethodPost)
	router.HandleFunc("/header", HandleHeader).Methods(http.MethodGet)
	router.HandleFunc("/", HandleDefault)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func HandleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func HandleParam(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func HandleData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "I got message:\n%s", string(body))
}

func HandleHeader(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Set("a+b", strconv.Itoa(a+b))
}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
