package backlog

import (
	"net/url"
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
