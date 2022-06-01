package main

import (
	"fmt"
)

//Consumer
type Consumer struct {
	items *chan int
}

//creating a new Consumer
func NewConsumer(items *chan int) *Consumer {
	return &Consumer{items: items}
}

//consuming items from the channel
func (c *Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		item := <-*c.items
		fmt.Println("consume: Received:", item)
	}
}

//Producer
type Producer struct {
	items *chan int
	done  *chan bool
}

//creating a new Producer
func NewProducer(items *chan int, done *chan bool) *Producer {
	return &Producer{items: items, done: done}
}

//producing and sending the item through the channel
func (p *Producer) produce(max int) {
	fmt.Println("produce: Started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		*p.items <- i
	}
	*p.done <- true // signal when done
	fmt.Println("produce: Done")
}

func main() {

	var items = make(chan int) // channel to send items
	var done = make(chan bool) // channel to control when production is done

	// Start a goroutine for Produce.produce
	go NewProducer(&items, &done).produce(5)

	// Start a goroutine for Consumer.consume
	go NewConsumer(&items).consume()

	// Finish the program when the production is done
	<-done

}
