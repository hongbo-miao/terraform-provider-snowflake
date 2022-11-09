---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "snowflake_schemas Data Source - terraform-provider-snowflake"
subcategory: ""
description: |-
  
---

# snowflake_schemas (Data Source)



## Example Usage

```terraform
data "snowflake_schemas" "current" {
    database = "MYDB"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database` (String) The database from which to return the schemas from.

### Read-Only

- `id` (String) The ID of this resource.
- `schemas` (List of Object) The schemas in the database (see [below for nested schema](#nestedatt--schemas))

<a id="nestedatt--schemas"></a>
### Nested Schema for `schemas`

Read-Only:

- `comment` (String)
- `database` (String)
- `name` (String)

