package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(getUserInfo("prasadsurase"))
}

func getUserInfo(handle string) (string, int, error) {
	url := "https://api.github.com/users/" + handle
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("error: %s", resp.Status)
	}
	var r struct {
		Name      string `json:"name"`
		RepoCount int    `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&r); err != nil {
		log.Fatalf("error: %s", err)
	}

	return r.Name, r.RepoCount, nil
}

func demo() {
	resp, err := http.Get("https://api.github.com/users/prasadsurase")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("error: %s", resp.Status)
	}
	// fmt.Printf("Content Type: %s\n", resp.Header.Get("content-type"))
	// fmt.Printf("Access Control Expose Headers: %s\n", resp.Header.Get("access-control-expose-headers"))

	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: %s", resp.Status)
	// }
	// fmt.Printf("%s\n", resp.Body)

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&r); err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Println(r)
	log.Printf("%#v\n", r)
}

type Reply struct {
	Name         string `json:"name"`
	Public_Repos int    `json:"public_repos"`
}
