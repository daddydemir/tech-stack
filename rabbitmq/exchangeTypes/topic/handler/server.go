package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"topic/internal/rabbitmq"
)

var consumerCount int = 0

func StartHttpServer() {

	log.Println("Http server started")

	router := mux.NewRouter()

	router.HandleFunc("/", nil)
	router.HandleFunc("/publish/{data}", publishMessage).Methods(http.MethodGet)
	router.HandleFunc("/consume/{data}", consumeMessage).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("http server is not started error: ", err)
	}

}

func publishMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	data := vars["data"]

	rabbitmq.PublishWithExchange("topic-1", data)

	err := json.NewEncoder(w).Encode("message send: " + string(data))
	if err != nil {
		log.Println("error encoding message: ", err)
	}
}

func consumeMessage(w http.ResponseWriter, r *http.Request) {

	consumerCount++

	vars := mux.Vars(r)
	data := vars["data"]

	log.Println("consumer count: ", consumerCount)

	rabbitmq.Consume(consumerCount, data)
}
