package helper

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 5 {
		// error
		t.Fail()
	}
	fmt.Println("Emteka")
}

func TestAddDede(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		// error
		t.FailNow()
	}
	fmt.Println("Amtaka")
}