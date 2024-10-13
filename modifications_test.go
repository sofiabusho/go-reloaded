package main

import (
	"testing"
)

func TestHandleAllModifications(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		{"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
	}

	for _, tt := range tests {
		actual := handleAllModifications(tt.input)
		if actual != tt.expectedOutput {
			t.Errorf("handleAllModifications(%s):\n expected: %s,\n got: %s", tt.input, tt.expectedOutput, actual)
		}
	}
}

// func TestHandleHexAndBin(t *testing.T) {
// 	tests := []struct {
// 		input           string
// 		expectedOutcome string
// 	}{
// 		{"234 (hex)", "564"},
// 		{"1011 (bin)", "11"},
// 		{"invalid (bin)", "invalid (bin)"}, // Test case for invalid input
// 	}

// 	for _, tt := range tests {
// 		actual := handleHexAndBin(tt.input)
// 		if actual != tt.expectedOutcome {
// 			t.Errorf("handleHexAndBin(%s): expected %s, got %s", tt.input, tt.expectedOutcome, actual)
// 		}
// 	}
// }

// func TestHandleTextModifications(t *testing.T) {
// 	tests := []struct {
// 		input           string
// 		expectedOutcome string
// 	}{
// 		{"this life (up, 2)", "THIS LIFE"},
// 		{"those pants (cap)", "those Pants"},
// 		{"WHY DID YOU DO THIS? (low, 5)", "why did you do this?"},
// 	}

// 	for _, tt := range tests {
// 		actual := handleTextModifications(tt.input)
// 		if actual != tt.expectedOutcome {
// 			t.Errorf("handleTextModifications(%s): expected %s, got %s", tt.input, tt.expectedOutcome, actual)
// 		}
// 	}
// }

// func TestCapitalizeWord(t *testing.T) {
// 	tests := []struct {
// 		input           string
// 		expectedOutcome string
// 	}{
// 		{"this life (cap)", "this Life"},
// 		{"everyday is different (cap, 3)", "Everyday Is Different"},
// 	}

// 	for _, tt := range tests {
// 		actual := capitalizeWord(tt.input)
// 		if actual != tt.expectedOutcome {
// 			t.Errorf("capitalize(%s): expected %s, got %s", tt.input, tt.expectedOutcome, actual)
// 		}
// 	}
// }

// func TestAdjustPunctuation(t *testing.T) {
// 	tests := []struct {
// 		input           string
// 		expectedOutcome string
// 	}{
// 		{"' whatever '", "'whatever"},
// 		{" it is what it is ... right ?", "it is what it is... right?"},
// 	}

// 	for _, tt := range tests {
// 		actual := adjustPunctuation(tt.input)
// 		if actual != tt.expectedOutcome {
// 			t.Errorf("adjustPunctuation(%s): expected %s, got %s", tt.input, tt.expectedOutcome, actual)
// 		}
// 	}
// }

// func TestHandleAtoAn(t *testing.T) {
// 	tests := []struct {
// 		input           string
// 		expectedOutcome string
// 	}{
// 		{"a watch", "a watch"},
// 		{"a actual rock", "an actual rock"},
// 	}

// 	for _, tt := range tests {
// 		actual := handleAtoAn(tt.input)
// 		if actual != tt.expectedOutcome {
// 			t.Errorf("handleAtoAn(%s): expected %s, got %s", tt.input, tt.expectedOutcome, actual)
// 		}
// 	}
// }
