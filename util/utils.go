package util

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"runtime/debug"
)

func HandlePanic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("there is a panic in goroutine:\n%v", string(debug.Stack()))
		}
	}()

	fn()
}

func PanicIfError(msg string, err error) {
	if err != nil {
		log.Panicf("%v, err=%v", msg, err)
	}
}

func LoadJSON(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}
