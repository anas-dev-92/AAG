package froms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values        // which hold all the value in giving url in our case our page
	Errors     errors // our errors type we created in errors file
}

// create func to make new form and get the vlaues from it
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// make function to check the vlaues we get from the above function is valid or not
func (f *Form) Has(filed string, r *http.Request) bool {
	x := r.Form.Get(filed)
	if x == "" {
		return false
	}
	return true
}
