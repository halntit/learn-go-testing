package main

import (
	"net/url"
	"strings"
)

type errors map[string][]string
type Form struct {
	Data   url.Values
	Errors errors
}

func (e errors) Get(field string) string {
	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]
}

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}

func (form *Form) Has(field string) bool {
	x := form.Data.Get(field)
	return x != ""
}

func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (form *Form) Check(ok bool, key, message string) {
	if !ok {
		form.Errors.Add(key, message)
	}
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}
