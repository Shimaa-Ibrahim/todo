package decorators

import (
	"net/http"
	"strconv"
)

var UserID uint

func IsAuthenticated(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := r.Cookie("user_id")
		if err == nil {
			id, _ := strconv.ParseUint(userID.Value, 10, 32)
			UserID = uint(id)
			endpoint(w, r)
		} else {
			http.Redirect(w, r, "/auth", http.StatusFound)
		}
	})
}
