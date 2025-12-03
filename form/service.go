package form

import "fmt"

// FormElement describes a single form input field. Similar to a interface in js/ts. No type inheritance.
type FormElement struct {
	Type     string   `json:"type"`              // Input type: text, email, number, select, etc.
	Label    string   `json:"label"`             // Display name
	Name     string   `json:"name"`              // Form field name
	Options  []string `json:"options,omitempty"` // Choices for select/radio fields
	Required bool     `json:"required"`          // Is field required?
}

// Service retrieves form definitions. Interface allows easy testing.
type Service interface {
	GetForm(formID string) ([]FormElement, error)
}

type formService struct{}

// GetForm returns form elements or error if not found.
func (formService) GetForm(formID string) ([]FormElement, error) {
	// Hardcoded forms. In production, load from database
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
		"survey": {
			{Type: "text", Label: "Age", Name: "age", Required: true},
			{Type: "select", Label: "Satisfaction Level", Name: "satisfaction_level", Options: []string{"Very Satisfied", "Satisfied", "Neutral", "Dissatisfied", "Very Dissatisfied"}, Required: true},
			{Type: "text", Label: "Comments", Name: "comments", Required: false},
		},
	}

	elements, exists := forms[formID]
	if !exists {
		return nil, fmt.Errorf("form with ID %s not found", formID)
	}
	return elements, nil
}

// NewService factory creates a Service instance.
func NewService() Service {
	return formService{}
}
