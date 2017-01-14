package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
	// "strings"
	// "unicode"
	// "io/ioutil"
	// "unicode/utf8"
)

func delim(r rune) bool {
	return r == '<' || r == '>' || r == '/' || r == ' ' || r == '.' || r == '!' || r == '"' || r == ',' || r == ';' || r == '(' || r == ')' || r == '=' || r == ':' || r == '&' || r == '\'' || r == '?' || r == '’' || r == '#' || r == '-' || r == '”' || r == '‘' || r == '“' || r == ' ' || r == '[' || r == ']' || r == '{' || r == '}' || r == '' || r == '\n' || r == '|' || r == ':' || r == '`'
}

func main() {

	files, err := filepath.Glob("*.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(files)

	for _, fname := range files {
		m := make(map[string]int)

		log.Println(fname)

		var content []byte
		content, err = ioutil.ReadFile(fname)
		if err != nil {
			log.Println(err)
		}
		/* Extract the plain text content out of the HTML mess */
		words := strings.FieldsFunc(string(content), delim)

		for _, word := range words {
			tamilWord := true
			token := word

			for i := 0; i < len(word); {
				r, l := utf8.DecodeRuneInString(token)
				if !unicode.Is(unicode.Tamil, r) {
					//fmt.Println("Falsifying because of:", r, token)
					tamilWord = false
					break
				}
				i += l
				//fmt.Println(token)
				token = token[l:]
			}
			if tamilWord {
				/* Count only tamil words */
				//fmt.Println(word)
				m[word] = m[word] + 1
			}
		}
		// log.Println(m)

		kvs := NewValueSorter(m)
		kvs.Sort()

		statsFileName := fname + ".stats"
		var statsFile *os.File
		statsFile, err = os.Create(statsFileName)
		if err != nil {
			fmt.Println("File creation error:")
			fmt.Println(err)
			return
		}
		var tokens []string
		tokens = make([]string, 0, len(kvs.Keys))
		for k, v := range kvs.Keys {
			tokens = append(tokens, v)
			io.WriteString(statsFile, fmt.Sprintf("%d,%s\n", kvs.Vals[k], v))
		}
		statsFile.Close()

		fmt.Println("Statistics for " + fname + " saved to " + statsFileName)

		sort.Strings(tokens)

		tokensFileName := fname + ".tokens"
		var tokensFile *os.File
		tokensFile, err = os.Create(tokensFileName)
		if err != nil {
			fmt.Println("File creation error:")
			fmt.Println(err)
			return
		}
		for _, str := range tokens {
			_, err = io.WriteString(tokensFile, str+"\n")
			if err != nil {
				fmt.Println("File write error:")
				fmt.Println(err)
				return
			}
		}
		tokensFile.Close()
		fmt.Println("Statistics for " + fname + " saved to " + tokensFileName)
		fmt.Println("========")
	}
}

type ValueSorter struct {
	Keys []string
	Vals []int
}

func NewValueSorter(m map[string]int) *ValueSorter {

	kvs := &ValueSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int, 0, len(m)),
	}

	for k, v := range m {
		kvs.Keys = append(kvs.Keys, k)
		kvs.Vals = append(kvs.Vals, v)
	}

	return kvs
}

func (kvs *ValueSorter) Sort() {
	sort.Sort(kvs)
}

func (kvs *ValueSorter) Len() int {
	return len(kvs.Vals)
}

func (kvs *ValueSorter) Less(i, j int) bool {
	return kvs.Vals[i] > kvs.Vals[j]
}

func (kvs *ValueSorter) Swap(i, j int) {
	kvs.Vals[i], kvs.Vals[j] = kvs.Vals[j], kvs.Vals[i]
	kvs.Keys[i], kvs.Keys[j] = kvs.Keys[j], kvs.Keys[i]
}
