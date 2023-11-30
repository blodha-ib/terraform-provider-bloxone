---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dns_auth_zone Resource - terraform-provider-bloxone"
subcategory: ""
description: |-
  
---

# bloxone_dns_auth_zone (Resource)



## Example Usage

```terraform
resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "tf-acc-test.com."
  primary_type = "cloud"

  # Other optional fields
  name    = "example_auth_zone"
  comment = "Example of an Authoritative Zone"
  tags = {
    site = "Test Site"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fqdn` (String) Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.
- `primary_type` (String) Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a BloxOne DDI host.

### Optional

- `comment` (String) Optional. Comment for zone configuration.
- `disabled` (Boolean) Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
- `external_primaries` (Attributes List) Optional. DNS primaries external to BloxOne DDI. Order is not significant. (see [below for nested schema](#nestedatt--external_primaries))
- `external_secondaries` (Attributes List) DNS secondaries external to BloxOne DDI. Order is not significant. (see [below for nested schema](#nestedatt--external_secondaries))
- `gss_tsig_enabled` (Boolean) _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.
- `inheritance_sources` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources))
- `initial_soa_serial` (Number) On-create-only. SOA serial is allowed to be set when the authoritative zone is created.
- `internal_secondaries` (Attributes List) Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant. (see [below for nested schema](#nestedatt--internal_secondaries))
- `notify` (Boolean) Also notify all external secondary DNS servers if enabled.  Defaults to _false_.
- `nsgs` (List of String) The resource identifier.
- `query_acl` (Attributes List) Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. (see [below for nested schema](#nestedatt--query_acl))
- `tags` (Map of String) Tagging specifics.
- `transfer_acl` (Attributes List) Optional. Clients must match this ACL to receive zone transfers. (see [below for nested schema](#nestedatt--transfer_acl))
- `update_acl` (Attributes List) Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. (see [below for nested schema](#nestedatt--update_acl))
- `use_forwarders_for_subzones` (Boolean) Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.
- `view` (String) The resource identifier.

### Read-Only

- `created_at` (String) Time when the object has been created.
- `external_providers` (Attributes List) list of external providers for the auth zone. (see [below for nested schema](#nestedatt--external_providers))
- `id` (String) The resource identifier.
- `inheritance_assigned_hosts` (Attributes List) The list of the inheritance assigned hosts of the object. (see [below for nested schema](#nestedatt--inheritance_assigned_hosts))
- `mapped_subnet` (String) Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.
- `mapping` (String) Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward.
- `parent` (String) The resource identifier.
- `protocol_fqdn` (String) Zone FQDN in punycode.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.
- `warnings` (Attributes List) The list of an auth zone warnings. (see [below for nested schema](#nestedatt--warnings))
- `zone_authority` (Attributes) (see [below for nested schema](#nestedatt--zone_authority))

<a id="nestedatt--external_primaries"></a>
### Nested Schema for `external_primaries`

Required:

- `type` (String) Allowed values: * _nsg_, * _primary_.

Optional:

- `address` (String) Optional. Required only if _type_ is _server_. IP Address of nameserver.
- `fqdn` (String) Optional. Required only if _type_ is _server_. FQDN of nameserver.
- `nsg` (String) The resource identifier.
- `tsig_enabled` (Boolean) Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--external_primaries--tsig_key))

Read-Only:

- `protocol_fqdn` (String) FQDN of nameserver in punycode.

<a id="nestedatt--external_primaries--tsig_key"></a>
### Nested Schema for `external_primaries.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.



<a id="nestedatt--external_secondaries"></a>
### Nested Schema for `external_secondaries`

Required:

- `address` (String) IP Address of nameserver.
- `fqdn` (String) FQDN of nameserver.

Optional:

- `stealth` (Boolean) If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_
- `tsig_enabled` (Boolean) If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.  Default: _false_
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--external_secondaries--tsig_key))

Read-Only:

- `protocol_fqdn` (String) FQDN of nameserver in punycode.

<a id="nestedatt--external_secondaries--tsig_key"></a>
### Nested Schema for `external_secondaries.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.



<a id="nestedatt--inheritance_sources"></a>
### Nested Schema for `inheritance_sources`

Optional:

- `gss_tsig_enabled` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--gss_tsig_enabled))
- `notify` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--notify))
- `query_acl` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--query_acl))
- `transfer_acl` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--transfer_acl))
- `update_acl` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--update_acl))
- `use_forwarders_for_subzones` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--use_forwarders_for_subzones))
- `zone_authority` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority))

