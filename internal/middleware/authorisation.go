package authorisation

import (
	"errors"
	"net/http"

	"github.com/pylinho/go-api/api"
	"github.com/pylinho/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorisedError = errors.New("Invalid username or token.")

func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error 

		if username == "" || token == "" {
			
		}
	}
}