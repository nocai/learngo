package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	pictureUrl string
	seaweedAddress string
)

func init() {
	flag.StringVar(&pictureUrl, "pictureUrl", "", "picture url")
	flag.StringVar(&seaweedAddress, "seaweedAddress", "", "the seaweed address")

	flag.Parse()
}

func main() {
	resp, err := http.Get(pictureUrl)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	pictureName := filepath.Base(pictureUrl)
	file, err := os.Create(pictureName)
	defer func() {
		_ = file.Close()
	}()

	if _, err := io.Copy(file, resp.Body); err != nil {
		panic(err)
	}

	client := Client{seaweedAddress}
	result, err := client.Store(pictureName, resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(result)
}
