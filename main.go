package main

import (
	"bufio"
	"fmt"
	"os"
)

// Options state
const (
	INITIAL = -1
	EXIT    = 0
	DISPLAY = 1
	SEARCH  = 2
	ADD     = 3
	REMOVE  = 4
)

type option int

type id string

type person struct {
	name  string
	email string
	phone string
}

func (p person) display() {
	fmt.Printf("Name: %s \n", p.name)
	fmt.Printf("Email: %s \n", p.email)
	fmt.Printf("Phone: %s \n", p.phone)
}

var phonebook = map[id]person{}

func main() {
	var input option = INITIAL
	for input != 0 {
		displayMenu()

		_, err := fmt.Scanln(&input)

		if err != nil {
			input = INITIAL
			fmt.Println("---------------------------------------")
			fmt.Println("Invalid input! Please enter a valid option.")
			fmt.Println("---------------------------------------")
		}

		dispatch(input)
	}
}

func displayMenu() {
	fmt.Println("---------------------------------------")
	fmt.Println("Hello Phonebook")
	fmt.Println("---------------------------------------")
	fmt.Println("Select the following options:")
	fmt.Println("0 - Exit")
	fmt.Println("1 - Display Phonebook")
	fmt.Println("2 - Search for contact")
	fmt.Println("3 - Add new contact")
	fmt.Println("4 - Delete new contact")
	fmt.Println("---------------------------------------")
}

func displayPhonebook(phonebook map[id]person) {
	for id, person := range phonebook {
		fmt.Printf("ID: %s \n", id)
		person.display()
		fmt.Println("---------------------------------------")
	}
}

func dispatch(option option) {
	switch option {
	case DISPLAY:
		displayPhonebook(phonebook)
	case SEARCH:
		searchContact(phonebook)
	case ADD:
		addContact(phonebook)
	case REMOVE:
		removeContact(phonebook)
	}
}

func addContact(phonebook map[id]person) {
	p := person{}

	fmt.Println("---------------------------------------")
	fmt.Println("Let's add a new contact!")
	fmt.Println("---------------------------------------")

	fmt.Println("What's the contact name?")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		p.name = scanner.Text()
	}
	fmt.Println("---------------------------------------")

	fmt.Println("What's the contact email?")
	fmt.Scanln(&p.email)
	fmt.Println("---------------------------------------")

	fmt.Println("What's the contact phone?")
	fmt.Scanln(&p.phone)
	fmt.Println("---------------------------------------")

	phonebook[id(p.name)] = p

	fmt.Println("Added!")
	p.display()
}

func searchContact(phonebook map[id]person) {
	fmt.Println("Let's search for a contact!")
	fmt.Println("---------------------------------------")

	var key id
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		key = id(scanner.Text())
	}

	if person, ok := phonebook[key]; ok == true {
		person.display()
	} else {
		fmt.Println("Contact not found")
	}
}

func removeContact(phonebook map[id]person) {
	fmt.Println("Let's remove a contact!")
	fmt.Println("---------------------------------------")
	fmt.Println("Enter contact name:")

	var key id
	fmt.Scanln(&key)

	if person, ok := phonebook[key]; ok == true {
		fmt.Println("Removing:")
		person.display()
		delete(phonebook, key)
	}
}
