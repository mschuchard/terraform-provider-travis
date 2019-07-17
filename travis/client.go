package travis

import (
  "net/http"
  "io/ioutil"
)

// primary helper function for client connections to api endpoint
func apiClient(method string, endpoint string, headers map[string]string, body string) (string, error) {
  // construct url
  url := "https://api.travis-ci.org/repo/" + endpoint

  // initialize http client
  client := &http.Client{}
  request, err := client.NewRequest(method, url, nil, body)

  // append base headers
  client.Header.Add("Content-Type", "application/json")
  client.Header.Add("Accept", "application/json")
  client.Header.Add("Travis-API-Version", "3")
  client.Header.Add("Authorization", travisOpts.token)

  // append extra headers; TODO: block conflict with above
  for headerKey, headerValue := range headers {
    client.Header.Add(headerKey, headerValue)
  }

  // initiate request for response
  response, err := client.Do(request)

  // TODO: err handle

  // read body of response and return it
  defer response.Body.Close()
  responseBody, err := ioutil.ReadAll(response.Body)

  return responseBody
}
