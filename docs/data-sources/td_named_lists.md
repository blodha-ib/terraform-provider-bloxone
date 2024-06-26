---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_td_named_lists Data Source - terraform-provider-bloxone"
subcategory: ""
description: |-
  Retrieves information about existing Named Lists.
---

# bloxone_td_named_lists (Data Source)

Retrieves information about existing Named Lists.

## Example Usage

```terraform
# Get Named Lists filtered by tag
data "bloxone_td_named_lists" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all Named Lists
data "bloxone_td_named_lists" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Optional:

- `tags` (Map of String) Tags associated with this Named List

Read-Only:

- `confidence_level` (String) The confidence level for a custom list. The possible values are ["LOW", "MEDIUM", "HIGH"]
- `created_time` (String) The time when this Named List object was created.
- `description` (String) The brief description for the named list.
- `id` (Number) The Named List object identifier.
- `item_count` (Number) The number of items in this named list.
- `name` (String) The name of the named list.
- `policies` (List of String) The list of the security policy names with which the named list is associated.
- `threat_level` (String) The threat level for a custom list. The possible values are ["INFO", "LOW", "MEDIUM", "HIGH"]
- `type` (String) The type of the named list, that can be "custom_list", "threat_insight", "fast_flux", "dga", "dnsm", "threat_insight_nde", "default_allow", "default_block".
- `updated_time` (String) The time when this Named List object was last updated.
