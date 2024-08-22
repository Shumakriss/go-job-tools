package handlers

import (
	"net/http"
	"strings"
)

type MethodTypeValidation struct {
	Valid        bool
	Msg          string
	StatusCode   int
	ExpectedType string
	ActualType   string
}

func NewMethodTypeValidation(expectedType string) MethodTypeValidation {
	v := MethodTypeValidation{}
	v.Valid = false
	v.Msg = "Method is not " + expectedType
	v.StatusCode = http.StatusMethodNotAllowed
	v.ExpectedType = expectedType
	return v
}

func (v *MethodTypeValidation) Validate() {
	v.Valid = v.ExpectedType == v.ActualType
}

func validateMethodType(r *http.Request, methodType string) MethodTypeValidation {
	v := NewMethodTypeValidation(methodType)
	v.ActualType = r.Method
	v.Validate()
	return v
}

type ContentTypeValidation struct {
	Valid        bool
	Msg          string
	StatusCode   int
	ExpectedType string
	ActualType   string
}

func NewContentTypeValidation(expectedType string) ContentTypeValidation {
	v := ContentTypeValidation{}
	v.Valid = false
	v.Msg = "Content-Type header is not " + expectedType
	v.StatusCode = http.StatusUnsupportedMediaType
	v.ExpectedType = expectedType
	return v
}

func (v *ContentTypeValidation) Validate() {
	v.Valid = v.ExpectedType == v.ActualType
}

func validateContentType(r *http.Request, contentType string) ContentTypeValidation {
	v := NewContentTypeValidation(contentType)

	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		v.ActualType = mediaType
	} else {
		v.ActualType = "UNKNOWN"
	}

	v.Validate()

	return v
}
