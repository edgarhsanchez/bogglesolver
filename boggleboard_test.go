package main

import (
	"fmt"
	"testing"
)

func TestMappedBoggleWordToString(t *testing.T) {
	//arrange
	langMap, err := LoadAllLanguageFiles()
	jsontxt := []byte(`{"lang":"en_US","rows":[{"cols":[{"char":"h"},{"char":"m"},{"char":"v"},{"char":"y"}]},{"cols":[{"char":"b"},{"char":"u"},{"char":"x"},{"char":"a"}]},{"cols":[{"char":"y"},{"char":"t"},{"char":"a"},{"char":"w"}]},{"cols":[{"char":"s"},{"char":"o"},{"char":"o"},{"char":"p"}]}]}`)
	boggleChars := BoggleChars{}
	json.Unmarshal(jsontxt, &boggleChars)
	mapped := ConvertToMapped(boggleChars)
	words, err := GetAllValidWords(langMap["en_US"], mapped)
	if err != nil {
		t.Errorf(err.Error())
	}
	//act
	word := words[0]
	fmt.Println(word)

	//assert

}
