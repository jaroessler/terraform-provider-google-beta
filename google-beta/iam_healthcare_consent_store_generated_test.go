// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHealthcareConsentStoreIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStoreIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccHealthcareConsentStoreIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccHealthcareConsentStoreIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccHealthcareConsentStoreIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccHealthcareConsentStoreIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareConsentStoreIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccHealthcareConsentStoreIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccHealthcareConsentStoreIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  provider = google-beta

  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_healthcare_consent_store_iam_member" "foo" {
  provider = google-beta
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccHealthcareConsentStoreIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  provider = google-beta

  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_healthcare_consent_store_iam_policy" "foo" {
  provider = google-beta
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccHealthcareConsentStoreIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  provider = google-beta

  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_healthcare_consent_store_iam_policy" "foo" {
  provider = google-beta
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccHealthcareConsentStoreIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  provider = google-beta

  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_healthcare_consent_store_iam_binding" "foo" {
 
  provider = google-beta
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccHealthcareConsentStoreIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  location = "us-central1"
  name     = "tf-test-my-dataset%{random_suffix}"
}

resource "google_healthcare_consent_store" "my-consent" {
  provider = google-beta

  dataset = google_healthcare_dataset.dataset.id
  name    = "tf-test-my-consent-store%{random_suffix}"
}

resource "google_healthcare_consent_store_iam_binding" "foo" {
  provider = google-beta
  dataset = google_healthcare_consent_store.my-consent.dataset
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
