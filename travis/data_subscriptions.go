package travis

import (
  "fmt"
  "encoding/json"

  "github.com/hashicorp/terraform/helper/schema"
)

// subscriptions data declaration and schema
func dataSubscriptions() *schema.Resource {
  return &schema.Resource{
    Read: subscriptionsRead,

    Schema: map[string]*schema.Schema {
      "subscriptions": {
        Type:     schema.TypeSet,
        Computed: true,
      },
    },
  }
}

// read function for the data
func subscriptionsRead(data *schema.ResourceData, meta interface{}) error {
  // construct endpoint
  endpoint := ":9293/subscriptions"

  // construct travisOpts
  opts := &travisOpts {
    endpoint: endpoint,
  }

  // receive response body
  responseBody, err := apiClient(opts)

  // error handle
  if err != nil {
    fmt.Errorf("Error interacting with Travis API.")
  }

  // convert response json to map
  var responseMap map[string][]string
  err = json.Unmarshal(responseBody, &responseMap)

  // error handle
  if err != nil {
    fmt.Errorf("Invalid JSON response from Travis.")
  // verify subscriptions returned from travis
  } else if _, exists := responseMap["subscriptions"]; exists {
    // set subscriptions attribute and return
    data.Set("subscriptions", responseMap["subscriptions"])
    return nil
  } else {
    fmt.Errorf("Subscriptions not found in response from Travis.")
  }

  return err
}
