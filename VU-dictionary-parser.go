// Author: Sankar P <sankar.curiosity@gmail.com>
// Can be used to parse the VU Dictionary at https://github.com/rprabhu/TamilDictionary
package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("TamilVUDictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)

	var words []string

	for _, jsonFile := range files {
		if !strings.HasSuffix(jsonFile.Name(), ".json") {
			continue
		}
		log.Println("Parsing: ", jsonFile.Name())

		b, err := ioutil.ReadFile(jsonFile.Name())
		if err != nil {
			log.Fatal(err)
		}

		var data map[string]string

		if err = json.Unmarshal(b, &data); err != nil {
			panic(err)
		}

		for k, v := range data {
			w.WriteString(k + "\n")
			// log.Println(k, v)
			ws := strings.FieldsFunc(v, func(r rune) bool {
				switch r {
				case '(', ';', ' ', '.', ')', ',', ':', '\'':
					return true
				}
				return false
			})
			words = append(words, ws...)
		}
	}
	log.Println(words[:10])
	sort.Strings(words)
	log.Println(words[:10])

	for _, word := range words {
		if !strings.ContainsAny(word, "+") {
			// Skip words ending in some letters
			suffixes := []string{"க்", "ங்", "ச்", "ஞ்", "ட்", "த்", "ந்", "ப்", "வ்", "ற்"}
			hasSuffix := false
			for _, suffix := range suffixes {
				if strings.HasSuffix(word, suffix) {
					hasSuffix = true
					break
				}
			}
			if !hasSuffix {
				w.WriteString(word + "\n")
			}
		}
	}
	w.Flush()
}
