package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"main/pkg/model"
	"math/rand/v2"
	"net/http"
	"time"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("version: 1.0.0")
		w.Write([]byte("version: 1.0.0"))
	default:
		fmt.Printf("Method not allowed\n")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var req model.DecodeRequest
		json.NewDecoder(r.Body).Decode(&req)
		fmt.Printf("Input string: %s\n", req.Base64String)
		decodeString, err := base64.StdEncoding.DecodeString(req.Base64String)
		if err != nil {
			fmt.Printf("Invalid base64 input\n")
			http.Error(w, "Invalid base64 input", http.StatusBadRequest)
			return
		}
		res := model.DecodeResponse{DecodeString: string(decodeString)}
		json.NewEncoder(w).Encode(res)
		fmt.Printf("Decoded string: %s\n", res.DecodeString)
	default:
		fmt.Printf("Method not allowed\n")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	duration := time.Duration(rand.IntN(10)+11) * time.Second
	select {
	case <-time.After(duration):
		if rand.IntN(2) == 0 {
			w.Write([]byte("OK"))
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case <-ctx.Done():
		fmt.Printf("HardOpHandler cancelled\n")
		return
	}
}
