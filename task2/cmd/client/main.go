package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"main/pkg/model"
	"net/http"
	"time"
)

func MyGetRequest() {
	request, err := http.NewRequest("GET", "http://localhost:8080/version", nil)
	if err != nil {
		fmt.Println("error in NewRequest:", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("error in Do:", err)
		return
	}
	defer response.Body.Close()
	respBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(respBody))
}

func MyPostRequest(base64String string) {
	req := model.DecodeRequest{Base64String: base64String}
	jsonBody, _ := json.Marshal(req)
	request, err := http.NewRequest("POST", "http://localhost:8080/decode", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("error in request:", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("error in Do:", err)
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	var resp model.DecodeResponse
	json.Unmarshal(body, &resp)
	fmt.Println(resp.DecodeString)
}

func MyHardOpRequest() {
	ctx, cansel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cansel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hard-op", nil)
	if err != nil {
		fmt.Println("error in NewRequest:", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("error in Do:", err)
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
func main() {
	MyGetRequest()
	MyPostRequest("aGVsbG8gZ29sYW5nIQ==")
	MyHardOpRequest()
}
