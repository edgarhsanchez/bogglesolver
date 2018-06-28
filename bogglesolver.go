package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/client9/gospell"
	jsoniter "github.com/json-iterator/go"

	"github.com/orcaman/concurrent-map"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// BoggleRows represents rows
type BoggleRows struct {
	Cols []BoggleCols `json:"cols"`
}

// BoggleCols represents columns
type BoggleCols struct {
	Char string `json:"char"`
}

// BoggleChars represents a boggle board and for a given language
type BoggleChars struct {
	Lang string       `json:"lang"`
	Rows []BoggleRows `json:"rows"`
}

// HunspellLanguage is used to create dictionaries for languages from hunspell files
type HunspellLanguage struct {
	Lang    string
	Speller *gospell.GoSpell
}

// LoadAllLanguageFiles looks through all the available dictionary file and loads them
func LoadAllLanguageFiles(maxWordSize int) (map[string]HunspellLanguage, error) {
	dir := "hunspell"

	hunSpellMap := make(map[string]HunspellLanguage)
	err := filepath.Walk(dir, func(path string, fInfo os.FileInfo, err error) error {
		ext := filepath.Ext(fInfo.Name())
		if ext == ".dic" {
			lang := fInfo.Name()[0 : len(fInfo.Name())-len(ext)]

			goSpell, err := gospell.NewGoSpell("hunspell/"+lang+".aff", "hunspell/"+lang+".dic")
			if err != nil {
				fmt.Println(err.Error())
				return err
			} else {
				hunSpell := HunspellLanguage{}
				hunSpell.Lang = lang
				hunSpell.Speller = goSpell
				largerWordsOnly := make(map[string]struct{})
				for key := range goSpell.Dict {
					wordSize := len(key)
					if wordSize >= 3 && wordSize <= maxWordSize {
						largerWordsOnly[strings.ToLower(key)] = struct{}{}
					}
				}
				goSpell.Dict = largerWordsOnly
				hunSpellMap[lang] = hunSpell
			}
		}

		return nil
	})

	return hunSpellMap, err
}

// GetAllValidWords finds all words valid or not for each piece on the boggle board
func GetAllValidWords(lang HunspellLanguage, mapped *MappedBoggleWords, maxWordSize int) ([]string, error) {
	var wg = new(sync.WaitGroup)
	validWords := cmap.New()

	// extract dictionary words for faster searching
	dictWords := cmap.New()
	dictKeys := reflect.ValueOf(lang.Speller.Dict).MapKeys()

	for keyIndex := range dictKeys {
		dictWord := dictKeys[keyIndex].String()
		if len(dictWord) >= 3 {
			dictWords.Set(dictWord, struct{}{})
		}
	}

	wg.Add(len(*mapped) * len((*mapped)[0]))
	for _, Row := range *mapped {
		for _, mBC := range Row {
			go NewWordBranch(mBC, wg, &validWords, &dictWords, maxWordSize)
		}
	}

	wg.Wait()

	orderWords := validWords.Keys()
	sort.Strings(orderWords)

	return orderWords, nil
}

// ArrayStartsWith returns true if any item in the strarr arrary starts with the given prefix.  Uses a binary search
func ArrayStartsWith(prefix string, strarr *[]string) bool {
	i := sort.SearchStrings(*strarr, prefix)
	return i >= 0
}

// NewWordBranch begins a new set of words starting from a single piece/char on the boggle board
func NewWordBranch(currentChar *MappedBoggleChar, wg *sync.WaitGroup, words *cmap.ConcurrentMap, dictWords *cmap.ConcurrentMap, maxWordSize int) {
	mappedWord := make(MappedBoggleWord, 0)
	var bWord bytes.Buffer
	RecurseWords(currentChar, mappedWord, bWord, words, dictWords, maxWordSize)
	wg.Done()
}

