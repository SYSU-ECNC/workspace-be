package auth

import (
	"encoding/gob"
	"net/http"

	"github.com/SYSU-ECNC/workspace-be/internal/pkg/config"
	"github.com/SYSU-ECNC/workspace-be/internal/pkg/sessions"
	"github.com/dghubble/gologin"
	gologinOAuth2 "github.com/dghubble/gologin/oauth2"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type ECNCUserInfo struct {
	Uid   int    `json:"uid"`
	NetID string `json:"netid"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Email string `json:"email"`
}

var endpoint = oauth2.Endpoint{
	AuthURL:  config.Get("sso_authorize_url"),
	TokenURL: config.Get("sso_token_url"),
}
var oauth2Config = oauth2.Config{
	ClientID:     config.Get("client_id"),
	ClientSecret: config.Get("client_secret"),
	RedirectURL:  config.Get("sso_redirect_url"),
	Endpoint:     endpoint,
}

var stateConfig = gologin.DefaultCookieConfig

var Authorize = gin.WrapH(gologinOAuth2.StateHandler(stateConfig, gologinOAuth2.LoginHandler(&oauth2Config, nil)))

func init() {
	gob.Register(&ECNCUserInfo{})
}

func callbackSuccess(c *gin.Context, token string) {
	session := sessions.Store(c)

	user := &ECNCUserInfo{}
	_, err := sling.New().Get(config.Get("sso_info_url")).Add("Authorization", "Bearer "+token).ReceiveSuccess(user)

	if err != nil {
		panic(err)
	}

	session.Set("user", user)
	session.Save()

	c.JSON(200, session.Get("user"))
}

func Callback(c *gin.Context) {
	callbackSuccessWrap := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		ecncToken, err := gologinOAuth2.TokenFromContext(ctx)

		if err != nil {
			c.Error(err)
			return
		}

		callbackSuccess(c, ecncToken.AccessToken)
	})
	callbackHandler := gologinOAuth2.CallbackHandler(&oauth2Config, callbackSuccessWrap, nil)
	gologinOAuth2.StateHandler(stateConfig, callbackHandler).ServeHTTP(c.Writer, c.Request)
}
