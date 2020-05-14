package fuzz

import (
    "os"
    "github.com/argoproj/argo/workflow/validate"
)

func Fuzz(data []byte) int {
	f, err := os.Create("payload.txt")
	if err != nil {
		return -1
	}
	defer f.Close()
	defer os.Remove("payload.txt")
	_, err = f.Write(data)
	if err != nil {
		return -1
	}
	_, err = validate.ParseWfFromFile("payload.txt", true)
	if err != nil {
		return 0
	}
	return 1
}
