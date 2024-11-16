package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	code := "muxi-backend"
	passport := checkpoint11(code)
	fmt.Println(passport)
	checkpoint12(passport)
	password := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMSIsImlhdCI6MTczMTc1MDcyMiwibmJmIjoxNzMxNzUwNzIyfQ.tmohy6YOvQH7l8gUJDpAnQpCwSPydljx8uNCIyP2I-M"
	checkpoint2(password)
}

func checkpoint11(code string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://gtainmuxi.muxixyz.com/api/v1/organization/code?", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("code", "1")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("checkpoint1-1:", response.Data.Text)
	fmt.Println()

	return resp.Header.Get("Passport")
}

func checkpoint12(password string) (link string, header string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://gtainmuxi.muxixyz.com/api/v1/organization/code?", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("code", password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ""
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("checkpoint1-2:", response.Data.Text)
	fmt.Println()

	re := regexp.MustCompile(`(http[^,]*)，`)
	link = re.FindStringSubmatch(string(body))[1]
	return link, resp.Header.Get("map-fragments")
}

func checkpoint2(password, info string) (link string, header string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("code", password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ""
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("checkpoint2:", response.Data.Text)
	return
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Text      string `json:"text"`
	ExtraInfo string `json:"extra_info"`
}

type DecodedInfo struct {
	SecretKey string `json:"secret_key"`
	ErrorCode string `json:"error_code"`
}
