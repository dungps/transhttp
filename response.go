package transhttp

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func RespondJSONFull(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	RespondJSONFullWithETag(w, httpStatusCode, payload, "")
}

func RespondJSONFullWithETag(w http.ResponseWriter, httpStatusCode int, payload interface{}, ETag string) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(ETag) > 0 {
		w.Header().Set("ETag", "W/\""+ETag+"\"")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}

func RespondJSON(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}

func RespondError(w http.ResponseWriter, httpStatusCode int, message string) {
	RespondJSON(w, httpStatusCode, map[string]string{"error": message})
}

func RespondJSONError(w http.ResponseWriter, payload interface{}) {
	var httpStatusCode = http.StatusInternalServerError
	if err, ok := payload.(error); ok {
		httpStatusCode = GetStatusCode(err)
	}
	RespondJSON(w, httpStatusCode, payload)
}

func RespondMessage(w http.ResponseWriter, httpStatusCode int, message string) {
	RespondJSON(w, httpStatusCode, map[string]string{"message": message})
}

func RespondMessageWithContentType(w http.ResponseWriter, httpStatusCode int, message string, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(message)))
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(message))
}

func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func RespondFile(w http.ResponseWriter, r *http.Request, fileName string) {
	http.ServeFile(w, r, fileName)
}
