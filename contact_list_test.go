package main

import (
	"testing"
)

var cl ContactList = *NewContactList()

func TestContactList_AddContact(t *testing.T) {
	testCases := []struct {
		name     string
		contact  Contact
		expected error
	}{
		{"John", Contact{Name: "John", Email: "john@example.com", Phone: "1234567890"}, nil},
		{"Jane", Contact{Name: "Jane", Email: "jane@example.com", Phone: "9876543210"}, nil},
		{"Duplicate", Contact{Name: "John", Email: "newjohn@example.com", Phone: "5555555555"}, ErrContactExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cl.AddContact(tc.contact)
			if err != tc.expected {
				t.Errorf("Expected error: %v, got error: %v", tc.expected, err)
			}
		})
	}
}

func TestContactList_GetContact(t *testing.T) {
	testCases := []struct {
		name     string
		getName  string
		expected Contact
		err      error
	}{
		{"John", "John", Contact{Name: "John", Email: "john@example.com", Phone: "1234567890"}, nil},
		{"Jane", "Jane", Contact{Name: "Jane", Email: "jane@example.com", Phone: "9876543210"}, nil},
		{"NotFound", "NonExistent", Contact{}, ErrContactNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contact, err := cl.GetContact(tc.getName)
			if err != tc.err {
				t.Errorf("Expected error: %v, got error: %v", tc.err, err)
			}
			if contact != tc.expected {
				t.Errorf("Expected contact: %+v, got contact: %+v", tc.expected, contact)
			}
		})
	}
}

func TestContactList_UpdateContact(t *testing.T) {
	testCases := []struct {
		name           string
		originalName   string
		updatedContact Contact
		expectedError  error
	}{
		{"John", "John", Contact{Name: "John", Email: "newjohn@example.com", Phone: "5555555555"}, nil},
		{"NotFound", "NonExistent", Contact{}, ErrContactNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cl.UpdateContact(tc.originalName, tc.updatedContact)
			if err != tc.expectedError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectedError, err)
			}
		})
	}
}

func TestContactList_DeleteContact(t *testing.T) {
	testCases := []struct {
		name          string
		deleteName    string
		expectedError error
	}{
		{"John", "John", nil},
		{"NotFound", "NonExistent", ErrContactNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cl.DeleteContact(tc.deleteName)
			if err != tc.expectedError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectedError, err)
			}
		})
	}
}
