package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// Session store - You can configure it to use a secure session store
var store = sessions.NewCookieStore([]byte("your-secret-key"))

// SetSession stores the user ID in the session
func SetSession(w http.ResponseWriter, r *http.Request, userID int) {
	session, _ := store.Get(r, "user-session")
	session.Values["user_id"] = userID
	err := session.Save(r, w)
	if err != nil {
		log.Println("Error saving session:", err)
	}
}

// ClearSession clears the user session
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		log.Println("Error clearing session:", err)
	}
}

// GetUserID retrieves the user ID from the session
func GetUserID(r *http.Request) int {
	session, _ := store.Get(r, "user-session")
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		// If user is not authenticated, you can handle this case (e.g., return an error or redirect)
		return 0 // Or you can return an error here or handle as per your app's flow
	}
	return userID
}
