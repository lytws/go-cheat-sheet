package main

import (
	"encoding/json"
	"net/http"
)

type Extractor[T any] struct {
	FromRequest          ExtractFromRequest[T]
	ExtractErrorResponse func(error) ToResponse[error]
}

type ExtractFromRequest[T any] func(*http.Request) (T, error)

func JsonBodyExtractor[T any]() Extractor[T] {

	fromRequest := func(r *http.Request) (T, error) {

		var t T

		err := json.NewDecoder(r.Body).Decode(&t)
		return t, err
	}

	return Extractor[T]{
		FromRequest: fromRequest,
		ExtractErrorResponse: func(err error) ToResponse[error] {
			return ToJsonErrResponse[error](http.StatusBadRequest, err)
		},
	}
}

func QueryExtractor[T any]() Extractor[T] {
	fromRequest := func(r *http.Request) (T, error) {

		var t T

		val := map[string]string{}
		for k := range r.URL.Query() {
			val[k] = r.URL.Query().Get(k)
		}

		d, _ := json.Marshal(val)
		err := json.Unmarshal(d, &t)
		return t, err
	}

	return Extractor[T]{
		FromRequest: fromRequest,
		ExtractErrorResponse: func(err error) ToResponse[error] {
			return ToJsonErrResponse[error](http.StatusBadRequest, err)
		},
	}
}
