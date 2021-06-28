package miro

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.miro_user.user3", "email", "{USER EMAIL}"),
					resource.TestCheckResourceAttr("data.miro_user.user3", "team_id", "{TEAM ID}"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(`	  
	resource "miro_user" "user2" {
		email        = "{USER EMAIL}"
		role		 = "{ROLE}"
		team_id		 = "{TEAM ID}"
	  }
	data "miro_user" "user3" {
		email        = "{USER EMAIL}"
		team_id		 = "{TEAM ID}"
	}
	`)
}