---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
layout: "google"
page_title: "Google: google_compute_organization_security_policy"
sidebar_current: "docs-google-compute-organization-security-policy"
description: |-
  Organization security policies are used to control incoming/outgoing traffic.
---

# google\_compute\_organization\_security\_policy

Organization security policies are used to control incoming/outgoing traffic.

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about OrganizationSecurityPolicy, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies)
* How-to Guides
    * [Creating a firewall policy](https://cloud.google.com/vpc/docs/using-firewall-policies#create-policy)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=organization_security_policy_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Organization Security Policy Basic


```hcl
resource "google_compute_organization_security_policy" "policy" {
  provider = google-beta

  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/123456789"
}
```

## Argument Reference

The following arguments are supported:


* `display_name` -
  (Required)
  A textual name of the security policy.

* `parent` -
  (Required)
  The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
  Format: organizations/{organization_id} or folders/{folder_id}


- - -


* `description` -
  (Optional)
  A textual description for the organization security policy.

* `type` -
  (Optional)
  The type indicates the intended use of the security policy.
  For organization security policies, the only supported type
  is "FIREWALL".
  Default value is `FIREWALL`.
  Possible values are `FIREWALL`.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `locations/global/securityPolicies/{{policy_id}}`

* `fingerprint` -
  Fingerprint of this resource. This field is used internally during
  updates of this resource.

* `policy_id` -
  The unique identifier for the resource. This identifier is defined by the server.


## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 4 minutes.
- `update` - Default is 4 minutes.
- `delete` - Default is 4 minutes.

## Import


OrganizationSecurityPolicy can be imported using any of these accepted formats:

```
$ terraform import google_compute_organization_security_policy.default locations/global/securityPolicies/{{policy_id}}
$ terraform import google_compute_organization_security_policy.default {{policy_id}}
```
