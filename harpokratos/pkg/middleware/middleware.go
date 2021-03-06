package middleware

import (
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/kpango/glg"
	"gopkg.in/oauth2.v3/utils/uuid"
	"harpokratos/pkg/models"
	"net/http"
	"strconv"
)

type Adapter func(http.HandlerFunc) http.HandlerFunc

// Iterate over adapters and run them one by one
func Adapt(h http.HandlerFunc, adapters ...Adapter) http.HandlerFunc {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func LogRequestDetails() Adapter {
	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
				glg.Infof("%s route %s", r.Method, r.URL.Path)
				f(w, r)
			}
		}
}

// middleware to validate proper methods
func ValidateRestMethod(method string) Adapter {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				var err models.MethodError
				e := models.MethodMessages{method, "Method " + r.Method + " not allowed at this endpoint"}
				err = models.MethodError{models.ErrorModel{CreateGUID()}, append(err.Messages, e)}
				glg.Errorf("%s %s", r.URL.Path, e.Message)
				ResponseWithJson(w, err)
				return
			}
			f(w, r)
		}
	}
}

// ResponseWithJson returns formed JSON
func ResponseWithJson(w http.ResponseWriter, payload interface{}) {
	code := 500

	switch payload.(type) {
	case models.ResultModel:
		code = 200
	case *vault.Secret:
		code = 200
	case map[string]interface{}:
		code = 200
	case models.ValidationError:
		code = 400
	case models.MethodError:
		code = 405
	default:
		code = 500
	}

	response, _ := json.Marshal(payload)
	resp := string(response)
	if code != 200 {
		glg.Errorf("responseCode: %s body: %s", strconv.Itoa(code), resp)
	} else {
		glg.Infof("responseCode: %s body: %s", strconv.Itoa(code), resp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(resp))
}

// creates a Guid for error tracing
func CreateGUID() string {
	b, _ := uuid.NewRandom()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}