package main

import "errors"

// var of errors
var (
	ErrContactNotFound = errors.New("contact not found")
	ErrContactExists   = errors.New("contact already exists")
)

/*
	 Hint : Expected output list contactd (map[string]Contact) looks like this:
	    [
		'asnur ramdani' : {Name: 'asnur ramdani', Email: 'asnurramdhani12@gmail.com', Phone: '085156156156'},
		'joko' : {Name: 'joko', Email: 'joko@gmail.com', Phone: '085156156156'},
		...
		]
*/
type Contact struct {
	Name  string
	Email string
	Phone string
}

type ContactList struct {
	contacts map[string]Contact
}

func NewContactList() *ContactList {
	return &ContactList{contacts: make(map[string]Contact)}
}

func (cl *ContactList) AddContact(contact Contact) error {
	// TODO 1: Add a new contact to the contact list. Return an error if the contact already exists.
	return nil
}

func (cl *ContactList) GetContact(name string) (Contact, error) {
	// TODO 2 : Retrieve a contact by their name. Return an error if the contact doesn't exist.
	return Contact{}, nil
}

func (cl *ContactList) UpdateContact(name string, updatedContact Contact) error {
	// TODO 3: Update the details of an existing contact. Return an error if the contact doesn't exist.
	return nil
}

func (cl *ContactList) DeleteContact(name string) error {
	// TODO 4: Delete a contact from the contact list. Return an error if the contact doesn't exist.
	return nil
}
