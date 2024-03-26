package handler

import (
	"encoding/json"
	"fanout/internal/rabbitmq"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var consumerCount int = 0

func StartHttpServer() {

	log.Println("Http server started")

	router := mux.NewRouter()

	router.HandleFunc("/", nil)
	router.HandleFunc("/publish/{data}", publishMessage).Methods(http.MethodGet)
	router.HandleFunc("/consume", consumeMessage).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("http server is not started error: ", err)
	}
}

func publishMessage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	data := vars["data"]

	rabbitmq.PublishWithQueue("public-queue", data)

	err := json.NewEncoder(w).Encode("message send: " + string(data))
	if err != nil {
		log.Println("error encoding message: ", err)
	}
}

func consumeMessage(w http.ResponseWriter, r *http.Request) {

	consumerCount++

	log.Printf("consumer %v is started. \n", consumerCount)

	rabbitmq.Consume(consumerCount, "public-queue")
}
