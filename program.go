package main

import (
	"fmt"
)

//Consumer
type Consumer struct {
	items *chan int
	timestamp *chan string
}

//creating a new Consumer
func NewConsumer(items *chan int) *Consumer {
	return &Consumer{items: items, timestamp: timestamp}
}

//consuming items from the channel
func (c *Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		item := <-*c.items
		timestamps := <-*c.timestamp
		fmt.Println("Received:", item, " which was produced at: ",timestamps)
	}
}

//Producer
type Producer struct {
	items *chan int
	done  *chan bool
	timestamp *chan string
}

//creating a new Producer
func NewProducer(items *chan int, done *chan bool, timestamp *chan string) *Producer {
	return &Producer{items: items, done: done, timestamp: timestamp}
}

//producing and sending the item through the channel
func (p *Producer) produce(max int) {
	fmt.Println("produce: Started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		*p.items <- i
		*p.timestamp <- string(time.crrentTime)
	}
	*p.done <- true // signal when done
	fmt.Println("produce: Done")
}

func main() {

	var items = make(chan int) // channel to send items
	var done = make(chan bool) // channel to control when production is done
	var timestamp = make(chan string)
	// Start a goroutine for Produce.produce
	go NewProducer(&items, &done, &timestamp).produce(5)

	// Start a goroutine for Consumer.consume
	go NewConsumer(&items, &timestamp).consume()

	// Finish the program when the production is done
	<-done

}
