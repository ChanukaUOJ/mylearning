package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {

	topic := "coffee_orders"
	msgcount := 0

	// 1. Create a new consumer and start it
	worker, err := connectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	// 2. Handle OS signals - used to stop the process
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 3. Create a Goroutine to consume messages
	donechan := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println("Error from consumer:", err)
			case msg := <-consumer.Messages():
				msgcount++
				log.Printf("Message claimed: %s\n", string(msg.Value))
				order := string(msg.Value)
				fmt.Printf("Brewing coffee for order : %s\n", order)
			case <-sigchan:
				log.Println("Interrupted")
				donechan <- struct{}{}
			}
		}
	}()

	<-donechan
	fmt.Printf("Processed %d messages\n", msgcount)

	// 4. Close the consumer on exit
	if err := worker.Close(); err != nil {
		log.Println("Error closing consumer:", err)
	}

}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}