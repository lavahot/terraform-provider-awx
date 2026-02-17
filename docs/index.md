
# AWX Provider

The AWX provider is used to interact with [Ansible AWX](https://github.com/ansible/awx) or [Red Hat Ansible Automation Platform](https://www.redhat.com/en/technologies/management/ansible) (formerly Ansible Tower). It allows you to manage AWX resources such as organizations, inventories, projects, job templates, credentials, teams, users, and more through Terraform.

This provider communicates with the AWX [REST API](https://docs.ansible.com/projects/awx/en/latest/rest_api/index.html) and is built on the [goawx](https://github.com/denouche/goawx) Go client library.

## Example Usage

Using username and password:
```hcl
provider "awx" {
  hostname = "http://localhost:8078"
  username = "test"
  password = "changeme"
}
```

Using token:
```hcl
provider "awx" {
  hostname = "http://localhost:8078"
  token    = "awxtoken"
}
```

> ⚠️ Be careful, if you set both token and username/password the token will have the precedence.

## Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The API endpoint for AWX. Defaults to `"http://localhost"`.
* `username` - (Optional) The username for API access. Defaults to `"admin"`.
* `password` - (Optional) The password for API access. Defaults to `"password"`.
* `token`    - (Optional) The AWX token for API access. Defaults to empty.
* `insecure` - (Optional) Whether to skip TLS certificate verification. Set to `true` if using self-signed certificates. Defaults to `false`.
