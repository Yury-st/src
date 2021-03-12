package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Структура для передачи в шаблон
type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandlerGuestBook(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signature.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}
func viewHandlerNew(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func getStrings(fileName string) []string {

	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func main() {

	http.HandleFunc("/book", viewHandlerGuestBook)
	http.HandleFunc("/new", viewHandlerNew)
	// http.HandleFunc("/hello", viewHandler1)
	// http.HandleFunc("/goodbay", viewHandler2)
	// textTemplate := "{{range .}}In loop: {{.}}\n{{end}}\n"
	// executeTemplate(textTemplate, []string{"do", "re", "mi"})
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}
