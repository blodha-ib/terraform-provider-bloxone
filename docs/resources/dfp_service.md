---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dfp_service Resource - terraform-provider-bloxone"
subcategory: "Threat Defense"
description: |-
  Manages a DNS Forwarding Proxy Service.
---

# bloxone_dfp_service (Resource)

Manages a DNS Forwarding Proxy Service.

## Example Usage

```terraform
data "bloxone_infra_services" "example" {
  name = "example_dfp_service"
}

resource "bloxone_td_internal_domain_list" "example_list" {
  name             = "example_internal_domain_list"
  internal_domains = ["example.domain.com"]
}

resource "bloxone_dfp_service" "example" {
  service_id = data.bloxone_infra_services.example.results.0.id

  # Other optional fields
  internal_domain_lists = [bloxone_td_internal_domain_list.example_list.id]
  resolvers_all = [
    {
      address     = "1.1.1.1"
      is_fallback = true
      is_local    = false
      protocols   = ["DOT", "DO53"]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `service_id` (String) The On-Prem Application Service identifier. For internal Use only

### Optional

- `forwarding_policy` (String) Policy Identifier for DNS Forwarding Proxy.
- `internal_domain_lists` (List of Number) The list of internal domains list IDs that are associated with this DFP
- `net_addr_policy_ids` (Attributes List) List of network-address-scoped security policy assignments (see [below for nested schema](#nestedatt--net_addr_policy_ids))
- `resolvers_all` (Attributes List) List of DNS resolvers (see [below for nested schema](#nestedatt--resolvers_all))

### Read-Only

- `created_time` (String) The time when this DNS Forwarding Proxy object was created.
- `elb_ip_list` (List of String) The list of internal or local DNS servers' IPv4 or IPv6 addresses that are used as ELB IPs.
- `host` (Attributes List) host information. For internal Use only. (see [below for nested schema](#nestedatt--host))
- `id` (Number) The DNS Forwarding Proxy object identifier.
- `name` (String) The name of the DNS Forwarding Proxy.
- `ophid` (String) The On-Prem Host identifier.
- `policy_id` (Number) The identifier of the security policy with which the DNS Forwarding Proxy is associated.
- `pop_region_id` (Number) Point of Presence (PoP) region
- `service_name` (String) The On-Prem Application Service name. For internal Use only
- `site_id` (String) The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.
- `updated_time` (String) The time when this DNS Forwarding Proxy object was last updated.

<a id="nestedatt--net_addr_policy_ids"></a>
### Nested Schema for `net_addr_policy_ids`

Optional:

- `addr_net` (String) Network address in IPv4 CIDR (address/bitmask length) string format
- `policy_id` (Number) Identifier of the security policy associated with this address block


<a id="nestedatt--resolvers_all"></a>
### Nested Schema for `resolvers_all`

Optional:

- `address` (String) Address that can be used as resolver
- `is_fallback` (Boolean) Mark it true to set default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable.
- `is_local` (Boolean) Mark it true to set internal or local DNS servers' IPv4 or IPv6 addresses that are used as DNS resolvers
- `protocols` (List of String) The list of DNS resolver communication protocols.


<a id="nestedatt--host"></a>
### Nested Schema for `host`

Optional:

- `legacy_host_id` (Number) The DNS Forwarding Proxy legacy ID object identifier.

Read-Only:

- `name` (String) The name of the DNS Forwarding Proxy.
- `ophid` (String) The On-Prem Host identifier.
- `site_id` (String) The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.