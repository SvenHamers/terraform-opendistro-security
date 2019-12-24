# Opendistro Security Terraform Provider

Work in progress

### working features:
- users
- roles
- rolemappings


### Example

```
provider "opendistro" {
  user = "admin"
  password = "admin"
  base_url = "https://opendistro.example.com:9200"
}

resource "opendistro_user" "example" {
    user_name = "example"
    password = "test123!1234"
    backend_roles = ["admin","example"]
}


resource "opendistro_role" "example" {

    role_name = "example"

    index_permissions{
        index_patterns=["test*"]
        allowed_actions = ["read"]

        ###
        dls = ""
        fls = [""]
        masked_fields = [""]
        ###
    }


    ###
    tenant_permissions {
        tenant_patterns = [""]
        allowed_actions = [""]
    }
    ###


    cluster_permissions = [""]



}

resource "opendistro_role_mapping" "example" {
  rolemapping_name = "exmaple"
  backend_roles = ["example"]
  hosts = ["*"]
  users = ["example"]
}

```