// RecurseWords navigates through all the pieces creating possible words form the boggle board
func RecurseWords(currentChar *MappedBoggleChar, priorMappedWord MappedBoggleWord, lastWord bytes.Buffer, words *cmap.ConcurrentMap, dictWords *cmap.ConcurrentMap, maxWordSize int) {

	// initiate a new mapped word
	mappedWord := make(MappedBoggleWord, 0)
	mappedWord = append(priorMappedWord, currentChar)

	//get new word
	var bWord bytes.Buffer

	if lastWord.Len() > 0 {
		bWord.WriteString(lastWord.String())
	}

	bWord.WriteString(currentChar.Char)
	word := bWord.String()
	if len(word) <= maxWordSize {
		if len(word) >= 3 {
			if _, ok := dictWords.Get(word); ok {
				(*words).Set(word, struct{}{})
			}
		}
		// ensure char doesn"t exist then add to bucket
		if currentChar.North != nil {
			// ensure this char has not been processed prior
			if !mappedWord.Contains(currentChar.North) {
				RecurseWords(currentChar.North, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.NorthWest != nil {
			// ensure this char has not been processed prior
			if !mappedWord.Contains(currentChar.NorthWest) {
				RecurseWords(currentChar.NorthWest, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.NorthEast != nil {
			if !mappedWord.Contains(currentChar.NorthEast) {
				RecurseWords(currentChar.NorthEast, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.West != nil {
			if !mappedWord.Contains(currentChar.West) {
				RecurseWords(currentChar.West, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.East != nil {
			if !priorMappedWord.Contains(currentChar.East) {
				RecurseWords(currentChar.East, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.South != nil {
			if !mappedWord.Contains(currentChar.South) {
				RecurseWords(currentChar.South, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.SouthWest != nil {
			if !mappedWord.Contains(currentChar.SouthWest) {
				RecurseWords(currentChar.SouthWest, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

		if currentChar.SouthEast != nil {
			if !mappedWord.Contains(currentChar.SouthEast) {
				RecurseWords(currentChar.SouthEast, mappedWord, bWord, words, dictWords, maxWordSize)
			}
		}

	}

}

// BoggleCharExists checks that a boggle piece has not already been used for a word
// pieces/chars may only be used once
func BoggleCharExists(ch *MappedBoggleChar, word []*MappedBoggleChar) bool {
	for _, char := range word {
		if char == ch {
			return true
		}
	}

	return false
}

// ConvertToMapped creates a map for each boggle piece with its surrounding pieces
func ConvertToMapped(bchars BoggleChars) *MappedBoggleWords {
	mapped := make(MappedBoggleWords, 0)
	for x, h := range bchars.Rows {
		row := make([]*MappedBoggleChar, 0)
		for y, cell := range h.Cols {
			row = append(row, &MappedBoggleChar{
				XYID:      fmt.Sprintf("%d%d", x, y),
				Char:      cell.Char,
				North:     nil,
				NorthWest: nil,
				NorthEast: nil,
				West:      nil,
				East:      nil,
				South:     nil,
				SouthWest: nil,
				SouthEast: nil,
			})
		}
		mapped = append(mapped, row)
	}

	for i1, h := range mapped {
		for i2, cell := range h {
			if i1 > 0 {
				cell.North = mapped[i1-1][i2]
			}
			if i1 < len(bchars.Rows)-1 {
				cell.South = mapped[i1+1][i2]
			}
			if i2 > 0 {
				cell.West = mapped[i1][i2-1]
			}
			if i2 < len(h)-1 {
				cell.East = mapped[i1][i2+1]
			}
			if i1 > 0 && i2 > 0 {
				cell.NorthWest = mapped[i1-1][i2-1]
			}
			if i1 < len(bchars.Rows)-1 && i2 > 0 {
				cell.SouthWest = mapped[i1+1][i2-1]
			}
			if i1 > 0 && i2 < len(h)-1 {
				cell.NorthEast = mapped[i1-1][i2+1]
			}
			if i1 < len(bchars.Rows)-1 && i2 < len(h)-1 {
				cell.SouthEast = mapped[i1+1][i2+1]
			}
		}
	}
	return &mapped
}
