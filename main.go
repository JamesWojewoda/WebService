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
type Error struct {
	Error string `json:"err_msg"`
}
func main() {
	h := mux.NewRouter()
	m := make(map[string]string)
	h.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			response := Response{Digest: PostHash(w, r, m)}
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(response)
		}
	})
	h.HandleFunc("/messages/{hash}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Print(GetHash(w, r, m))
			if GetHash(w, r, m) != "" {
				response := Message{Content: GetHash(w, r, m)}
				json.NewEncoder(w).Encode(response)
			} else {
				response := Error{Error: "Not Found"}
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(response)
			}
			
		}
	})
	log.Fatal(http.ListenAndServeTLS(":5000", "/etc/ssl/certs/localhost.crt", "/etc/ssl/certs/localhost.key", h))
}

func GetHash(w http.ResponseWriter, r *http.Request, m map[string]string) string{
	defer r.Body.Close()

	vars := mux.Vars(r)
	hash := vars["hash"]

	return m[hash]
}

func PostHash(w http.ResponseWriter, r *http.Request, m map[string]string) string{
	defer r.Body.Close()

	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	var message Message
	json.Unmarshal(byteData, &message)
	sum := fmt.Sprintf("%x", sha256.Sum256([]byte(message.Content)))
	m[sum] = message.Content
	return sum
}