/*
 * A parser to identify the list of tamil words from a wordpress
 * archive sorted both alphabetically and usage count wise
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

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Title   string `xml:"title"`
	Content string `xml:"contentencoded"`
}

type Content struct {
	Data string `xml:",chardata"`
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

		v := Rss{}

		/*
					x := `<rss>
				<channel>
				      <item>
					<title>எழுத ஆசை</title>
					<link>http://nchokkan.wordpress.com/2008/11/23/%e0%ae%8e%e0%ae%b4%e0%af%81%e0%ae%a4-%e0%ae%86%e0%ae%9a%e0%af%88/</link>
					<pubDate>Sun, 23 Nov 2008 07:07:33 +0000</pubDate>
					<dc:creator>nchokkan</dc:creator>
					<guid isPermaLink="false">http://nchokkan.wordpress.com/?p=3</guid>
					<description></description>
					<content:encoded><![CDATA[கழுத்தில் யாரேனும் டெட்லைன் கத்தி வைத்தால்மட்டுமே எழுத வரும் என்கிற அளவு கெட்டுப்போய்விட்ட முழுச் சோம்பேறி நான்.
			என். சொக்கன் ...
			]]></content:encoded>
				</item>
				</channel>
				</rss>`

					err = xml.Unmarshal([]byte(x), &v)
		*/
		err = xml.Unmarshal(b, &v)
		if err != nil {
			fmt.Println("Parsing error : ")
			fmt.Println(err)
			return
		}

		//fmt.Println(v)

		m := make(map[string]int)

		for _, i := range v.Channel.Item {

			fmt.Printf("Parsing %s\n", i.Title)
			//fmt.Printf("Content is : [[[[%s]]]\n", i.Content)

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
