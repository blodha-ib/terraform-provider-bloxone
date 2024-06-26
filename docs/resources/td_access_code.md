---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_td_access_code Resource - terraform-provider-bloxone"
subcategory: ""
description: |-
  Manages an access code.
---

# bloxone_td_access_code (Resource)

Manages an access code.

## Example Usage

```terraform
resource "bloxone_td_access_code" "example" {
  name       = "example_access_code"
  activation = timestamp()
  expiration = timeadd(timestamp(), "24h")
  rules = [
    {
      data = "terraform_test",
      type = "custom_list"
    }
  ]
  # Other optional fields
  description = "Example Access Code"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `activation` (String) The time when the Bypass Code object was activated.
- `expiration` (String) The time when the Bypass Code object expires.
- `name` (String) The name of Bypass Code
- `rules` (Attributes List) The list of selected security rules (see [below for nested schema](#nestedatt--rules))

### Optional

- `description` (String) The brief description for an access code.
- `policy_ids` (List of Number) The list of SecurityPolicy object identifiers.

### Read-Only

- `access_key` (String) Auto generated unique Bypass Code value
- `created_time` (String) The time when the Bypass Code object was created.
- `id` (String) The resource identifier.
- `updated_time` (String) The time when the Bypass Code object was last updated.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `action` (String) The action to be used in the rule.
- `data` (String) The data to be used in the rule.
- `description` (String) The brief description of the rule.
- `redirect_name` (String) The name of the redirect to be used in the rule.
- `type` (String) The type of the rule.
