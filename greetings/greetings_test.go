package greetings

import (
	"regexp"
	"testing"
)

//TestHelloName calls greetings.Hello with a name,
//checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions - With Update" {
		t.Errorf("Greet() = %s; want Hello GitHub actions", result)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Bisola", "Bunmi", "Fatmat", "Ngozi"}
	messages, err := Hellos(names)
	if err != nil {
		t.Error("Hellos() creates error")
	}
	if len(messages) < 3 {
		t.Error("No message generated.")
	}
}

func TestHellosWithEmptyName(t *testing.T) {
	names := []string{"", "", "Fatmat", "Ngozi"}
	messages, err := Hellos(names)
	if err == nil {
		t.Error("Hellos() with empty name expects to throw error")
	}
	if len(messages) > 0 {
		t.Error("No message expected.")
	}
}
