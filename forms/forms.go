package froms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values        // which hold all the value in giving url in our case our page
	Errors     errors // our errors type we created in errors file
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
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
		f.Errors.Add(filed, "this field can't be empty")
		return false
	}
	return true
}

// check the form var and return true if there is no errors
