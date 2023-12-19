package utils

import (
	"encoding/json"
	"log"
	"os"
)

func LogDetails(title string, val any, isPanic bool) {
	if isPanic {
		log.Panicf("%v :: %v\n", title, val)
	} else {
		log.Printf("%v :: %v\n", title, val)
	}
}

func ExitProgram(failed bool) {
	if failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func UnmarshalJson[D any](value string) (*D, error) {
	var data D
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return nil, err
	}
	return &data, nil
}
