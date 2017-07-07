package backlog

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"
)

const (
	version = "0.2.0"
)

type Client struct {
	URL    *url.URL
	client *http.Client

	apiKey string

	logger *log.Logger
}

type requestOption struct {
	params map[string]string
	body   map[string]string
}

func NewClient(urlStr, apiKey string, logger *log.Logger) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("missing api key")
	}
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}
	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	return &Client{
		URL:    parsedURL,
		client: http.DefaultClient,

		apiKey: apiKey,

		logger: logger,
	}, nil
}

var userAgent = fmt.Sprintf("orisano-backlog/%s (%s)", version, runtime.Version())

func (c *Client) newRequest(ctx context.Context, method, spath string, opt *requestOption) (*http.Request, error) {
	if ctx == nil {
		return nil, errors.New("nil context")
	}
	if len(method) == 0 {
		return nil, errors.New("missing method")
	}
	if len(spath) == 0 {
		return nil, errors.New("missing spath")
	}

	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	var r io.Reader
	if len(opt.body) != 0 {
		kv := make([]string, 0, len(opt.body))
		for k, v := range opt.body {
			kv = append(kv, fmt.Sprintf("%s=%s", k, v))
		}
		r = strings.NewReader(strings.Join(kv, "&"))
	}
	req, err := http.NewRequest(method, u.String(), r)
	if err != nil {
		return nil, err
	}

	values := req.URL.Query()
	values.Add("apiKey", c.apiKey)
	if len(opt.params) != 0 {
		for k, v := range opt.params {
			values.Add(k, v)
		}
	}
	req.URL.RawQuery = values.Encode()

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req = req.WithContext(ctx)

	return req, nil
}

func assertStatusCode(res *http.Response, expected int) error {
	if res.StatusCode != expected {
		return errors.Errorf("invalid status code: %s", res.Status)
	}
	return nil
}
