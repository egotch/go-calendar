// Maintains a csv file that servers as a local celanard storage
// Takes in a CLI argument for date, time, place
// Formats the input and saves it to the running csv file
// Also can print out the contents of the csv file to "display"
// the calendar
// TODO - initiate contact list
// TODO - add some contacts to the contact list & print out the contact list
// TODO - remove some contacts from contact list & print contacts
package main

import (
	"github.com/contacts"
)

// globals for holding months, days of week
//globals for holding known holidays

func main() {
	// Initiate contact list
	c_list := contacts.InitContactList()
	for _, contact := range c_list {
		contact.PrintContact()
	}

}

// function to call init for contacts list and print out current contacts
func startContacts() {

}
