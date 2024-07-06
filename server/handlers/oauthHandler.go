package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func getOauthConfig(withPrivate bool) *oauth2.Config {
	scopes := []string{}
	if withPrivate {
		scopes = append(scopes, "repo")
	}
	return &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/auth/gh/callback", os.Getenv("SERVER_URL")),
		ClientID:     os.Getenv("GITHUB_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_OAUTH_CLIENT_SECRET"),
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
	}
}

func OauthGHLogin(c echo.Context) error {
	oauthState := generateStateOauthCookie(c.Response().Writer)

	u := getOauthConfig(false).AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, u)
	return nil
}
func OauthGHLoginExtended(c echo.Context) error {
	oauthState := generateStateOauthCookie(c.Response().Writer)

	u := getOauthConfig(true).AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, u)
	return nil
}

func OauthGHCallback(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer

	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return fmt.Errorf("invalid oauth state")
	}
	token, err := getAccessToken(r.FormValue("code"))
	if err != nil {
		return err
	}

	s, err := session.Get("session", c)
	if err != nil {
		log.Println(err.Error())
	}
	s.Values["token"] = token
	err = s.Save(c.Request(), c.Response())
	if err != nil {
		log.Println(err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("SERVER_URL"))
	return nil
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getAccessToken(code string) (string, error) {
	token, err := getOauthConfig(false).Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("code exchange error: %s", err.Error())
	}
	return token.AccessToken, nil
}
