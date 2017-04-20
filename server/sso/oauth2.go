package sso

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

// JWTAuthConfig represents the config for the JWT authentication middleware.
type JWTAuthConfig struct {
	TokenCookieName string
	PrivateKey      *rsa.PrivateKey
}

// OAuth2Config represents the config for the OAuth2 middleware.
type OAuth2Config struct {
	JWTAuthConfig

	QueryParamNextPage string
	LoginPath          string
	LogoutPath         string
	CallbackPath       string
	OAuth2             *oauth2.Config
	EmailLookupFunc    func(t *oauth2.Token) (string, error)
	NoAuthn            bool
}

// JWTAuthFromConfig returns a JWT authentication middleware.
func JWTAuthFromConfig(conf *JWTAuthConfig) echo.MiddlewareFunc {
	if conf == nil {
		conf = &JWTAuthConfig{}
	}

	if conf.TokenCookieName == "" {
		conf.TokenCookieName = "jwt-token"
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return authn(conf, c, next)
		}
	}
}

// OAuth2FromConfig returns an OAuth2 JWT authn middleware.
func OAuth2FromConfig(conf *OAuth2Config) echo.MiddlewareFunc {
	if conf == nil {
		conf = &OAuth2Config{}
	}

	if conf.QueryParamNextPage == "" {
		conf.QueryParamNextPage = "redirect_url"
	}
	if conf.TokenCookieName == "" {
		conf.TokenCookieName = "jwt-token"
	}
	if conf.LoginPath == "" {
		conf.LoginPath = "/login"
	}
	if conf.LogoutPath == "" {
		conf.LogoutPath = "/logout"
	}
	if conf.CallbackPath == "" {
		conf.CallbackPath = "/auth/callback"
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			switch c.Request().URL.Path {
			case conf.LoginPath:
				return login(conf, c)
			case conf.LogoutPath:
				return logout(conf, c)
			case conf.CallbackPath:
				return callback(conf, c)
			default:
				if conf.NoAuthn {
					return next(c)
				}
				return authn(&conf.JWTAuthConfig, c, next)
			}
		}
	}
}

//return c.Redirect(http.StatusFound, loginPath+"?"+keyNextPage+"="+url.QueryEscape(c.Request().URL.RequestURI()))
func authn(conf *JWTAuthConfig, c echo.Context, next echo.HandlerFunc) error {
	cookie, err := c.Cookie(conf.TokenCookieName)
	if err != nil && err != http.ErrNoCookie {
		return err
	}
	if cookie == nil {
		c.Logger().Printf("no authn cookie")
		return c.NoContent(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.GetSigningMethod("RS256") {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return &conf.PrivateKey.PublicKey, nil
	})
	if err != nil {
		c.Logger().Printf("error parsing jwt, %s", err)
		return c.NoContent(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("authenticated", true)
		c.Set("email", claims["email"])
		return next(c)
	}

	return c.NoContent(http.StatusUnauthorized)
}

func login(conf *OAuth2Config, c echo.Context) error {
	to := c.QueryParam(conf.QueryParamNextPage)
	return c.Redirect(http.StatusFound, conf.OAuth2.AuthCodeURL(to, oauth2.AccessTypeOnline))
}

func logout(conf *OAuth2Config, c echo.Context) error {
	to := c.QueryParam(conf.QueryParamNextPage)

	cookie := &http.Cookie{
		Name:    conf.TokenCookieName,
		Value:   "",
		Expires: time.Now(),
		MaxAge:  -1,
		Path:    "/",
	}

	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, to)
}

func callback(conf *OAuth2Config, c echo.Context) error {
	next := c.QueryParam("state")
	code := c.QueryParam("code")
	t, err := conf.OAuth2.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}

	email, err := conf.EmailLookupFunc(t)
	if err != nil {
		return err
	}

	expire := time.Now().Add(time.Hour * 24 * 3)

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims{
		"email": email,
		"exp":   expire.Unix(),
	})

	tokenString, err := token.SignedString(conf.PrivateKey)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	//next = next + "?token=" + tokenString

	cookie := &http.Cookie{
		Name:    conf.TokenCookieName,
		Value:   tokenString,
		Expires: expire,
		Path:    "/",
	}

	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, next)
}
