/*
Package tag implements functions for tagging parts of speech.
*/
package tag

import (
	"bytes"
	"encoding/gob"
	"strings"
)

// Token represents a tagged section of text.
type Token struct {
	Text string
	Tag  string
}

// TupleSlice is a slice of tuples in the form (words, tags).
type TupleSlice [][][]string

// Len returns the length of a Tuple.
func (t TupleSlice) Len() int { return len(t) }

// Swap switches the ith and jth elements in a Tuple.
func (t TupleSlice) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// ReadTagged converts pre-tagged input into a TupleSlice suitable for training.
func ReadTagged(text, sep string) TupleSlice {
	t := TupleSlice{}
	for _, sent := range strings.Split(text, "\n") {
		tokens := []string{}
		tags := []string{}
		for _, token := range strings.Split(sent, " ") {
			parts := strings.Split(token, sep)
			tokens = append(tokens, parts[0])
			tags = append(tags, parts[1])
		}
		t = append(t, [][]string{tokens, tags})
	}
	return t
}

func getAsset(name string) *gob.Decoder {
	loc := "tag/" + name
	return gob.NewDecoder(bytes.NewReader(MustAsset(loc)))
}

func mustDecode(name string, e interface{}) {
	dec := getAsset(name)
	err := dec.Decode(e)
	if err != nil {
		panic(err)
	}
}
