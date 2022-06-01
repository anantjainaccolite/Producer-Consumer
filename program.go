package main

import (
	"fmt"
	"time"
)

//Consumer
type Consumer struct {
	items     *chan int
	timestamp *chan string
}

//creating a new Consumer
func NewConsumer(items *chan int, timestamp *chan string) *Consumer {
	return &Consumer{items: items, timestamp: timestamp}
}

//consuming items from the channel
func (c *Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		// time.Sleep(45 * time.Millisecond)
		if len(*c.items) <= 0 {
			continue
		}
		fmt.Println("consuming...")
		item := <-*c.items
		timestamps := <-*c.timestamp
		fmt.Println("Received:", item, " which was produced at: ", timestamps)
	}
}

//Producer
type Producer struct {
	items     *chan int
	done      *chan bool
	timestamp *chan string
}

//creating a new Producer
func NewProducer(items *chan int, done *chan bool, timestamp *chan string) *Producer {
	return &Producer{items: items, done: done, timestamp: timestamp}
}

//producing and sending the item through the channel
func (p *Producer) produce(max int) {
	i := 1
	fmt.Println("produce: Started")
	for {

		if len(*p.items) >= max {
			continue
		}
		fmt.Println("starting.....")
		*p.items <- i
		fmt.Println("starting 1.....")
		*p.timestamp <- time.Now().String()
		fmt.Println("starting 2.....")
		fmt.Println("produce: Sending ", i)
		i++
		// time.Sleep(30 * time.Millisecond)
	}
	// *p.done <- true // signal when done
	// fmt.Println("produce: Done")
}

func main() {

	var items = make(chan int)        // channel to send items
	var done = make(chan bool)        // channel to control when production is done
	var timestamp = make(chan string) //channel to  add timestamps

	// Start a goroutine for Produce.produce
	len := 5
	go NewProducer(&items, &done, &timestamp).produce(len)
	fmt.Println("in betwenn")
	// Start a goroutine for Consumer.consume
	go NewConsumer(&items, &timestamp).consume()

	time.Sleep(100 * time.Second)
	// Finish the program when the production is done
	fmt.Println("last")
	// <-done

}
