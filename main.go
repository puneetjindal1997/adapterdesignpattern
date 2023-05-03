package main

import "fmt"

// Welcome to your channel go guruji

// What is adaptor design pattern

// Why we use it

// Implementation of it

type Boat struct{}

func (w *Boat) travelToDestination() {
	fmt.Println("boat is navigating to destination")
}

type Client struct{}

func (c *Client) StartingMyJourney(com Transportation) {
	fmt.Println("Starting navigation process")
	com.Travel()
}

type Transportation interface {
	Travel()
}

type Car struct{}

func (m *Car) Travel() {
	fmt.Println("car is navigating to destination")
}

type BoatAdapter struct {
	boat *Boat
}

func (w *BoatAdapter) Travel() {
	fmt.Println("Adaptor used to move boat on roads")
	w.boat.travelToDestination()
}

func main() {

	client := &Client{}
	car := &Car{}

	fmt.Println("car started")
	client.StartingMyJourney(car)

	fmt.Println("boat started")
	boat := &Boat{}
	boatAdaptor := &BoatAdapter{
		boat: boat,
	}

	client.StartingMyJourney(boatAdaptor)

}
