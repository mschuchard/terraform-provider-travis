package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
)

// subscriptions data declaration and schema
func dataSubscriptions() *schema.Resource {
  return &schema.Resource{
    Read: subscriptionsRead,

    Schema: map[string]*schema.Schema {
      "subscriptions": {
        Type:     schema.TypeArray,
        Computed: true,
      },
    },
  }
}

// read function for the data
func subscriptionsRead(meta interface{}) error {
  // construct endpoint
  endpoint = ":9293/subscriptions"

  // construct travisOpts
  opts := &travisOpts {
    endpoint: endpoint,
  }

  // receive response body
  responseBody, err := apiClient(opts)

  // error handle
  if err != nil {
    fmt.Errorf("Error interacting with Travis API.")
  // verify subscriptions returned from travis
  } else if _, exists := responseBody["subscriptions"]; exists {
    // set subscriptions attribute and return
    data.set("subscriptions", responseBody["subscriptions"])
    return nil
  } else {
    fmt.Errorf("Subscriptions not found in response from Travis.")
  }
}
