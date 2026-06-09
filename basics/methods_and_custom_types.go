package basics

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	HTTPMethodGet    HTTPMethod = "GET"
	HTTPMethodPost   HTTPMethod = "POST"
	HTTPMethodPut    HTTPMethod = "PUT"
	HTTPMethodDelete HTTPMethod = "DELET"
)

type HTTPMethod string

func (h HTTPMethod) String() string {
	return string(h)
}

type URLParams map[string][]string

func (u *URLParams) Add(name string, value string) {
	params := *u
	if _, ok := params[name]; ok {
		params[name] = append(params[name], value)
		return
	}
	params[name] = []string{value}
}

func (u *URLParams) Get(name string) []string {
	param, _ := (*u)[name]
	return param
}

// Não utilizar como referência pois não estamos fazendo a substituição correta
func (u *URLParams) String() string {
	params := *u
	formattedParams := strings.Builder{}
	for k, v := range params {
		if _, err := fmt.Fprintf(&formattedParams, "%s=%s&", k, strings.Join(v, ",")); err != nil {
			log.Fatalf("ERROR: %s", err)
		}
	}
	return strings.TrimSuffix(formattedParams.String(), "&")
}

type URL struct {
	host   string
	path   string
	params URLParams
}

func NewURL(host string, path string, params URLParams) *URL {
	return &URL{
		host:   host,
		path:   path,
		params: params,
	}
}

func (u URL) String() string {
	return fmt.Sprintf("%s%s?%s", u.host, u.path, u.params.String())
}

type Request struct {
	method HTTPMethod
	url    URL
	client http.Client
}

func NewRequest(method HTTPMethod, url URL) *Request {
	return &Request{
		url:    url,
		method: method,
	}
}

func (r *Request) Execute() {
	log.Printf("INFO: %s %s", r.method, r.url)

	req, err := http.NewRequest(r.method.String(), r.url.String(), nil)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	res, err := r.client.Do(req)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	log.Printf("INFO: response.status: %d, response.body: %s\n", res.StatusCode, string(body))
}
