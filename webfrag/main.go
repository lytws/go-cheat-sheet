package main

import (
	"fmt"
	"net/http"
)

type P struct {
	A string `json:"a"`
}

type R struct {
	Ret string `json:"ret"`
}

func Ctrl(p P) ToResponse[R] {
	fmt.Println(p.A)

	if p.A == "" {
		return ToJsonErrResponse[R](400, fmt.Errorf("a is empty"))
	}

	return ToJsonResponse(200, R{Ret: fmt.Sprintf("ok, got: %s", p.A)})
}

func Ctrl2(query map[string]string, p P) ToResponse[R] {
	fmt.Println(query)

	if p.A == "" {
		return ToJsonErrResponse[R](400, fmt.Errorf("a is empty"))
	}

	return ToJsonResponse(200, R{Ret: fmt.Sprintf("ok, got: %s, %+v", p.A, query)})
}

func main() {
	http.HandleFunc("POST /test", WrapOneParam(Ctrl, JsonBodyExtractor[P]()))
	http.HandleFunc("POST /test2", WrapTwoParam(
		Ctrl2,
		QueryExtractor[map[string]string](),
		JsonBodyExtractor[P]()),
	)

	http.ListenAndServe(":8080", nil)
}
