package main

import "net/http"

type OneParamController[T any, U any] func(T) ToResponse[U]

func WrapOneParam[T any, U any](
	fn OneParamController[T, U],
	ext Extractor[T],
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := ext.FromRequest(r)
		if err != nil {
			ext.ExtractErrorResponse(err)(w)
			return
		}

		result := fn(t)

		if err := result(w); err != nil {
			// log error
			return
		}
	}
}

type TwoParamController[T any, U any, V any] func(T, U) ToResponse[V]

func WrapTwoParam[T any, U any, V any](
	fn TwoParamController[T, U, V],
	ext1 Extractor[T],
	ext2 Extractor[U],
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := ext1.FromRequest(r)
		if err != nil {
			ext1.ExtractErrorResponse(err)(w)
			return
		}

		u, err := ext2.FromRequest(r)
		if err != nil {
			ext2.ExtractErrorResponse(err)(w)
			return
		}

		result := fn(t, u)

		if err := result(w); err != nil {
			// log
			return
		}
	}
}
