package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//func check(err error) check a mistakes in a err - variable
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//func MakeRequest(link string) []byte take a link from our variable and Get respounse
// to server. If server push request we a bode from request.
func MakeRequest(link string) []byte {

	resp, err := http.Get(link)
	check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body
}

//func WriteToFile(nameFile string, text []byte)
//Func write text to file
func WriteToFile(nameFile string, text []byte) {

	option := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(nameFile, option, os.FileMode(0600))
	check(err)

	_, err = file.Write(text)
	check(err)

	err = file.Close()
	check(err)
}

//func writeText(textByte string) string take a text and oarcing the text and likes
// And return clear text. Only text and likes from textByte
func writeText(textByte string) string {
	myText := ""
	textTemplate := "\"text\":"
	// j := 0
	for strings.Index(textByte, textTemplate) != -1 {

		indexOfText := strings.Index(textByte, textTemplate)
		textByte = textByte[indexOfText+len(textTemplate)+1:]
		i := strings.Index(textByte, "\"")
		myText += textByte[:i] + "\n" + "===================\n"
	}
	fmt.Println(myText)
	return myText
}

func main() {
	var bodyJson []byte
	link := "https://api.vk.com/method/wall.get?PARAMETERS&access_token=eb34bd94eb34bd94eb34bd9438eb42630feeb34eb34bd948b79b09179195d7ad35a5a0e&v=5.130&domain=adkazy&count=100&offset="
	for i := 0; i < 2500; i += 100 {
		bodyJson = MakeRequest(link + strconv.Itoa(i))
		WriteToFile("request.json", []byte(writeText(string(bodyJson))))
	}
}
