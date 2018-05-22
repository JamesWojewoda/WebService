package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"crypto/sha256"
	"fmt"
	"os"
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
	f, err := os.OpenFile("/var/log/go-service/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	h := mux.NewRouter()
	m := make(map[string]string)
	h.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			response := Response{Digest: PostHash(w, r, m)}
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(response)
			log.Println("Post request 201: " + response.Digest)
		}
	})
	h.HandleFunc("/messages/{hash}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getresult := GetHash(w, r, m)
			if getresult != "" {
				response := Message{Content: getresult}
				json.NewEncoder(w).Encode(response)
				log.Println("Get request 200: " + getresult)
			} else {
				response := Error{Error: "Not Found"}
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(response)
				log.Println("Get request 404")
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