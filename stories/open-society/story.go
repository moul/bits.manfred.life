package main

import "fmt"

// Defining the structures of the society
type Location struct {
	Name string
}

type Person struct {
	Name   string
	Age    int
	Design string
}

type Event struct {
	Name        string
	Location    Location
	Participants []Person
}

// Functions to describe certain actions in the story
func (p *Person) visitLocation(l Location) {
	fmt.Printf("%s is now at %s.\n", p.Name, l.Name)
}

func (p *Person) contributeToEvent(e Event) {
	fmt.Printf("%s is participating in the event: %s at %s.\n", p.Name, e.Name, e.Location.Name)
}

func main() {
	// Locations in the story
	grandAgora := Location{Name: "Grand Agora"}
	diningCommons := Location{Name: "Dining Commons"}
	localLibrary := Location{Name: "Local Library"}
	observatory := Location{Name: "Observatory"}

	// Main character
	lia := Person{Name: "Lia", Age: 28}

	// Describing the journey using the functions
	lia.visitLocation(grandAgora)
	lia.visitLocation(diningCommons)
	lia.visitLocation(localLibrary)
	lia.visitLocation(observatory)

	// An example event
	artEvent := Event{
		Name:        "Open Art Installation",
		Location:    grandAgora,
		Participants: []Person{lia},
	}

	lia.contributeToEvent(artEvent)
}
