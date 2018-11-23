package main

import (
	"fmt"
	"github.com/Dmitriy-Opria/bet_parser/go-log"
	"io/ioutil"
	"net/http"
	"os"
)

const source = "https://ua1xbet.com/ua/line/eSports/"

func main() {

	logger := log.WithFields(log.Fields{"test_key": "test_value"})
	resp, err := http.Get(source)
	if err != nil {
		log.Fatal("fatal")
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Error(fmt.Sprintf("Can`t get source response, error: %s", err.Error()))
	}
	fmt.Println(string(body))
}

// Die kills the program.
func Die(err error) {
	if err == nil {
		return
	}
	log.Fatal(err.Error())
	os.Exit(1)
}
