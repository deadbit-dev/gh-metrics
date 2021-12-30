package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome, please GET git name!\n")
}

func Data(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	nickname := ps.ByName("name")
	endpoint := "https://api.github.com/users/" + nickname + "/repos?type=owner"
	log.Print("INFO: " + endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: parse json response
	// TODO: loop request by languages

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/:name", Data)

	log.Fatal(http.ListenAndServe(":8000", router))
}
