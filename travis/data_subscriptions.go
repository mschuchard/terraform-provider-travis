package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
)

// subscriptions data declaration and schema
func dataSubscriptions() *schema.Resource {
  return &schema.Resource{
    Read: subscriptionsRead,
  }
}

// read function for the data
func subscriptionsRead(meta interface{}) error {

}
