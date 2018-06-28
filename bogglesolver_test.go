package main

import (
	"fmt"
	testing "testing"

	"github.com/json-iterator/go"
)

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

	hunLangs, err := LoadAllLanguageFiles(10)
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
	allwords, err := GetAllValidWords(hunLangs[boggleChars.Lang], mapped, 10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, word := range allwords {
		fmt.Println(word)
	}
}

func TestGetAllValidWords(t *testing.T) {
	langMap, err := LoadAllLanguageFiles(10)

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
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 10)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}
}

func TestLargeBoard(t *testing.T) {

	//arrange
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	langMap, err := LoadAllLanguageFiles(10)
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"h"},{"char":"m"},{"char":"v"},{"char":"y"}]},{"cols":[{"char":"b"},{"char":"u"},{"char":"x"},{"char":"a"}]},{"cols":[{"char":"y"},{"char":"t"},{"char":"a"},{"char":"w"}]},{"cols":[{"char":"s"},{"char":"o"},{"char":"o"},{"char":"p"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)

	//act
	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 10)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}

	if len(validWords) != 41 {
		t.Errorf("ConvertToMapped did not find 41 words found: %s ", string(len(validWords)))
	}

	if (*mapped)[0][0].Char != "h" {
		t.Error("ConvertToMapped maps chars (a)")
	}
	if (*mapped)[0][0].South.Char != "b" {
		t.Error("ConvertToMapped maps chars (b")
	}
	if (*mapped)[0][0].East.Char != "m" {
		t.Error("ConvertToMapped maps chars (m)")
	}
	if (*mapped)[0][0].SouthEast.Char != "u" {
		t.Error("ConvertToMapped maps chars (m)")
	}
	// spells but
	if (*mapped)[1][0].Char != "b" {
		t.Error("Does NOT spell but")
	}
	if (*mapped)[1][0].East.Char != "u" {
		t.Error("Does NOT spell but")
	}
	if (*mapped)[1][0].SouthEast.Char != "t" {
		t.Error("Does NOT spell but")
	}
}
