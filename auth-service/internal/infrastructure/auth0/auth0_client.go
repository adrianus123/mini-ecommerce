package auth0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Auth0Client struct {
	Domain       string
	ClientID     string
	ClientSecret string
	Audience     string
	MgmtToken    string
}

func (a *Auth0Client) Register(email, password string) (string, error) {
	url := "https://" + a.Domain + "/api/v2/users"

	payload := map[string]interface{}{
		"email":      email,
		"password":   password,
		"connection": "Username-Password-Authentication",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+a.MgmtToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("Response:", result)

	return result["user_id"].(string), nil
}

func (a *Auth0Client) Login(email, password string) (string, error) {
	url := "https://" + a.Domain + "/oauth/token"

	payload := map[string]interface{}{
		"grant_type":    "http://auth0.com/oauth/grant-type/password-realm",
		"username":      email,
		"password":      password,
		"client_id":     a.ClientID,
		"client_secret": a.ClientSecret,
		"audience":      a.Audience,
		"scope":         "openid profile email",
		"realm":         "Username-Password-Authentication",
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["error"] != nil {
		return "", errors.New(result["error_description"].(string))
	}

	return result["access_token"].(string), nil
}
