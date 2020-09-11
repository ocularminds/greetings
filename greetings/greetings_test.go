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