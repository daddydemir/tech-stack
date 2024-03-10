package handler

import (
	"direct/internal/rabbitmq"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var consumerCount int = 0

func StartHttpServer() {

	router := mux.NewRouter()

	router.HandleFunc("/", nil)
	router.HandleFunc("/publish/{data}", publishMessage)
	router.HandleFunc("/publishQueue/{data}", publishToQueue)
	router.HandleFunc("/consume", consumeMessages)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("http server is not started error: ", err)
	}

}

func publishMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	data := vars["data"]

	rabbitmq.PublishWithExchange("personal-2", "", data)

	err := json.NewEncoder(w).Encode("message send: " + string(data))
	if err != nil {
		log.Println("error encoding message: ", err)
	}
}

func publishToQueue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	data := vars["data"]

	rabbitmq.PublishWithQueue("queue", data)

	err := json.NewEncoder(w).Encode("message send: " + string(data))
	if err != nil {
		log.Println("error encoding message: ", err)
	}
}

func consumeMessages(w http.ResponseWriter, r *http.Request) {

	consumerCount++

	log.Printf("consumer %v is started. \n", consumerCount)

	rabbitmq.Consume(consumerCount, "queue")

}
