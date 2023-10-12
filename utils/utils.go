package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func LogDetails(val any, title string, isPanic bool) {
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

func UnmarshalJson[D any](value string, target D) (*D, error) {
	var data D
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		fmt.Println("Error unmarshaling secret data:", err)
		return nil, err
	}
	return &data, nil
}
