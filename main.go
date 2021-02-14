package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	HTMLVersion "github.com/lestoni/html-version"
	"golang.org/x/net/html"
)

func version(w http.ResponseWriter, r *http.Request) {
	// url := "https://stackoverflow.com"
	URL := r.URL.Query().Get("v")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	input := URL
	version, err := HTMLVersion.DetectFromURL(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<h1>Home 24</h1><p>The HTML version of this page %s is %s </p>", input, version)
}
func getTitle(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("get")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	textTags := []string{
		"title",
	}

	tag := ""
	enter := false

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					fmt.Fprintf(w, "<h1>Home 24</h1><p>Title for this page is: %s</p>", data)
				}
			}
		}
	}
}
func getHeadings(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("get")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	textTags := []string{
		"h1,h2,h3",
	}

	tag := ""
	enter := false

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					fmt.Fprintf(w, "<h1>Home 24</h1><p>Headings for this page are: %s</p>", data)
				}
			}
		}
	}
}
func getContent(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("get")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	textTags := []string{
		"p",
	}

	tag := ""
	enter := false

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}

		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					fmt.Fprintf(w, "<p><b>Content</b> \n Content of this page is: %s</p>", data)
				}
			}
		}
	}
}
func getLinks(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("get")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	textTags := []string{
		"a",
	}

	tag := ""
	enter := false

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}

		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					fmt.Fprintf(w, "<p><b>Links</b> \n Links of this page are: %s</p>", data)
				}
			}
		}
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("has")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	input := URL
	str := "login"

	if strings.Contains(str, input) {
		fmt.Fprintf(w, "<h1>Home 24</h1><p>This link %s  has a login form</p>", input)
	}
	fmt.Fprintf(w, "<h1>Home 24</h1><p>There is no login form for this link %s</p>", input)

}
func main() {
	http.HandleFunc("/", version)
	http.HandleFunc("/title", getTitle)
	http.HandleFunc("/heading", getHeadings)
	http.HandleFunc("/content", getContent)
	http.HandleFunc("/link", getLinks)
	http.HandleFunc("/log", login)
	http.ListenAndServe(":8000", nil)

}
