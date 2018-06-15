package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/client9/gospell"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
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

// MappedBoggleChar is a boggle board character/piece mapped to surrounding pieces
type MappedBoggleChar struct {
	Char      string
	North     *MappedBoggleChar
	NorthWest *MappedBoggleChar
	NorthEast *MappedBoggleChar
	West      *MappedBoggleChar
	East      *MappedBoggleChar
	South     *MappedBoggleChar
	SouthWest *MappedBoggleChar
	SouthEast *MappedBoggleChar
}

// HunspellLanguage is used to create dictionaries for languages from hunspell files
type HunspellLanguage struct {
	Lang    string
	Speller *gospell.GoSpell
}

func main() {

	langMap := LoadAllLanguageFiles()

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Use(gin.Recovery())

	router.Use(static.Serve("/", static.LocalFile("/public", true)))
	router.Use(static.Serve("/public", static.LocalFile("/public", true)))

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong: "+os.Getenv("VERSION"))
	})

	api.POST("/possiblewords", func(c *gin.Context) {
		boggleChars := BoggleChars{}

		if c.ShouldBindJSON(&boggleChars) == nil {
			mapped := ConvertToMapped(boggleChars)
			words := GetAllValidWords(langMap[boggleChars.Lang], mapped)
			c.JSON(http.StatusOK, words)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}

	})

	router.Run(":" + os.Getenv("PORT"))

}

// LoadAllLanguageFiles looks through all the available dictionary file and loads them
func LoadAllLanguageFiles() map[string]HunspellLanguage {
	dir := "hunspell"

	hunSpellMap := make(map[string]HunspellLanguage)
	err := filepath.Walk(dir, func(path string, fInfo os.FileInfo, err error) error {
		ext := filepath.Ext(fInfo.Name())
		if ext == ".dic" {
			lang := fInfo.Name()[0 : len(fInfo.Name())-len(ext)]

			affFile, affError := os.Open("hunspell/" + lang + ".aff")
			dicFile, dicError := os.Open("hunspell/" + lang + ".dic")
			goSpell, err := gospell.NewGoSpellReader(affFile, dicFile)

			if affError != nil {
				return affError
			}

			if dicError != nil {
				return dicError
			}

			if err != nil {
				return err
			}

			hunSpell := HunspellLanguage{}
			hunSpell.Lang = lang
			hunSpell.Speller = goSpell
			hunSpellMap[lang] = hunSpell
		}

		return nil
	})

	if err != nil {
		log.Println(err.Error())
	}

	return hunSpellMap
}

// GetAllValidWords finds all words valid or not for each piece on the boggle board
func GetAllValidWords(lang HunspellLanguage, mapped *[][]*MappedBoggleChar) []string {
	validWords := cmap.New()
	allWords := cmap.New()
	chans := make([]chan bool, len(*mapped))
	for chani, Row := range *mapped {
		for _, mBC := range Row {
			chanx := make(chan bool)
			chans[chani] = chanx
			go NewWordBranch(mBC, chanx, &allWords)
		}
	}

	for _, chanel := range chans {
		if <-chanel {
			for _, word := range allWords.Keys() {
				found := lang.Speller.Spell(word)
				if found {
					validWords.Set(word, struct{}{})
				}
			}
		}
	}

	orderWords := validWords.Keys()
	sort.Strings(orderWords)

	return orderWords
}

// NewWordBranch begins a new set of words starting from a single piece/char on the boggle board
func NewWordBranch(currentChar *MappedBoggleChar, channel chan bool, words *cmap.ConcurrentMap) {
	RecurseWords(currentChar, nil, words)
	channel <- true
}

// RecurseWords navigates through all the pieces creating possible words form the boggle board
func RecurseWords(currentChar *MappedBoggleChar, lastWord []*MappedBoggleChar, words *cmap.ConcurrentMap) {
	word := make([]*MappedBoggleChar, 0)
	if lastWord != nil {
		word = append(lastWord, currentChar)
	} else {
		word = append(word, currentChar)
	}
	// create word entry if char count greater than 3
	if len(word) >= 3 {
		var buffer bytes.Buffer
		for _, char := range word {
			buffer.Write([]byte(char.Char))
		}
		words.Set(buffer.String(), struct{}{})
	}

	// ensure char doesn't exist then add to bucket
	if currentChar.North != nil && !BoggleCharExists(currentChar.North, lastWord) {
		RecurseWords(currentChar.North, word, words)
	}
	if currentChar.NorthWest != nil && !BoggleCharExists(currentChar.NorthWest, lastWord) {
		RecurseWords(currentChar.NorthWest, word, words)
	}
	if currentChar.NorthEast != nil && !BoggleCharExists(currentChar.NorthEast, lastWord) {
		RecurseWords(currentChar.NorthEast, word, words)
	}
	if currentChar.West != nil && !BoggleCharExists(currentChar.West, lastWord) {
		RecurseWords(currentChar.West, word, words)
	}
	if currentChar.East != nil && !BoggleCharExists(currentChar.East, lastWord) {
		RecurseWords(currentChar.East, word, words)
	}
	if currentChar.South != nil && !BoggleCharExists(currentChar.South, lastWord) {
		RecurseWords(currentChar.South, word, words)
	}
	if currentChar.SouthWest != nil && !BoggleCharExists(currentChar.SouthWest, lastWord) {
		RecurseWords(currentChar.SouthWest, word, words)
	}
	if currentChar.SouthEast != nil && !BoggleCharExists(currentChar.SouthEast, lastWord) {
		RecurseWords(currentChar.SouthEast, word, words)
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
func ConvertToMapped(bchars BoggleChars) *[][]*MappedBoggleChar {
	mapped := make([][]*MappedBoggleChar, 0)
	for _, h := range bchars.Rows {
		row := make([]*MappedBoggleChar, 0)
		for _, cell := range h.Cols {
			row = append(row, &MappedBoggleChar{
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
