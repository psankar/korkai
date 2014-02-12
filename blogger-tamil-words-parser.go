/*
 * A parser to identify the list of tamil words from a blogger archive
 * sorted both alphabetically and usage count wise
 * 
 * Author: Sankar சங்கர் <sankar.curiosity@gmail.com>
 */

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Category struct {
	Term string `xml:"term,attr"`
}

type Entry struct {
	Title    string     `xml:"title"`
	Content  string     `xml:"content"`
	Category []Category `xml:"category"`
}

type Feed struct {
	Entry []Entry `xml:"entry"`
}

func delim(r rune) bool {
	return r == '<' || r == '>' || r == '/' || r == ' ' || r == '.' || r == '!' || r == '"' || r == ',' || r == ';' || r == '(' || r == ')' || r == '=' || r == ':' || r == '&' || r == '\'' || r == '?' || r == '’' || r == '#' || r == '-' || r == '”' || r == '‘' || r == '“' || r == ' ' || r == '[' || r == ']' || r == '{' || r == '}' || r == '' || r == '\n' || r == '|' || r == ':' || r == '`'
}

func main() {

	for _, fname := range os.Args[1:] {

		fmt.Println("\nReading " + fname + "\n")
		b, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println("Read failed")
			return
		}

		v := Feed{}
		err = xml.Unmarshal(b, &v)
		if err != nil {
			fmt.Println("Parsing error : ")
			fmt.Println(err)
			return
		}
		m := make(map[string]int)

		for _, i := range v.Entry {

			isPost := false

			/* proceed only if the entry is a post
			 * ignore comments, settings etc. */
			for _, cat := range i.Category {
				if cat.Term == "http://schemas.google.com/blogger/2008/kind#post" {
					isPost = true
					break
				}
			}

			if isPost == true {
				fmt.Printf("Parsing %s\n", i.Title)

				/* Extract the plain text content out of the HTML mess */
				tokens := strings.FieldsFunc(i.Content, delim)

				for _, token := range tokens {
					r, l := utf8.DecodeRuneInString(token)
					if l != 1 {
						if unicode.Is(unicode.Tamil, r) {
							/* Count only tamil words */
							m[token] = m[token] + 1
						}
					}
				}

			}
		}

		fmt.Println("========")
		fmt.Println("Parsing complete. Now onto tokenizing and ranking.")
		fmt.Println("========")

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
	return
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
