package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v57/github"
)

// GitHub App ID
const GithubAppID = ""

// pem 形式 の秘密鍵
const GitHubSecretKey = ""

func main() {
	now := time.Now()

	payload := jwt.MapClaims{
		"exp": now.Unix() + 60,
		"iat": now.Unix(),
		"iss": GithubAppID,
	}

	secret := []byte(GitHubSecretKey)

	block, _ := pem.Decode(secret)

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	jwt, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	client := github.NewClient(&http.Client{
		Transport: &BearerTokenTransport{JWT: jwt},
	})

	installations, _, err := client.Apps.ListInstallations(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	if len(installations) != 1 {
		panic("expected 1 installation")
	}

	accessToken, _, err := client.Apps.CreateInstallationToken(context.Background(), installations[0].GetID(), nil)
	if err != nil {
		panic(err)
	}

	client = github.NewClient(&http.Client{
		Transport: &TokenTransport{Token: accessToken.GetToken()},
	})

	if _, _, err := client.Repositories.Dispatch(context.Background(), "otakakot", "playground-github-actions", github.DispatchRequestOptions{
		EventType: "repository-dispatch",
	}); err != nil {
		panic(err)
	}
}

type BearerTokenTransport struct {
	JWT string
}

func (btt *BearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+btt.JWT)

	return http.DefaultTransport.RoundTrip(req)
}

type TokenTransport struct {
	Token string
}

func (tt *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "token "+tt.Token)

	return http.DefaultTransport.RoundTrip(req)
}
