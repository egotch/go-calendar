// contact keeping module

//TODO - method to print single contact
//TODO - method to add contact
//TODO - method to delete contact (by name)
//TODO - function to generate/return contact list

package contacts

import "fmt"

type contact struct {
	name  string
	email string
	phone string
}

// printContact prints the contact information for the contact
func (c *contact) PrintContact() {
	fmt.Printf("Name: %s\nEmail: %s\nPhone:%s\n\n\n", c.name, c.email, c.phone)
}

func (c *contact) UpdateContact(email string, phone string) {

	c.email = email
	c.phone = phone
}

// PrintContacts prints contact information of all contacts
// func PrintContacts(*map[string]contact) {
// 	fmt.Printf("Your Contact List...\n\n")
// 	for _, c := range map {
// 		c.printContact()
// 	}
// }

// InitContact list sets up the main contact list
func InitContactList() map[string]contact {
	var c_list = make(map[string]contact)

	c_list["Eric"] = contact{"Eric", "egotkowski@gmail.com", "6302538183"}
	c_list["Emily"] = contact{"Emily", "emilywinans@gmail.com", "2182052998"}
	c_list["Lucy"] = contact{"Lucy", "lucy@gmail.com", "1234569879"}
	c_list["Kylie"] = contact{"Kylie", "kylie@gmail.com", "9871235467"}

	return c_list
}

// func UpdateContact(name string, email string, phone string) {
// 	c := c_list[name]
// 	c.updateContact(email, phone)
// 	fmt.Printf(c.email)
// }

// func AddContact() {
// 	c_list["Eric2"] = contact{"Eric2", "shklumpo@hotmail.com", "6302538183"}
// }
