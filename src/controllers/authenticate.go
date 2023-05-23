package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Access_token string `json:"access_token"`
	Expires_in int `json:"expires_in"`
	Token_type string `json:"token_type"`
}

func GetToken(c *gin.Context) {
	url := os.Getenv("AUTH0_TOKEN_URL")

	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")
	audience := os.Getenv("AUTH0_AUDIENCE")
	payload := fmt.Sprintf("{\"client_id\":\"%s\",\"client_secret\":\"%s\",\"audience\":\"%s\",\"grant_type\":\"client_credentials\"}", client_id, client_secret, audience)

	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(payload))

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// fmt.Println(res)
	// fmt.Println(string(body))
	c.JSON(200, gin.H{
		"data": response,
	})
}