package client

import (
	"net/http"
	"io"
	"net/url"
	"fmt"
	"log"
	"todo/auth"
	"os"
)

type Client struct {
	Config *auth.Config
}

func (k *Client) Do(method string, url string, query *url.Values, body io.Reader) (*http.Response, error) {
	accessToken, err := auth.GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, k.Config.Endpoint+url, body)
	if err != nil {
		return nil, err
	}

	if accessToken != "" {
		req.Header.Add("Authorization", "Bearer "+accessToken)
	}

	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Can't contact remote server : %s", err)
	}
	if resp.StatusCode == 401 {
		log.Fatalln("Seems that you are not logged in ! Please use login command...")
	}

	return resp, err
}

func (k *Client) Get(url string, query *url.Values) (*http.Response, error) {
	return k.Do("GET", url, query, nil)
}

func (k *Client) Delete(url string, query *url.Values) (*http.Response, error) {
	return k.Do("DELETE", url, query, nil)
}

func (k *Client) Post(url string, query *url.Values, body io.Reader) (*http.Response, error) {
	return k.Do("POST", url, query, body)
}

func (k *Client) Put(url string, query *url.Values, body io.Reader) (*http.Response, error) {
	return k.Do("PUT", url, query, body)
}

func TodoMVCClient(endpoint string) (*Client, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Invalid server endpoint %s: %s", endpoint, err)
	}

	creds, err := auth.LoadTokenInfos()
	if err != nil {
		return nil, err
	}

	config := &auth.Config{
		Endpoint:    endpoint,
		Credentials: creds,
		ClientId: "todo-cli",
		ClientSecret: os.Getenv("TODOMVC_CLIENT_SECRET"),
	}

	return &Client{
		Config: config,
	}, nil
}