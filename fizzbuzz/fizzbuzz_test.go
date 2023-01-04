package fizzbuzz_test

import (
	"testing"

	"github.com/SaphMB/fizzbuzz/fizzbuzz"
)

func TestRespondSlowly(t *testing.T) {
	t.Run("when the input is not a multiple of 3 or 5", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 1, expectedResponse: "1"},
			{input: 2, expectedResponse: "2"},
			{input: 4, expectedResponse: "4"},
			{input: 7, expectedResponse: "7"},
			{input: 8, expectedResponse: "8"},
		}

		for _, test := range tests {
			if fizzbuzz.RespondSlowly(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.RespondSlowly(test.input))
			}
		}
	})

	t.Run("when the input is a multiple of 3", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 3, expectedResponse: "Fizz"},
			{input: 6, expectedResponse: "Fizz"},
			{input: 9, expectedResponse: "Fizz"},
			{input: 12, expectedResponse: "Fizz"},
		}

		for _, test := range tests {
			if fizzbuzz.RespondSlowly(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.RespondSlowly(test.input))
			}
		}
	})

	t.Run("when the input is a multiple of 5", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 5, expectedResponse: "Buzz"},
			{input: 10, expectedResponse: "Buzz"},
		}

		for _, test := range tests {
			if fizzbuzz.RespondSlowly(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.RespondSlowly(test.input))
			}
		}
	})

	t.Run("when the input is a multiple of 3 and 5", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 15, expectedResponse: "FizzBuzz"},
			{input: 30, expectedResponse: "FizzBuzz"},
		}

		for _, test := range tests {
			if fizzbuzz.RespondSlowly(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.RespondSlowly(test.input))
			}
		}
	})
}

func TestRespond(t *testing.T) {
	t.Run("when the input is not a multiple of 3 or 5", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 1, expectedResponse: "1"},
			{input: 2, expectedResponse: "2"},
			{input: 4, expectedResponse: "4"},
			{input: 7, expectedResponse: "7"},
			{input: 8, expectedResponse: "8"},
		}

		for _, test := range tests {
			if fizzbuzz.Respond(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.Respond(test.input))
			}
		}
	})

	t.Run("when the input is a multiple of 3", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 3, expectedResponse: "Fizz"},
			{input: 6, expectedResponse: "Fizz"},
			{input: 9, expectedResponse: "Fizz"},
			{input: 12, expectedResponse: "Fizz"},
		}

		for _, test := range tests {
			if fizzbuzz.Respond(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.Respond(test.input))
			}
		}
	})

	t.Run("when the input is a multiple of 5", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 5, expectedResponse: "Buzz"},
			{input: 10, expectedResponse: "Buzz"},
		}

		for _, test := range tests {
			if fizzbuzz.Respond(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.Respond(test.input))
			}
		}
	})

	t.Run("when the input is a multiple of 3 and 5", func(t *testing.T) {
		type cases struct {
			input            int
			expectedResponse string
		}

		tests := []cases{
			{input: 15, expectedResponse: "FizzBuzz"},
			{input: 30, expectedResponse: "FizzBuzz"},
		}

		for _, test := range tests {
			if fizzbuzz.Respond(test.input) != test.expectedResponse {
				t.Errorf("Expected %s, Got %s", test.expectedResponse, fizzbuzz.Respond(test.input))
			}
		}
	})
}

func BenchmarkRespondSlowly(b *testing.B) {
	b.Run("Respond slowly", func(b *testing.B) {
		for _, num := range []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 30} {
			fizzbuzz.RespondSlowly(num)
		}
	})
}

func BenchmarkRespond(t *testing.B) {
	t.Run("Respond", func(b *testing.B) {
		for _, num := range []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 30} {
			fizzbuzz.Respond(num)
		}
	})
}
