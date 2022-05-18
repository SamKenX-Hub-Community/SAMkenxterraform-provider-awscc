---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_voiceid_domain Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.
---

# awscc_voiceid_domain (Resource)

The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `server_side_encryption_configuration` (Attributes) (see [below for nested schema](#nestedatt--server_side_encryption_configuration))

### Optional

- `description` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `domain_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--server_side_encryption_configuration"></a>
### Nested Schema for `server_side_encryption_configuration`

Required:

- `kms_key_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_voiceid_domain.example <resource ID>
```