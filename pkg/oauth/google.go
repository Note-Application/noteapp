package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"noteapp/internal/models"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google OAuth2 config
var googleOauthConfig = oauth2.Config{
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

// GoogleUser represents a user fetched from Google
type GoogleUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfilePic string `json:"picture"`
}

// GetGoogleUserInfo exchanges the code for a token and fetches user info
func GetGoogleUserInfo(code string) (models.User, error) {
	// Exchange the code for a token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return models.User{}, errors.New("failed to exchange token")
	}

	// Get user info from Google API
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + token.AccessToken)
	if err != nil {
		return models.User{}, errors.New("failed to get user info")
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, errors.New("failed to read response body")
	}

	// Parse the user data
	var googleUser GoogleUser
	if err := json.Unmarshal(body, &googleUser); err != nil {
		return models.User{}, errors.New("failed to parse user info")
	}

	// Map the Google user data to our internal user model
	user := models.User{
		Email:      googleUser.Email,
		Name:       googleUser.Name,
		ProfilePic: googleUser.ProfilePic,
	}

	return user, nil
}
