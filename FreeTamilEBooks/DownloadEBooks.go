package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

// Book is the representation of each book
// in the https://github.com/kishorek/Free-Tamil-Ebooks/blob/master/booksdb.xml file
type Book struct {
	//  <book>
	// 	<bookid>98628e8c-fcb3-404b-93ad-8782d0125efa</bookid>
	// 	<title>காதலென்பது </title>
	// 	<author>கா.பாலபாரதி</author>
	// 	<image>https://i1.wp.com/freetamilebooks.com/wp-content/uploads/2017/01/love.jpg?resize=215%2C300</image>
	// 	<link>http://freetamilebooks.com/ebooks/kadhalenpathu/</link>
	// 	<epub>http://freetamilebooks.com/download/%e0%ae%95%e0%ae%be%e0%ae%a4%e0%ae%b2%e0%af%86%e0%ae%a9%e0%af%8d%e0%ae%aa%e0%ae%a4%e0%af%81-epub/</epub>
	// 	<pdf>http://freetamilebooks.com/download/%e0%ae%95%e0%ae%be%e0%ae%a4%e0%ae%b2%e0%af%86%e0%ae%a9%e0%af%8d%e0%ae%aa%e0%ae%a4%e0%af%81-6-inch-pdf/</pdf>
	// 	<category> கட்டுரைகள் </category>
	// 	<date />
	// </book>

	BookID  string `xml:"bookid"`
	Title   string `xml:"title"`
	Author  string `xml:"author"`
	EpubURL string `xml:"epub"`
}

// Books is the representation of the xmldocument
type Books struct {
	Books []Book `xml:"book"`
}

var wg sync.WaitGroup

func main() {
	for _, fname := range os.Args[1:] {
		// fmt.Println("\nReading " + fname + "\n")
		b, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Println("Read failed", err)
			return
		}

		books := Books{}
		err = xml.Unmarshal(b, &books)
		if err != nil {
			log.Println("Parsing error : ", err)
			return
		}

		for i, book := range books.Books {
			go func(i, n int, title, author, url string) {
				wg.Add(1)
				defer wg.Done()

				prefix := title + "-" + author
				// fmt.Printf("\n%d/%d\nHTTP GET %s\n", i, n, url)
				var resp *http.Response
				resp, err = http.Get(url)
				if err != nil {
					log.Println("ERROR:", prefix, err)
					return
				}
				defer resp.Body.Close()

				var body []byte
				body, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println("ERROR:", prefix, err)
					return
				}
				defer resp.Body.Close()

				// fmt.Printf("Creating %s.pub\n", prefix)
				err = ioutil.WriteFile(prefix+".epub", body, 0666)
				if err != nil {
					log.Println("ERROR:", prefix, err)
					return
				}

				fmt.Printf("Creating %s.txt file\n", prefix)
				cmd := exec.Command("ebook-convert", prefix+".epub", prefix+".txt")
				err = cmd.Run()
				if err != nil {
					log.Println("ERROR:", prefix, err)
					return
				}
			}(i, len(books.Books), book.Title, book.Author, book.EpubURL)
		}
		wg.Wait()
	}
}
