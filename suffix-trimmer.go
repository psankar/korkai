package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(content), "\n")
	if words[len(words)-1] == "" {
		words = words[:len(words)-1]
	}

	f, err := os.Create("SuffixTrimmed.txt")
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)

	for _, word := range words {
		if !strings.ContainsAny(word, "+♥ஃ௦௪௫௰௯௩௬௨௲௵௴௷௱+$ஜஷஸஹ0123456789abcdefghijklmnopqrstuvxyz~`-´௳௶௸௹௺") {
			// Skip words ending in some letters
			suffixes := []string{"க்", "ங்", "ச்", "ஞ்", "ட்", "த்", "ந்", "ப்", "வ்", "ற்", "னும்", "னின்", "னிடம்", "ரிடம்", "ரின்", "கின்றன", "கின்ற", "லும்", "வும்", "ளில்", "ளுடன்"}
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
