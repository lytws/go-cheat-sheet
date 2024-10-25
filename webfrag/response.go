package main

import (
	"encoding/json"
	"net/http"
)

// ToResponse 响应函数, 将数据转换到 http 响应中
type ToResponse[T any] func(http.ResponseWriter) error

func ToJsonResponse[T any](respCode int, data T) ToResponse[T] {
	return func(w http.ResponseWriter) error {
		w.Header().Set("Content-Type", "application/json")

		resp := struct {
			Code int `json:"code"`
			Data T   `json:"data"`
		}{
			Code: respCode,
			Data: data,
		}

		d, err := json.Marshal(resp)
		if err != nil {
			return err
		}

		_, err = w.Write(d)

		return err
	}
}

func ToJsonErrResponse[T any](respCode int, err error) ToResponse[T] {
	return func(w http.ResponseWriter) error {
		w.Header().Set("Content-Type", "application/json")

		resp := struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  respCode,
			Error: err.Error(),
		}

		d, err := json.Marshal(resp)
		if err != nil {
			return err
		}

		_, err = w.Write(d)

		return err
	}
}
