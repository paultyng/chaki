package sso

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

//GithubEnterpriseEndpoint generates an OAuth2 endpoint for a Github Enterprise installation.
func GithubEnterpriseEndpoint(domain string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  "https://" + domain + "/login/oauth/authorize",
		TokenURL: "https://" + domain + "/login/oauth/access_token",
	}
}

//GithubEnterpriseEmailLookup generated an email lookup function for the specified domain.
func GithubEnterpriseEmailLookup(domain string) func(t *oauth2.Token) (string, error) {
	return func(t *oauth2.Token) (string, error) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://"+domain+"/api/v3/user/emails", nil)
		if err != nil {
			return "", err
		}

		req.Header.Set("Authorization", "token "+t.AccessToken)
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)

		emails := []struct {
			email    string
			verified bool
			primary  bool
		}{}

		err = json.Unmarshal(contents, &emails)
		if err != nil {
			return "", err
		}

		email := ""
		for _, e := range emails {
			if e.primary && e.verified {
				email = e.email
			}
		}

		return email, nil
	}
}
