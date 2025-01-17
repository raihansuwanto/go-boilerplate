package errors

import (
	"encoding/json"
	"fmt"
	"sort"
)

type (
	Errors map[string]error
)

func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}

	keys := []string{}
	for key := range es {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	s := ""
	for i, key := range keys {
		if i > 0 {
			s += "; "
		}
		if errs, ok := es[key].(Errors); ok {
			s += fmt.Sprintf("%v: (%v)", key, errs)
		} else {
			s += fmt.Sprintf("%v: %v", key, es[key].Error())
		}
	}
	return s + "."
}

// MarshalJSON converts the Errors into a valid JSON.
func (es Errors) MarshalJSON() ([]byte, error) {
	errs := map[string]interface{}{}
	for key, err := range es {
		if ms, ok := err.(json.Marshaler); ok {
			errs[key] = ms
		} else {
			errs[key] = err.Error()
		}
	}
	return json.Marshal(errs)
}

func (es Errors) WithMessage(message string) Errors {
	es["message"] = fmt.Errorf(message)
	return es
}

func NewInternalSystemError() Errors {
	errs := Errors{}

	errs["message"] = fmt.Errorf("internal system error")
	errs["code"] = fmt.Errorf("500")
	errs["messageKey"] = fmt.Errorf("internal_system_error")

	return errs
}

func NewDataNotFoundError() Errors {
	errs := Errors{}

	errs["message"] = fmt.Errorf("data not found")
	errs["code"] = fmt.Errorf("404")
	errs["messageKey"] = fmt.Errorf("data_not_found")

	return errs
}

func NewBadRequestError() Errors {
	errs := Errors{}

	errs["message"] = fmt.Errorf("bad request")
	errs["code"] = fmt.Errorf("400")
	errs["messageKey"] = fmt.Errorf("bad_request")

	return errs
}
