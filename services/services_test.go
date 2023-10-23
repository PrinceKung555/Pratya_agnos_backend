package services_test

import (
	services "password_strength_api/services"
	"testing"
)

func TestIsStrong(t *testing.T) {
	tests := []struct {
		password     string
		expectStrong bool
	}{
		{"aA1", false},
		{"1445D1cd", true},
		{"abc123", false},
		{"Aa1Bb2Ccc3Ddd4Eeee5Fffff6", false},
		{"abc123!@#", false},
		{"a", false},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			actual := services.IsStrong(test.password)
			if actual != test.expectStrong {
				t.Errorf("Expected isStrong(%s) to be %v, but got %v", test.password, test.expectStrong, actual)
			}
		})
	}
}

func TestActionsNeededToMakeStrong(t *testing.T) {
	tests := []struct {
		password      string
		expectedSteps int
	}{
		{"aA1", 3},
		{"1445D1cd", 0},
		{"abc123", 1},
		{"Aa1Bb2Ccc3Ddd4Eeee5Fffff6", 5},
		{"abc123!@#", 1},
		{"a", 5},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			actual := services.ActionsNeededToMakeStrong(test.password)
			if actual != test.expectedSteps {
				t.Errorf("Expected actionsNeededToMakeStrong(%s) to be %v, but got %v", test.password, test.expectedSteps, actual)
			}
		})
	}
}
