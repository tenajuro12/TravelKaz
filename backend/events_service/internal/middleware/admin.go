package middleware

import (
	"net/http"
)

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from cookie
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "No authorization token", http.StatusUnauthorized)
			return
		}

		// Validate token with auth service
		authServiceURL := "http://localhost:8082/validate-admin"
		req, err := http.NewRequest("GET", authServiceURL, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Cookie", "session_token="+cookie.Value)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()

		// If valid, proceed
		next.ServeHTTP(w, r)
	})
}
