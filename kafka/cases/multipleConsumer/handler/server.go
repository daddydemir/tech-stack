package handler

import (
	"github.com/gorilla/mux"
	"log"
	"multiple/internal/kafka"
	"net/http"
)

var consumerCount int = 1

func StartHttpServer() {

	router := mux.NewRouter()

	router.HandleFunc("/", nil)
	router.HandleFunc("/publish/{data}", publisherHandler)
	router.HandleFunc("/consume", consumerHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func consumerHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Conusmer starting...: %v", consumerCount)
	//consumer := kafka.NewConsumer(fmt.Sprint("cg-", consumerCount))
	consumer := kafka.NewConsumer("cg-55")
	consumerCount++
	consumer.ListenMessages()
}

func publisherHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	data := vars["data"]

	producer := kafka.NewProducer()
	producer.SendMessage(data)

}
