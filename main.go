package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var linkList map[string]string

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	linkList = map[string]string{}

	http.HandleFunc("/addlink", addLink)
	http.HandleFunc("/shorted", getLink)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addLink(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()["link"]
	if ok {
		if _, ok := linkList[key[0]]; !ok {
			generateString := fmt.Sprint(rand.Int63n(1000))
			linkList[generateString] = key[0]
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusAccepted)
			linkString := fmt.Sprintf("<a href=\"http://localhost:8080/short/%s\">http://localhost:8080/short/%s</a>", generateString, generateString)
			fmt.Fprintf(w, "added shortlink\n")
			fmt.Fprintf(w, linkString)
			return
		}
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Already have this link")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Failed to add link")
	return
}
