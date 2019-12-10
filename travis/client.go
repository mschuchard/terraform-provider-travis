package travis

import (
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
)

// struct for passing opts from provider to resources
type travisOpts struct {
  token      string
  commercial bool
  method     string
  endpoint   string
  headers    map[string]string
  body       []byte
}

// primary helper function for client connections to api endpoint
func apiClient(opts *travisOpts) (string, error) {
  // construct url
  var url string
  if opts.commercial {
    url = "https://api.travis-ci.com/repo/" + opts.endpoint
  } else {
    url = "https://api.travis-ci.org/repo/" + opts.endpoint
  }

  // initialize http client
  client := &http.Client{}

  // body into bytes.Buffer pointer
  var bodyBuffer *bytes.Buffer
  bodyBuffer = bytes.NewBuffer(opts.body)

  // body always request body or nil for reads
  request, err := http.NewRequest(opts.method, url, bodyBuffer)

  // error handle
  if err != nil {
    fmt.Errorf("Error constructing REST API request.")
  }

  // append base headers
  request.Header.Add("Content-Type", "application/json")
  request.Header.Add("Accept", "application/json")
  request.Header.Add("Travis-API-Version", "3")
  request.Header.Add("Authorization", opts.token)

  // append extra headers; TODO: block conflict with above
  for headerKey, headerValue := range opts.headers {
    request.Header.Add(headerKey, headerValue)
  }

  // initiate request for response
  response, err := client.Do(request)

  // error handle
  if err != nil {
    fmt.Errorf("Error making request to Travis API.")
  }

  // read body of response and return it
  defer response.Body.Close()
  responseBody, err := ioutil.ReadAll(response.Body)

  return string(responseBody), err
}
