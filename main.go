package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"crypto/sha256"
	"fmt"
	"github.com/gorilla/mux"
)

type Message struct {
	Content string `json:"message"`
}

type Response struct {
	Digest string `json:"digest"`
}

func main() {
	h := mux.NewRouter()
	//var m map[string]string
	h.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		//switch r.Method {
		if r.Method == http.MethodPost {
			//log.Print(PostHash(w, r))
			response := Response{Digest: PostHash(w, r)}
			json.NewEncoder(w).Encode(response)
			//break
		}
	})
	h.HandleFunc("/messages/{hash}", func(w http.ResponseWriter, r *http.Request) {
		//switch r.Method {
		if r.Method == http.MethodGet {
			log.Print(GetHash(w, r))
			//break			
		}
	})
	log.Fatal(http.ListenAndServeTLS(":5000", "localhost.crt", "localhost.key", h))
}

func GetHash(w http.ResponseWriter, r *http.Request) string{
	defer r.Body.Close()

	vars := mux.Vars(r)
	hash := vars["hash"]

	return hash
}

func PostHash(w http.ResponseWriter, r *http.Request) string{
	defer r.Body.Close()

	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	var message Message
	json.Unmarshal(byteData, &message)
	sum := sha256.Sum256([]byte(message.Content))
	return fmt.Sprintf("%x", sum)
	//return sha256.Sum256([]byte(message.Content))
}