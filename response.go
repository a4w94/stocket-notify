package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type tokenResponse struct {
	Status      string `json:"Name"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

func newTokenResponse(raw []byte) *tokenResponse {
	ret := &tokenResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return ret
	}
	fmt.Println("TokenResponse:", "Status:", &ret.Status, "Message:", &ret.Message, "Token:", &ret.AccessToken, ret)
	return ret
}
