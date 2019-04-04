package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsDescribeAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsDescribeAccountRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"joined_method": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"joined_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsDescribeAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient).organizationsconn
	account_id := d.Get("account_id").(string)
	req := &organizations.DescribeAccountInput{
		AccountId: aws.String(account_id),
	}

	log.Printf("[DEBUG] Getting account description: %s", req)
	res, err := client.DescribeAccount(req)
	if err != nil {
		return fmt.Errorf("Error describing account: %v", err)
	}

	log.Printf("[DEBUG] Received account description: %s", res)

	account := res.Account
	d.SetId(account_id)
	d.Set("arn", account.Arn)
	d.Set("email", account.Email)
	d.Set("joined_method", account.JoinedMethod)
	d.Set("joined_timestamp", aws.TimeValue(account.JoinedTimestamp).Unix())
	d.Set("name", account.Name)
	d.Set("status", account.Status)

	return nil
}
