package backlog

import (
	"io"
	"net/url"
	"strings"
)

func mergeValues(a, b url.Values) url.Values {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	r := a
	for k, vs := range b {
		for _, v := range vs {
			r.Add(k, v)
		}
	}
	return r
}

func encodeForm(form url.Values) io.Reader {
	return strings.NewReader(form.Encode())
}
