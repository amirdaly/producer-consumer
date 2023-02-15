package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// Producer is a type for structs that holds tow chammels: one for pizzas, with all
// information for a given pizza order including whether it  was made
// successfully, and another to handle end of processing (when quit the channel)
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder is a type for structs that describes a given pizza order. It has the order
// number, a message indicating what happened to the order, and a boolean
// indicating if the order was successfully completed.
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is simply a method of closing the channel when we are done it (i.e)
// something is pushed to the quit channel)
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// makePizza attempts to make pizza. We generate a reandom number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza whithout issue. To make things intresting, each pizza
// will take a different length ig time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quite while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	i := 0

	// run forever or until we recive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		// try to make a pizza

		// decision
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out the ending message
}
