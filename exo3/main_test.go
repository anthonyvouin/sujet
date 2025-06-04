package main

import (
	"reflect"
	"testing"
)

func TestEvaluateGuess(t *testing.T) {
	tests := []struct {
		secret string
		guess  string
		want   GuessResult
	}{
		{
			secret: "apple",
			guess:  "apple",
			want:   GuessResult{CorrectPosition, CorrectPosition, CorrectPosition, CorrectPosition, CorrectPosition},
		},
		{
			secret: "apple",
			guess:  "apron",
			want:   GuessResult{CorrectPosition, CorrectPosition, Absent, Absent, Absent},
		},
		{
			secret: "apple",
			guess:  "pleap",
			want:   GuessResult{WrongPosition, WrongPosition, WrongPosition, WrongPosition, WrongPosition},
		},
		{
			secret: "apple",
			guess:  "zzzzz",
			want:   GuessResult{Absent, Absent, Absent, Absent, Absent},
		},
		{
			secret: "apple",
			guess:  "aplep",
			want:   GuessResult{CorrectPosition, CorrectPosition, WrongPosition, WrongPosition, WrongPosition},
		},
		{
			secret: "hello",
			guess:  "llama",
			want:   GuessResult{Absent, WrongPosition, WrongPosition, Absent, Absent},
		},
		{
			secret: "books",
			guess:  "spoon",
			want:   GuessResult{WrongPosition, Absent, WrongPosition, Absent, Absent},
		},
		{
			secret: "speed",
			guess:  "erase",
			want:   GuessResult{WrongPosition, Absent, Absent, WrongPosition, WrongPosition},
		},
	}

	for _, tt := range tests {
		t.Run(tt.secret+"-"+tt.guess, func(t *testing.T) {
			got := EvaluateGuess(tt.secret, tt.guess)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvaluateGuess(%q, %q) = %v, want %v", tt.secret, tt.guess, got, tt.want)
			}
		})
	}
}

func TestDictionaryGenerator(t *testing.T) {
	dict := []string{"apple", "grape", "test"}
	gen := DictionaryGenerator{Dictionary: dict}

	word := gen.Generate(5)
	if word != "apple" && word != "grape" {
		t.Errorf("Expected 'apple' or 'grape', got %q", word)
	}

	word = gen.Generate(10)
	if word != "" {
		t.Errorf("Expected empty string, got %q", word)
	}
}

func TestRandomGenerator(t *testing.T) {
	gen := RandomGenerator{}
	word := gen.Generate(5)

	if len(word) != 5 {
		t.Errorf("Expected word of length 5, got %d", len(word))
	}

	for _, char := range word {
		if char < 'a' || char > 'z' {
			t.Errorf("Expected lowercase letter, got %c", char)
		}
	}
}
