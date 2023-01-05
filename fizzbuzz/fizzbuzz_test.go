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
	for n := 0; n < b.N; n++ {
		fizzbuzz.RespondSlowly(15)
	}
}

func BenchmarkRespond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fizzbuzz.Respond(15)
	}
}

func benchmarkRespond(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fizzbuzz.Respond(i)
	}
}

func benchmarkRespondSlowly(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fizzbuzz.RespondSlowly(i)
	}
}

func BenchmarkRespond1(b *testing.B)  { benchmarkRespond(1, b) }
func BenchmarkRespond3(b *testing.B)  { benchmarkRespond(3, b) }
func BenchmarkRespond5(b *testing.B)  { benchmarkRespond(5, b) }
func BenchmarkRespond10(b *testing.B) { benchmarkRespond(10, b) }
func BenchmarkRespond15(b *testing.B) { benchmarkRespond(15, b) }
func BenchmarkRespond30(b *testing.B) { benchmarkRespond(30, b) }

func BenchmarkRespondSlowly1(b *testing.B)  { benchmarkRespondSlowly(1, b) }
func BenchmarkRespondSlowly3(b *testing.B)  { benchmarkRespondSlowly(3, b) }
func BenchmarkRespondSlowly5(b *testing.B)  { benchmarkRespondSlowly(5, b) }
func BenchmarkRespondSlowly10(b *testing.B) { benchmarkRespondSlowly(10, b) }
func BenchmarkRespondSlowly15(b *testing.B) { benchmarkRespondSlowly(15, b) }
func BenchmarkRespondSlowly30(b *testing.B) { benchmarkRespondSlowly(30, b) }

func FuzzRespond(f *testing.F) {
	f.Add(2)
	f.Fuzz(func(t *testing.T, input int) {
		response := fizzbuzz.Respond(input)
		if response == "" {
			t.Fail()
		}
	})
}

func FuzzRespondSlowly(f *testing.F) {
	f.Add(2)
	f.Fuzz(func(t *testing.T, input int) {
		response := fizzbuzz.RespondSlowly(input)
		if response == "" {
			t.Logf("failed running against input: %d", input)
			t.Fail()
		}
	})
}
