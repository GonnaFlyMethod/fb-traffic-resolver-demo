package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readRequestBody(r *http.Request) ([]byte, error) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer func(r *http.Request) {
		if err = r.Body.Close(); err != nil {
			fmt.Println("failed to close request body")
		}
	}(r)

	r.Body = io.NopCloser(bytes.NewBuffer(content))
	return content, nil
}

type ErrResponse struct {
	Err string `json:"err"`
}

func writeInternalErrResponse(w http.ResponseWriter) {
	errResponse := ErrResponse{Err: "something went wrong on our end"}
	writeResponse(w, http.StatusInternalServerError, errResponse)
}

func writeResponse(w http.ResponseWriter, statusCode int, responseBody interface{}) {
	w.WriteHeader(statusCode)

	marshalledResponseBody, err := json.Marshal(responseBody)
	if err != nil {
		fmt.Println("can't marshall response body")
		return
	}

	if _, err = w.Write(marshalledResponseBody); err != nil {
		fmt.Println(err)
		fmt.Println("can't write response body")
		return
	}
}
