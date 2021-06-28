package miro

import(
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("miro_user.user1", "email", "{USER EMAIL}"),
					resource.TestCheckResourceAttr("miro_user.user1", "role", "{USER ROLE}"),
					resource.TestCheckResourceAttr("miro_user.user1", "team_id", "{TEAM ID}"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
		resource "miro_user" "user1" {
			email	= "{USER EMAIL}"
			role	= "{USER ROLE}"
			team_id = "{TEAM ID}"
		}
	`)
}

func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"miro_user.user1", "email", "{USER EMAIL}"),	
					resource.TestCheckResourceAttr(
						"miro_user.user1", "role", "{USER ROLE}"),
					resource.TestCheckResourceAttr(
						"miro_user.user1", "team_id", "{TEAM ID}"),
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"miro_user.user1", "email", "{EMAIL}"),	
					resource.TestCheckResourceAttr(
						"miro_user.user1", "role", "{ROLE}"),
					resource.TestCheckResourceAttr(
						"miro_user.user1", "team_id", "{TEAM ID}"),
				),
			},
		},
	})
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
		resource "miro_user" "user1" {
			email	= "{EMAIL}"
			role	= "{ROLE}"
			team_id = "{TEAM ID}"
		}
	`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
		resource "miro_user" "user1" {
			email        = "{EMAIL}"
			role		 = "{ROLE}"
			team_id		 = "{TEAM ID}"
		}
	`)
}
