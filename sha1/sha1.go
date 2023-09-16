package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sum, err := sha1sum("http.log.gz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)

	sum, err = sha1sum("sha1.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}

// cat http.log.gz | gunzip | sha1sum
func sha1sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		reader, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer reader.Close()
		r = reader
	}

	writer := sha1.New()

	if _, err := io.Copy(writer, r); err != nil {
		return "", err
	}

	signature := writer.Sum(nil)
	return fmt.Sprintf("%x", signature), nil
}