<a id="nestedatt--inheritance_sources--gss_tsig_enabled"></a>
### Nested Schema for `inheritance_sources.gss_tsig_enabled`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Boolean) The inherited value.


<a id="nestedatt--inheritance_sources--notify"></a>
### Nested Schema for `inheritance_sources.notify`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Boolean) The inherited value.


<a id="nestedatt--inheritance_sources--query_acl"></a>
### Nested Schema for `inheritance_sources.query_acl`

Optional:

- `action` (String) Optional. Inheritance setting for a field. Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) Human-readable display name for the object referred to by _source_.
- `value` (Attributes List) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--query_acl--value))

<a id="nestedatt--inheritance_sources--query_acl--value"></a>
### Nested Schema for `inheritance_sources.query_acl.value`

Required:

- `access` (String) Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
- `element` (String) Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.

Optional:

- `acl` (String) The resource identifier.
- `address` (String) Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--query_acl--value--tsig_key))

<a id="nestedatt--inheritance_sources--query_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.query_acl.value.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.




<a id="nestedatt--inheritance_sources--transfer_acl"></a>
### Nested Schema for `inheritance_sources.transfer_acl`

Optional:

- `action` (String) Optional. Inheritance setting for a field. Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) Human-readable display name for the object referred to by _source_.
- `value` (Attributes List) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--transfer_acl--value))

<a id="nestedatt--inheritance_sources--transfer_acl--value"></a>
### Nested Schema for `inheritance_sources.transfer_acl.value`

Required:

- `access` (String) Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
- `element` (String) Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.

Optional:

- `acl` (String) The resource identifier.
- `address` (String) Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--transfer_acl--value--tsig_key))

<a id="nestedatt--inheritance_sources--transfer_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.transfer_acl.value.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.




<a id="nestedatt--inheritance_sources--update_acl"></a>
### Nested Schema for `inheritance_sources.update_acl`

Optional:

- `action` (String) Optional. Inheritance setting for a field. Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) Human-readable display name for the object referred to by _source_.
- `value` (Attributes List) Inherited value. (see [below for nested schema](#nestedatt--inheritance_sources--update_acl--value))

<a id="nestedatt--inheritance_sources--update_acl--value"></a>
### Nested Schema for `inheritance_sources.update_acl.value`

Required:

- `access` (String) Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
- `element` (String) Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.

Optional:

- `acl` (String) The resource identifier.
- `address` (String) Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--update_acl--value--tsig_key))

<a id="nestedatt--inheritance_sources--update_acl--value--tsig_key"></a>
### Nested Schema for `inheritance_sources.update_acl.value.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.




<a id="nestedatt--inheritance_sources--use_forwarders_for_subzones"></a>
### Nested Schema for `inheritance_sources.use_forwarders_for_subzones`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Boolean) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority"></a>
### Nested Schema for `inheritance_sources.zone_authority`

Optional:

- `default_ttl` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--default_ttl))
- `expire` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--expire))
- `mname_block` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--mname_block))
- `negative_ttl` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--negative_ttl))
- `protocol_rname` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--protocol_rname))
- `refresh` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--refresh))
- `retry` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--retry))
- `rname` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--rname))

<a id="nestedatt--inheritance_sources--zone_authority--default_ttl"></a>
### Nested Schema for `inheritance_sources.zone_authority.default_ttl`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Number) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority--expire"></a>
### Nested Schema for `inheritance_sources.zone_authority.expire`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Number) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority--mname_block"></a>
### Nested Schema for `inheritance_sources.zone_authority.mname_block`

Optional:

