package miro

import(
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
	"log"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	var token string   = "{MIRO TOKEN}"
	os.Setenv("MIRO_TOKEN", token)
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"miro": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ",err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T)  {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MIRO_TOKEN"); v == "" {
		t.Fatal("Miro APi TOKEN must be set for acceptance tests.")
	}
}