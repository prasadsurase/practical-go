package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type Result struct {
	FileName      string
	CalculatedSha string
	ProvidedSha   string
	Match         bool
}

func main() {
	rootDir := "tmp/taxi/"
	file, err := os.Open(path.Join(rootDir, "sha256sum.txt"))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	results, err := calculateResults(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(results)
}

func calculateResults(r io.Reader) ([]Result, error) {
	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanLines)

	results := []Result{}
	for fileScanner.Scan() {
		fields := strings.Fields(fileScanner.Text())
		if len(fields) < 2 {
			fmt.Errorf("bad line: %q", fileScanner.Text())
		} else {
			r := Result{
				FileName:    fields[1],
				ProvidedSha: fields[0],
			}
			sha, err := sha256sum(fields[1])
			if err != nil {
				fmt.Errorf("bad line: %q", err)
			}
			r.CalculatedSha = sha
			r.Match = r.CalculatedSha == r.ProvidedSha
			fmt.Println("-------------")
			fmt.Printf("%#v\n", r)
			results = append(results, r)
			// fmt.Println(results)
		}
	}
	// fmt.Println(results)
	return results, nil
}

// cat http.log.gz | gunzip | sha1sum
func sha256sum(fileName string) (string, error) {
	rootDir := "tmp/taxi"
	file, err := os.Open(path.Join(rootDir, fileName) + ".bz2")

	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
