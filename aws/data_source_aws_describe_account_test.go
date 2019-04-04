package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func testAccAwsDescribeAccount_basic(t *testing.T) {
	var organization organizations.Organization
	orgsEmailDomain, ok := os.LookupEnv("TEST_AWS_ORGANIZATION_ACCOUNT_EMAIL_DOMAIN")

	if !ok {
		t.Skip("'TEST_AWS_ORGANIZATION_ACCOUNT_EMAIL_DOMAIN' not set, skipping test.")
	}

	rInt := acctest.RandInt()
	name := fmt.Sprintf("tf_acctest_%d", rInt)
	email := fmt.Sprintf("tf-acctest+%d@%s", rInt, orgsEmailDomain)

	resourceName := "aws_organizations_account.test"
	datasourceName := "data.aws_describe_account.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccOrganizationsAccountPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsOrganizationsOrganizationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsDescribeAccountConfig(name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "email", resourceName, "email"),
					resource.TestCheckResourceAttrSet(datasourceName, "joined_method"),
					resource.TestCheckResourceAttrSet(datasourceName, "joined_timestamp"),
					resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrSet(datasourceName, "status"),
				),
			},
		},
	})
}

func testAccAwsDescribeAccountConfig(name, email string) string {
	return fmt.Sprintf(`
resource "aws_organizations_account" "test" {
  name  = "%s"
  email = "%s"
}

data "aws_describe_account" "test" {
  account_id = "${aws_organizations_account.id}"
}
`, name, email)
}

const testAccCheckAwsDescribeAccountConfig = `
`
