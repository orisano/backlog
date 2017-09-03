package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

const (
	version = "0.3.0"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  string

	apiKey string

	logger *log.Logger
	debug  bool
}

var defaultUserAgent = fmt.Sprintf("%s go-backlog/%s", runtime.Version(), version)

type options struct {
	UserAgent *string
	Logger    *log.Logger
	Debug     bool
}

type option func(*options)

func SetLogger(logger *log.Logger) option {
	return func(opt *options) {
		opt.Logger = logger
	}
}

func SetUserAgent(userAgent string) option {
	return func(opt *options) {
		opt.UserAgent = &userAgent
	}
}

func SetDebug(debug bool) option {
	return func(opt *options) {
		opt.Debug = debug
	}
}

func NewClient(urlStr, apiKey string, opts ...option) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("missing api key")
	}
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	var options options
	for _, opt := range opts {
		opt(&options)
	}

	logger := options.Logger
	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile)
	}
	userAgent := options.UserAgent
	if userAgent == nil {
		userAgent = &defaultUserAgent
	}

	return &Client{
		baseURL:    parsedURL,
		httpClient: http.DefaultClient,
		userAgent:  *userAgent,

		apiKey: apiKey,

		logger: logger,
		debug:  options.Debug,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	if ctx == nil {
		return nil, errors.New("nil context")
	}
	if len(method) == 0 {
		return nil, errors.New("missing method")
	}
	if len(spath) == 0 {
		return nil, errors.New("missing spath")
	}
	u := *c.baseURL
	u.Path = path.Join(u.Path, spath)

	param := u.Query()
	param.Set("apiKey", c.apiKey)
	u.RawQuery = param.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) do(req *http.Request, expected int, out interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expected {
		return errors.Errorf("unexpected status. expected %v, actual %v", expected, resp.Status)
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return errors.Wrap(err, "response decode failed")
		}
	}
	return nil
}

func (c *Client) get(ctx context.Context, spath string, expected int, out interface{}) error {
	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return errors.Wrap(err, "request construct failed")
	}
	if err := c.do(req, expected, out); err != nil {
		return errors.Wrap(err, "request failed")
	}
	return nil
}

func (c *Client) post(ctx context.Context, spath string, form url.Values, expected int, out interface{}) error {
	body := strings.NewReader(form.Encode())
	req, err := c.newRequest(ctx, http.MethodPost, spath, body)
	if err != nil {
		return errors.Wrap(err, "request construct failed")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err := c.do(req, expected, out); err != nil {
		return errors.Wrap(err, "request failed")
	}
	return nil
}
