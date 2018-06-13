package main

import (
	"fmt"
	testing "testing"
)

func TestValidWords(t *testing.T) {
	allWords := []string{
		"help",
		"googglegoo",
		"airport",
	}

	validWords := ValidWords("en_US", allWords)

	if len(validWords) != 2 {
		t.Errorf("valid words returns valid words only")
	}

	for _, word := range validWords {
		fmt.Println(word)
	}
}

func TestConvertToMapped(t *testing.T) {
	boggleChars := BoggleChars{
		Lang: "en_US",
		Rows: []BoggleRows{
			{
				Cols: []BoggleCols{
					{
						Char: "a",
					}, {
						Char: "c",
					},
				},
			}, {
				Cols: []BoggleCols{
					{
						Char: "t",
					}, {
						Char: "m",
					},
				},
			},
		},
	}

	if len(boggleChars.Rows) != 2 {
		t.Errorf("format error")
	}
	if len(boggleChars.Rows[0].Cols) != 2 {
		t.Errorf("format error")
	}

	mapped := ConvertToMapped(boggleChars)

	if (*mapped)[0][0].Char != "a" {
		t.Error("ConvertToMapped maps chars (a)")
	}
	if (*mapped)[0][1].Char != "c" {
		t.Error("ConvertToMapped maps chars (c)")
	}
	if (*mapped)[1][0].Char != "t" {
		t.Error("ConvertToMapped maps chars (t)")
	}
	if (*mapped)[1][1].Char != "m" {
		t.Error("ConvertToMapped maps chars (m)")
	}

	if (*mapped)[0][0].East.Char != "c" {
		t.Error("ConvertToMapped maps East (c)")
	}
	if (*mapped)[0][0].South.Char != "t" {
		t.Error("ConvertToMapped maps South (c)")
	}
	if (*mapped)[0][0].SouthEast.Char != "m" {
		t.Error("ConvertToMapped maps SouthEast (m)")
	}

	if (*mapped)[0][1].West.Char != "a" {
		t.Error("ConvertToMapped maps West (a)")
	}
	if (*mapped)[0][1].SouthWest.Char != "t" {
		t.Error("ConvertToMapped maps West (a)")
	}
	if (*mapped)[0][1].South.Char != "m" {
		t.Error("ConvertToMapped maps West (a)")
	}

	if (*mapped)[1][0].East.Char != "m" {
		t.Error("ConvertToMapped maps West (m)")
	}
	if (*mapped)[1][0].North.Char != "a" {
		t.Error("ConvertToMapped maps North (a)")
	}
	if (*mapped)[1][0].NorthEast.Char != "c" {
		t.Error("ConvertToMapped maps NorthEast (c)")
	}

	if (*mapped)[1][1].West.Char != "t" {
		t.Error("ConvertToMapped maps West (t)")
	}
	if (*mapped)[1][1].North.Char != "c" {
		t.Error("ConvertToMapped maps North (c)")
	}
	if (*mapped)[1][1].NorthWest.Char != "a" {
		t.Error("ConvertToMapped maps NorthWest (a)")
	}
}

func TestGetAllPossibleWords(t *testing.T) {
	boggleChars := BoggleChars{
		Lang: "en_US",
		Rows: []BoggleRows{
			{
				Cols: []BoggleCols{
					{
						Char: "a",
					}, {
						Char: "c",
					},
				},
			}, {
				Cols: []BoggleCols{
					{
						Char: "t",
					}, {
						Char: "m",
					},
				},
			},
		},
	}

	mapped := ConvertToMapped(boggleChars)
	allwords := GetAllPossibleWords(mapped)

	for _, word := range allwords {
		fmt.Println(word)
	}
}

func TestGetAllValidWords(t *testing.T) {
	boggleChars := BoggleChars{
		Lang: "en_US",
		Rows: []BoggleRows{
			{
				Cols: []BoggleCols{
					{
						Char: "a",
					}, {
						Char: "c",
					},
				},
			}, {
				Cols: []BoggleCols{
					{
						Char: "t",
					}, {
						Char: "m",
					},
				},
			},
		},
	}

	mapped := ConvertToMapped(boggleChars)
	allwords := GetAllPossibleWords(mapped)
	validWords := ValidWords("en_US", allwords)
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}
}
