package main

import (
	"encoding/json"
	"os"

	"github.com/juddbaguio/go-melon-chart/melon"
)

func createfile(fileName string, data interface{}) error {
	if data == nil {
		return nil
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	payloadByte, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(payloadByte)
	return err
}

func main() {
	melon := melon.New()
	top100 := melon.Top100()
	daily := melon.TopToday()

	createfile("melon-top-100.json", top100)
	createfile("melon-top-today.json", daily)

}
