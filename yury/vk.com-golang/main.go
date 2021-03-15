package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MakeRequest(link string) {
	resp, err := http.Get(link)
	check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	WriteToFile("request.json", body)
}

type Items struct {
	response struct {
		count int
		items []string
	}
}

func WriteToFile(nameFile string, text []byte) {

	option := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(nameFile, option, os.FileMode(0600))
	check(err)

	_, err = file.Write(text)
	check(err)

	err = file.Close()
	check(err)
}

func main() {
	link := "https://api.vk.com/method/wall.get?PARAMETERS&access_token=eb34bd94eb34bd94eb34bd9438eb42630feeb34eb34bd948b79b09179195d7ad35a5a0e&v=5.130&domain=adkazy&count=1&offset="
	for i := 0; i < 1; i += 1 {
		MakeRequest(link + strconv.Itoa(i))
		fmt.Println(link + strconv.Itoa(i))

	}
}
