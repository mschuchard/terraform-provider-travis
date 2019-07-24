package travis

import (
  "bytes"
  "net/http"
  "io/ioutil"
)

// primary helper function for client connections to api endpoint
func apiClient(method string, endpoint string, headers map[string]string, body string) (string, error) {
  // construct url
  url := "https://api.travis-ci.org/repo/" + endpoint

  // initialize http client
  client := &http.Client{}
  // ternary to convert body into bytes.Buffer pointer
  var bodyBuffer *bytes.Buffer
  if body != nil {
    bodyBuffer = bytes.NewBuffer([]byte(body))
  } else {
    bodyBuffer = nil
  }
  // body always request body or nil for reads
  request, err := http.NewRequest(method, url, bodyBuffer)

  // append base headers
  request.Header.Add("Content-Type", "application/json")
  request.Header.Add("Accept", "application/json")
  request.Header.Add("Travis-API-Version", "3")
  request.Header.Add("Authorization", travisOpts.token)

  // append extra headers; TODO: block conflict with above
  for headerKey, headerValue := range headers {
    request.Header.Add(headerKey, headerValue)
  }

  // initiate request for response
  response, err := client.Do(request)

  // TODO: err handle

  // read body of response and return it
  defer response.Body.Close()
  responseBody, err := ioutil.ReadAll(response.Body)

  return responseBody
}
