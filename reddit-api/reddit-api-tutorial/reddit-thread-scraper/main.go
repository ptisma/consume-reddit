package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiration  uint64 `json:"expires_in"`
	Scope       string `json:"scope"`
}

func main() {

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "beskucnik_na_feru")
	data.Add("password", "Wizard21298")

	encodedData := data.Encode()

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(encodedData))
	failOnError(err, "REQUEST")
	req.Header.Add("Authorization", "Basic "+basicAuth("95wusYL4TfUPURtBnswGQg", "qlMMTd2bN3ijKT7wrRlJtGpxQVyZ2w"))
	req.Header.Add("User-Agent", "reddit-api-scraper-test:0.01 by Pepe")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	res, err := client.Do(req)
	failOnError(err, "RESPONSE")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	failOnError(err, "BODY")
	fmt.Println(string(body))

	var token Token
	json.Unmarshal([]byte(body), &token)
	fmt.Println(token)

	// req, err = http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	// failOnError(err, "REQUEST 2")
	// req.Header.Add("Authorization", "bearer "+token.AccessToken)
	// req.Header.Add("User-Agent", "reddit-api-scraper-test:0.01 by Pepe")

	// res, err = client.Do(req)
	// failOnError(err, "RESPONSE 2")
	// defer res.Body.Close()

	// body, err = ioutil.ReadAll(res.Body)
	// failOnError(err, "BODY 2")
	// fmt.Println(string(body))

	req, err = http.NewRequest("GET", "https://oauth.reddit.com/r/croatia/hot", nil)
	failOnError(err, "REQUEST 3")
	req.Header.Add("Authorization", "bearer "+token.AccessToken)
	req.Header.Add("User-Agent", "reddit-api-scraper-test:0.01 by Pepe")

	res, err = client.Do(req)
	failOnError(err, "RESPONSE 3")
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	failOnError(err, "BODY 3")
	fmt.Println(string(body))

}
