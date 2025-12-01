package main

import "fmt"

type FormElement struct {
	Type     string   `json:"type"`              // I.e., "text", "number", "select", etc.
	Label    string   `json:"label"`             // i.e., "First Name", "Age", etc.
	Name     string   `json:"name"`              // i.e., "first_name", "age", etc.
	Options  []string `json:"options,omitempty"` // Used for select, radio, etc.
	Required bool     `json:"required"`          // Whether the field is required
}

type Service interface {
	GetForm(formID string) ([]FormElement, error)
}

type formService struct{}

func (formService) GetForm(formID string) ([]FormElement, error) {
	forms := map[string][]FormElement{
		"registration": {
			{Type: "text", Label: "First Name", Name: "first_name", Required: true},
			{Type: "text", Label: "Last Name", Name: "last_name", Required: true},
			{Type: "email", Label: "Email", Name: "email", Required: true},
			{Type: "phone_number", Label: "Phone Number", Name: "phone_number", Required: true},
		},
		"feedback": {
			{Type: "text", Label: "Name", Name: "name", Required: false},
			{Type: "email", Label: "Email", Name: "email", Required: false},
		},
	}

	elements, exists := forms[formID]
	if !exists {
		return nil, fmt.Errorf("form with ID %s not found", formID)
	}
	return elements, nil
}

func NewService() Service {
	return formService{}
}
