package main

import (
	"log"
	"testing"
)

var Mytests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"positive-secenario", 100.0,10.0,10.0,false},
	{"0-secenario", 100.0,0,0,true},
	{"negnum-secenario", 100.0,-10.0,-10.0,false},
}

func TestDivision(t *testing.T){
	
	for _, tt := range Mytests{
		actual, err := divide(tt.dividend,tt.divisor)
		if err != nil{
			t.Error("expected an error but did not get one")
		}else{
			if err != nil{
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if actual != tt.expected{
			t.Errorf("expected %f but got %f", tt.expected, actual)
		}
		
	}
	log.Println(Mytests[1])
}

// these are manual test which is inefficient
// func Test_positive_divide(t *testing.T) {
// 	_, err := divide(10.0, 3.3)

// 	if err != nil {
// 		t.Error("ERROR : EVERYONE PANIC !!!! AHHHHH boogly boogly boogly ~~~", err)
// 	}
// }

// func Test_0_divide(t *testing.T) {
// 	_, err := divide(10.0, 0)

// 	if err == nil {
// 		t.Error("I expected an error and didnt get one *thinking emote*", err)
// 	}
// }

// func Test_negnum_divide(t *testing.T) {
// 	_, err := divide(10.0, -10)

// 	if err != nil {
// 		t.Error("ERROR : EVERYONE PANIC !!!! AHHHHH boogly boogly boogly ~~~", err)
// 	}
// }
