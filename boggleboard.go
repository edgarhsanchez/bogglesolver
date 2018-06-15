package main

import (
	"bytes"
)

// MappedBoggleChar is a boggle board character/piece mapped to surrounding pieces
type MappedBoggleChar struct {
	XYID      string
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

// MappedBoggleWord represents a boggle word with adjacent characters/pieces specified
type MappedBoggleWord []*MappedBoggleChar

// MappedBoggleWords is an array MappedBoggleWord
type MappedBoggleWords []MappedBoggleWord

// ToString converts a mappedboggleword to a string
func (mword *MappedBoggleWord) ToString() string {
	var buffer bytes.Buffer
	for _, mbc := range *mword {
		buffer.WriteString((*mbc).Char)
	}

	return buffer.String()
}

// Contains searches for a given MappedBoggleChar and returns true if found or false if not found in the MappedBoggleWord
// if nil is passed in Contains always returns true
func (mword *MappedBoggleWord) Contains(char *MappedBoggleChar) bool {
	if char == nil {
		return true
	}

	for _, ch := range *mword {
		if ch == char {
			return true
		}
	}

	return false
}
