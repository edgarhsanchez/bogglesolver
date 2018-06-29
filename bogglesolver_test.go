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

func Test5x5Board(t *testing.T) {
	//arrange
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	langMap, err := LoadAllLanguageFiles(7)
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"g"},{"char":"i"},{"char":"x"},{"char":"b"},{"char":"x"}]},{"cols":[{"char":"r"},{"char":"e"},{"char":"y"},{"char":"b"},{"char":"i"}]},{"cols":[{"char":"b"},{"char":"w"},{"char":"t"},{"char":"y"},{"char":"t"}]},{"cols":[{"char":"k"},{"char":"y"},{"char":"u"},{"char":"l"},{"char":"i"}]},{"cols":[{"char":"e"},{"char":"i"},{"char":"l"},{"char":"h"},{"char":"f"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)

	//act
	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 7)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}
}

func Test20x20Board(t *testing.T) {
	//arrange
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	langMap, err := LoadAllLanguageFiles(5)
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"p"},{"char":"w"},{"char":"l"},{"char":"z"},{"char":"z"},{"char":"p"},{"char":"x"},{"char":"k"},{"char":"j"},{"char":"z"},{"char":"k"},{"char":"z"},{"char":"m"},{"char":"d"},{"char":"q"},{"char":"g"},{"char":"f"},{"char":"a"},{"char":"x"},{"char":"f"}]},{"cols":[{"char":"j"},{"char":"e"},{"char":"v"},{"char":"s"},{"char":"n"},{"char":"v"},{"char":"v"},{"char":"h"},{"char":"o"},{"char":"x"},{"char":"y"},{"char":"p"},{"char":"m"},{"char":"s"},{"char":"q"},{"char":"e"},{"char":"b"},{"char":"k"},{"char":"z"},{"char":"m"}]},{"cols":[{"char":"r"},{"char":"d"},{"char":"x"},{"char":"r"},{"char":"r"},{"char":"o"},{"char":"s"},{"char":"t"},{"char":"x"},{"char":"y"},{"char":"c"},{"char":"r"},{"char":"s"},{"char":"b"},{"char":"b"},{"char":"u"},{"char":"b"},{"char":"r"},{"char":"m"},{"char":"o"}]},{"cols":[{"char":"t"},{"char":"s"},{"char":"f"},{"char":"d"},{"char":"r"},{"char":"l"},{"char":"o"},{"char":"z"},{"char":"h"},{"char":"f"},{"char":"w"},{"char":"v"},{"char":"l"},{"char":"h"},{"char":"l"},{"char":"b"},{"char":"n"},{"char":"w"},{"char":"w"},{"char":"m"}]},{"cols":[{"char":"t"},{"char":"o"},{"char":"m"},{"char":"u"},{"char":"k"},{"char":"y"},{"char":"i"},{"char":"o"},{"char":"c"},{"char":"p"},{"char":"w"},{"char":"a"},{"char":"w"},{"char":"h"},{"char":"l"},{"char":"q"},{"char":"g"},{"char":"g"},{"char":"r"},{"char":"z"}]},{"cols":[{"char":"v"},{"char":"b"},{"char":"z"},{"char":"w"},{"char":"l"},{"char":"a"},{"char":"r"},{"char":"j"},{"char":"a"},{"char":"x"},{"char":"r"},{"char":"r"},{"char":"p"},{"char":"e"},{"char":"j"},{"char":"l"},{"char":"n"},{"char":"s"},{"char":"f"},{"char":"z"}]},{"cols":[{"char":"s"},{"char":"c"},{"char":"t"},{"char":"r"},{"char":"q"},{"char":"m"},{"char":"s"},{"char":"w"},{"char":"l"},{"char":"n"},{"char":"p"},{"char":"i"},{"char":"d"},{"char":"z"},{"char":"o"},{"char":"w"},{"char":"z"},{"char":"e"},{"char":"e"},{"char":"h"}]},{"cols":[{"char":"u"},{"char":"e"},{"char":"k"},{"char":"w"},{"char":"a"},{"char":"w"},{"char":"x"},{"char":"i"},{"char":"b"},{"char":"t"},{"char":"w"},{"char":"f"},{"char":"u"},{"char":"t"},{"char":"r"},{"char":"i"},{"char":"l"},{"char":"e"},{"char":"u"},{"char":"y"}]},{"cols":[{"char":"u"},{"char":"q"},{"char":"v"},{"char":"a"},{"char":"l"},{"char":"o"},{"char":"s"},{"char":"h"},{"char":"h"},{"char":"l"},{"char":"u"},{"char":"m"},{"char":"x"},{"char":"m"},{"char":"q"},{"char":"u"},{"char":"i"},{"char":"e"},{"char":"t"},{"char":"r"}]},{"cols":[{"char":"f"},{"char":"w"},{"char":"b"},{"char":"t"},{"char":"z"},{"char":"g"},{"char":"o"},{"char":"j"},{"char":"j"},{"char":"n"},{"char":"h"},{"char":"l"},{"char":"i"},{"char":"i"},{"char":"n"},{"char":"t"},{"char":"x"},{"char":"z"},{"char":"w"},{"char":"g"}]},{"cols":[{"char":"n"},{"char":"g"},{"char":"w"},{"char":"b"},{"char":"e"},{"char":"n"},{"char":"q"},{"char":"k"},{"char":"k"},{"char":"y"},{"char":"b"},{"char":"f"},{"char":"s"},{"char":"f"},{"char":"o"},{"char":"e"},{"char":"q"},{"char":"v"},{"char":"i"},{"char":"b"}]},{"cols":[{"char":"h"},{"char":"h"},{"char":"s"},{"char":"c"},{"char":"g"},{"char":"n"},{"char":"n"},{"char":"i"},{"char":"v"},{"char":"h"},{"char":"g"},{"char":"n"},{"char":"i"},{"char":"g"},{"char":"m"},{"char":"a"},{"char":"g"},{"char":"l"},{"char":"r"},{"char":"x"}]},{"cols":[{"char":"r"},{"char":"y"},{"char":"b"},{"char":"p"},{"char":"g"},{"char":"u"},{"char":"p"},{"char":"p"},{"char":"m"},{"char":"e"},{"char":"f"},{"char":"i"},{"char":"z"},{"char":"h"},{"char":"s"},{"char":"u"},{"char":"u"},{"char":"q"},{"char":"f"},{"char":"b"}]},{"cols":[{"char":"g"},{"char":"v"},{"char":"h"},{"char":"i"},{"char":"i"},{"char":"m"},{"char":"u"},{"char":"e"},{"char":"i"},{"char":"e"},{"char":"p"},{"char":"w"},{"char":"v"},{"char":"h"},{"char":"s"},{"char":"l"},{"char":"s"},{"char":"u"},{"char":"e"},{"char":"i"}]},{"cols":[{"char":"d"},{"char":"g"},{"char":"i"},{"char":"b"},{"char":"r"},{"char":"m"},{"char":"x"},{"char":"f"},{"char":"o"},{"char":"h"},{"char":"k"},{"char":"g"},{"char":"e"},{"char":"x"},{"char":"t"},{"char":"l"},{"char":"l"},{"char":"e"},{"char":"h"},{"char":"x"}]},{"cols":[{"char":"f"},{"char":"p"},{"char":"s"},{"char":"f"},{"char":"l"},{"char":"t"},{"char":"o"},{"char":"n"},{"char":"b"},{"char":"a"},{"char":"l"},{"char":"n"},{"char":"o"},{"char":"k"},{"char":"a"},{"char":"t"},{"char":"k"},{"char":"n"},{"char":"z"},{"char":"j"}]},{"cols":[{"char":"c"},{"char":"s"},{"char":"w"},{"char":"j"},{"char":"v"},{"char":"b"},{"char":"a"},{"char":"t"},{"char":"i"},{"char":"q"},{"char":"k"},{"char":"o"},{"char":"t"},{"char":"r"},{"char":"u"},{"char":"j"},{"char":"a"},{"char":"w"},{"char":"u"},{"char":"j"}]},{"cols":[{"char":"w"},{"char":"y"},{"char":"p"},{"char":"s"},{"char":"t"},{"char":"x"},{"char":"b"},{"char":"o"},{"char":"y"},{"char":"y"},{"char":"p"},{"char":"s"},{"char":"r"},{"char":"p"},{"char":"i"},{"char":"r"},{"char":"q"},{"char":"e"},{"char":"u"},{"char":"y"}]},{"cols":[{"char":"v"},{"char":"d"},{"char":"t"},{"char":"b"},{"char":"z"},{"char":"q"},{"char":"k"},{"char":"d"},{"char":"p"},{"char":"w"},{"char":"x"},{"char":"p"},{"char":"r"},{"char":"q"},{"char":"c"},{"char":"h"},{"char":"a"},{"char":"c"},{"char":"g"},{"char":"d"}]},{"cols":[{"char":"h"},{"char":"z"},{"char":"h"},{"char":"q"},{"char":"u"},{"char":"t"},{"char":"i"},{"char":"o"},{"char":"r"},{"char":"h"},{"char":"c"},{"char":"c"},{"char":"b"},{"char":"q"},{"char":"v"},{"char":"t"},{"char":"v"},{"char":"r"},{"char":"f"},{"char":"c"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)

	//act
	mapped := ConvertToMapped(boggleChars)
	validWords, err := GetAllValidWords(langMap["en_US"], mapped, 6)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//assert
	fmt.Println("valid words...")
	for _, word := range validWords {
		fmt.Println(word)
	}
}
