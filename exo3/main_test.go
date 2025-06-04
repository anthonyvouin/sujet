package main

import (
	"testing"
)


func TestDictionaryGenerator(t *testing.T) {
	tests := []struct {
		name       string
		dictionary []string
		wordLength int
		expectWord bool
	}{
		{
			name:       "word 5 letters found",
			dictionary: []string{"apple", "grape", "test"},
			wordLength: 5,
			expectWord: true,
		},
		{
			name:       "No word of the requested length",
			dictionary: []string{"apple", "grape"},
			wordLength: 3,
			expectWord: false,
		},
		{
			name:       "Empty dictionary",
			dictionary: []string{},
			wordLength: 5,
			expectWord: false,
		},
		{
			name:       "Several words of the same length",
			dictionary: []string{"cat", "dog", "bat", "rat"},
			wordLength: 3,
			expectWord: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := DictionaryGenerator{Dictionary: tt.dictionary}
			result := gen.Generate(tt.wordLength)
			
			if tt.expectWord && result == "" {
				t.Error("Expected a word, but received an empty string")
			}
			
			if !tt.expectWord && result != "" {
				t.Errorf("Expected an empty string, but received: %s", result)
			}
			
			if result != "" && len(result) != tt.wordLength {
				t.Errorf("Incorrect word length: got %d, want %d", len(result), tt.wordLength)
			}
			
			if result != "" {
				found := false
				for _, word := range tt.dictionary {
					if word == result {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("The generated word '%s' is not in the dictionary", result)
				}
			}
		})
	}
}

func TestRandomGenerator(t *testing.T) {
	gen := RandomGenerator{}
	
	tests := []struct {
		name       string
		wordLength int
	}{
		{"word 5 letters", 5},
		{"word 3 letters", 3},
		{"word 10 letters", 10},
		{"word 1 letter", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gen.Generate(tt.wordLength)
			
			if len(result) != tt.wordLength {
				t.Errorf("Incorrect word length: got %d, want %d", len(result), tt.wordLength)
			}
			
			for i, r := range result {
				if r < 'a' || r > 'z' {
					t.Errorf("Invalid character at position %d: got %c", i, r)
				}
			}
		})
	}
}

func TestRandomLettersGenerator(t *testing.T) {
	tests := []int{1, 3, 5, 10, 26}
	
	for _, length := range tests {
		t.Run(string(rune(length+'0')), func(t *testing.T) {
			result := RandomLettersGenerator(length)
			
			if len(result) != length {
				t.Errorf("Incorrect word length: got %d, want %d", len(result), length)
			}
			
			for i, r := range result {
				if r < 'a' || r > 'z' {
					t.Errorf("Invalid character at position %d: got %c", i, r)
				}
			}
			if length > 1 {
				results := make(map[string]bool)
				for i := 0; i < 10; i++ {
					word := RandomLettersGenerator(length)
					results[word] = true
				}
				if len(results) < 2 {
					t.Log("Attention: the generator might not be sufficiently random")
				}
			}
		})
	}
}

func TestIAGuesser(t *testing.T) {
	tests := []struct {
		name       string
		dictionary []string
		wordLength int
		expectWord bool
	}{
		{
			name:       "Dictionary with words of the correct length",
			dictionary: []string{"apple", "grape", "peach"},
			wordLength: 5,
			expectWord: true,
		},
		{
			name:       "Dictionary without words of the correct length",
			dictionary: []string{"cat", "dog"},
			wordLength: 5,
			expectWord: true, 
		},
		{
			name:       "Empty dictionary",
			dictionary: []string{},
			wordLength: 5,
			expectWord: true, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guesser := IAGuesser{Dictionary: tt.dictionary}
			result := guesser.Guess(tt.wordLength)
			
			if len(result) != tt.wordLength {
				t.Errorf("Incorrect word length: got %d, want %d", len(result), tt.wordLength)
			}
			validChars := true
			for _, r := range result {
				if r < 'a' || r > 'z' {
					validChars = false
					break
				}
			}
			
			if !validChars {
				t.Error("The generated word contains invalid characters")
			}
		})
	}
}


func TestPrintResult(t *testing.T) {
	t.Run("Does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PrintResult panicked: %v", r)
			}
		}()
		
		guess := "apple"
		result := []LetterStatus{CorrectPosition, WrongPosition, Absent, CorrectPosition, WrongPosition}
		
		PrintResult(guess, result)
	})
}