- `action` (String) Defaults to _inherit_.
- `source` (String) The resource identifier.
- `value` (Attributes) (see [below for nested schema](#nestedatt--inheritance_sources--zone_authority--mname_block--value))

Read-Only:

- `display_name` (String) Human-readable display name for the object referred to by _source_.

<a id="nestedatt--inheritance_sources--zone_authority--mname_block--value"></a>
### Nested Schema for `inheritance_sources.zone_authority.mname_block.display_name`

Optional:

- `mname` (String) Defaults to empty.
- `use_default_mname` (Boolean) Optional. Use default value for master name server.  Defaults to true.

Read-Only:

- `protocol_mname` (String) Optional. Master name server in punycode.  Defaults to empty.



<a id="nestedatt--inheritance_sources--zone_authority--negative_ttl"></a>
### Nested Schema for `inheritance_sources.zone_authority.negative_ttl`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Number) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority--protocol_rname"></a>
### Nested Schema for `inheritance_sources.zone_authority.protocol_rname`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (String) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority--refresh"></a>
### Nested Schema for `inheritance_sources.zone_authority.refresh`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Number) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority--retry"></a>
### Nested Schema for `inheritance_sources.zone_authority.retry`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (Number) The inherited value.


<a id="nestedatt--inheritance_sources--zone_authority--rname"></a>
### Nested Schema for `inheritance_sources.zone_authority.rname`

Optional:

- `action` (String) The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.
- `source` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the object referred to by _source_.
- `value` (String) The inherited value.




<a id="nestedatt--internal_secondaries"></a>
### Nested Schema for `internal_secondaries`

Required:

- `host` (String) The resource identifier.


<a id="nestedatt--query_acl"></a>
### Nested Schema for `query_acl`

Required:

- `access` (String) Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
- `element` (String) Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.

Optional:

- `acl` (String) The resource identifier.
- `address` (String) Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--query_acl--tsig_key))

<a id="nestedatt--query_acl--tsig_key"></a>
### Nested Schema for `query_acl.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.



<a id="nestedatt--transfer_acl"></a>
### Nested Schema for `transfer_acl`

Required:

- `access` (String) Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
- `element` (String) Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.

Optional:

- `acl` (String) The resource identifier.
- `address` (String) Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--transfer_acl--tsig_key))

<a id="nestedatt--transfer_acl--tsig_key"></a>
### Nested Schema for `transfer_acl.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.



<a id="nestedatt--update_acl"></a>
### Nested Schema for `update_acl`

Required:

- `access` (String) Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
- `element` (String) Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.

Optional:

- `acl` (String) The resource identifier.
- `address` (String) Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--update_acl--tsig_key))

<a id="nestedatt--update_acl--tsig_key"></a>
### Nested Schema for `update_acl.tsig_key`

Optional:

- `algorithm` (String) TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.
- `comment` (String) Comment for TSIG key.
- `key` (String) The resource identifier.
- `name` (String) TSIG key name, FQDN.
- `secret` (String, Sensitive) TSIG key secret, base64 string.

Read-Only:

- `protocol_name` (String) TSIG key name in punycode.



<a id="nestedatt--external_providers"></a>
### Nested Schema for `external_providers`

Optional:

- `name` (String) The name of the external provider.
- `type` (String) Defines the type of external provider. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform.

Read-Only:

- `id` (String) The identifier of the external provider.


<a id="nestedatt--inheritance_assigned_hosts"></a>
### Nested Schema for `inheritance_assigned_hosts`

Optional:

- `host` (String) The resource identifier.

Read-Only:

- `display_name` (String) The human-readable display name for the host referred to by _ophid_.
- `ophid` (String) The on-prem host ID.


<a id="nestedatt--warnings"></a>
### Nested Schema for `warnings`

Optional:

- `message` (String) Warning message.
- `name` (String) Name of a warning.


<a id="nestedatt--zone_authority"></a>
### Nested Schema for `zone_authority`

Optional:

- `default_ttl` (Number) Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).  Defaults to 28800.
- `expire` (Number) Optional. ZoneAuthority expire time in seconds.  Defaults to 2419200.
- `mname` (String) Defaults to empty.
- `negative_ttl` (Number) Optional. ZoneAuthority negative caching (minimum) ttl in seconds.  Defaults to 900.
- `refresh` (Number) Optional. ZoneAuthority refresh.  Defaults to 10800.
- `retry` (Number) Optional. ZoneAuthority retry.  Defaults to 3600.
- `rname` (String) Optional. ZoneAuthority rname.  Defaults to empty.
- `use_default_mname` (Boolean) Optional. Use default value for master name server.  Defaults to true.

Read-Only:

- `protocol_mname` (String) Optional. ZoneAuthority master name server in punycode.  Defaults to empty.
- `protocol_rname` (String) Optional. A domain name which specifies the mailbox of the person responsible for this zone.  Defaults to empty.