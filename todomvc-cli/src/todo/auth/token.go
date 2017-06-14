package auth

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"errors"
	"net/http"
	"net/url"
	"io"
)

type Config struct {
	Endpoint     string
	ClientId     string
	ClientSecret string
	Credentials  *TokenInfos
}

type AuthRealmResponse struct {
	Realm        string `json:"realm"`
	TokenService string `json:"token-service"`
}

type Credentials struct {
	Username string
	Password string
}

type TokenInfos struct {
	TokenService     string `json:"token_service"`
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

func GetAuthRealm(authRealm string) (*AuthRealmResponse, error) {
	r, err := http.Get(authRealm)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result AuthRealmResponse
	err = json.NewDecoder(r.Body).Decode(&result)
	return &result, err
}

func GetOfflineToken(tokenServiceUrl string, creds *Credentials) (*TokenInfos, error) {
	r, err := http.PostForm(tokenServiceUrl+"/token", url.Values{
		"client_id":  {"todo-cli"},
		"client_secret": {os.Getenv("TODOMVC_CLIENT_SECRET")},
		"username":   {creds.Username},
		"password":   {creds.Password},
		"grant_type": {"password"},
		"scope":      {"offline_access"},
	})
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(r.Body)
		return nil, errors.New(r.Status + " : " + string(body))
	}
	var result TokenInfos
	err = json.NewDecoder(r.Body).Decode(&result)
	result.TokenService = tokenServiceUrl
	return &result, err
}

func SaveTokenInfos(infos *TokenInfos) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	dir := path.Join(usr.HomeDir, ".todomvc")
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		} else {
			return err
		}
	}
	file := path.Join(dir, "creds.json")
	os.Remove(file)

	b, err := json.Marshal(infos)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	return err
}

func LoadTokenInfos() (*TokenInfos, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	filename := path.Join(usr.HomeDir, ".todomvc", "creds.json")

	if _, err := os.Stat(filename); err == nil {
		file, e := ioutil.ReadFile(filename)
		if e != nil {
			return nil, err
		}
		var infos TokenInfos
		err = json.Unmarshal(file, &infos)
		return &infos, err
	}
	return nil, nil
}

func GetAccessToken(config *Config) (string, error) {
	// Do not throw any error when there is no credentials
	if config.Credentials == nil {
		return "", nil
	}
	r, err := http.PostForm(config.Credentials.TokenService+"/token", url.Values{
		"grant_type": {"refresh_token"},
		"client_id":  {config.ClientId},
		"client_secret": {config.ClientSecret},
		"refresh_token": {config.Credentials.RefreshToken},
	})
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		io.Copy(os.Stderr, r.Body)
		return "", errors.New(r.Status)
	}
	var result TokenInfos
	err = json.NewDecoder(r.Body).Decode(&result)
	return result.AccessToken, err
}

func RemoveTokenInfos() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	filename := path.Join(usr.HomeDir, ".todomvc", "creds.json")
	if _, err := os.Stat(filename); err == nil {
		return os.Remove(filename)
	}
	return nil
}