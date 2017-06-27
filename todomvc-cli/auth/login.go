package auth

import (
	"strings"
	"log"
)

func Login(authServer string, username string, password string, realm string) {
	if username = strings.TrimSpace(username); username == "" {
		log.Fatalln("Username not specified.")
	}

	realmResp, err := GetAuthRealm(authServer + "realms/" + realm)
	if err != nil {
		log.Fatalf("Unable to get realm info : %s", err)
	}

	creds := &Credentials{
		Username: username,
		Password: password,
	}

	tokenInfos, err := GetOfflineToken(realmResp.TokenService, creds)
	if err != nil {
		log.Fatalf("Unable to get offline token : %s", err)
	}

	saveError := SaveTokenInfos(tokenInfos)

	if saveError != nil {
		log.Fatalf("Unable to save your credentials : %s", saveError)
	}

	log.Println("You are now logged in!")
}
