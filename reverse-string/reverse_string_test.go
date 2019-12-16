package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	actualResult := Reverse("Hello")
	expectedResult := "olleH"

	if expectedResult != actualResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
