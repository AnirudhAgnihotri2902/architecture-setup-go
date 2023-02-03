package test

import (
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/service"
	"log"
	"testing"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 13},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {
	var count int = 0
	for _, test := range addTests {
		count++
		if output := service.Add(test.arg1, test.arg2); output != test.expected {
			log.Println("Test no.", count, "failed: Output", output, "not equal to expected ", test.expected)
		}
	}
}
