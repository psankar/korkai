/*
 * A parser to identify the list of tamil words used in the wikipedia
 * and sorted them alphabetically and based on usage count
 * 
 * Author: Sankar சங்கர் <sankar.curiosity@gmail.com>
 */

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

type Page struct {
	Title string `xml:"title"`
	Text  string `xml:"revision>text"`
}

var m map[string]int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) != 2 {
		fmt.Println(`

Usage: wikipedia-tamil-words-parser.go <wikipedia-archive-file-location>.xml

Example: wikipedia-tamil-words-parser.go tawiki-latest-pages-articles.xml

`)
		return
	}

	xmlFile, err := os.Open(os.Args[1:][0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	fmt.Println("Fetching each page from the archive and calculating each unique word and its frequency")

	m = make(map[string]int)

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch elem := t.(type) {
		case xml.StartElement:
			if elem.Name.Local == "page" {
				var p Page
				decoder.DecodeElement(&p, &elem)
				wg.Add(1)
				go analyze(&p)
			}
		default:
		}

	}

	/* Waits until all the go routines return */
	wg.Wait()
	fmt.Println("Word fetching done")

	fmt.Println("Now sorting the words based on usage")
	kvs := NewValueSorter(m)
	kvs.Sort()
	fmt.Println("Sorting words based on usage Done")

	statsFileName := "wikipedia-tamil-words.stats"
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

	fmt.Println("Statistics for wikipedia saved to " + statsFileName)

	sort.Strings(tokens)

	tokensFileName := "wikipedia-tamil-words.tokens"
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
	fmt.Println("Tokens for wikipedia saved to " + tokensFileName)

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

func analyze(p *Page) {
	//fmt.Println("Analyzing " + p.Title)

	/* Extract the plain text content out of the HTML mess */
	tokens := strings.FieldsFunc(p.Text, delim)

	for _, token := range tokens {
		r, l := utf8.DecodeRuneInString(token)
		if l != 1 {
			if unicode.Is(unicode.Tamil, r) {
				/* Count only tamil words */
				mutex.Lock()
				m[token] = m[token] + 1
				mutex.Unlock()
			}
		}
	}
	wg.Done()
}

func delim(r rune) bool {
	return r == '<' || r == '>' || r == '/' || r == ' ' || r == '.' || r == '!' || r == '"' || r == ',' || r == ';' || r == '(' || r == ')' || r == '=' || r == ':' || r == '&' || r == '\'' || r == '?' || r == '’' || r == '#' || r == '-' || r == '”' || r == '‘' || r == '“' || r == ' ' || r == '[' || r == ']' || r == '{' || r == '}' || r == '' || r == '\n' || r == '|' || r == ':' || r == '`'
}
