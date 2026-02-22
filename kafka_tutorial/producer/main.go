package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	CoffeeType   string `json:"coffee_type"`
}

func main() {
	http.HandleFunc("/order", placeOrder)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func connectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func pushOrderToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}

	producer, err := connectProducer(brokers)
	if err != nil {
		return err
	}
	defer producer.Close()

	// create a new message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// send message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

// placeOrder handler
func placeOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parse request body into order
	order := new(Order)

	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		log.Println("Error parsing order:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Convert body into bytes
	orderInBytes, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Send to Kafka
	err = pushOrderToQueue("coffee_orders", orderInBytes)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 4. Send response to client
	response := map[string]interface{}{
		"success": true,
		"msg":     "Order for " + order.CustomerName + " placed successfully!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		http.Error(w, "Error placing order", http.StatusInternalServerError)
		return
	}
}