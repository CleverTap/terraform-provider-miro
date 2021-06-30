This Terraform provider enables to create, read, update, delete, and import operations for Miro users.
 
 
## Requirements
 
 
* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* Application: [Miro](https://miro.com/login/) (API is supported in all plans.)
* [Miro API Documentation](https://developers.miro.com/reference#introduction)
 
## Application Account
 
### Setup<a id="setup"></a>
1. Create a Miro account. (https://miro.com/signup/)
2. And Sign in to the Miro.
3. Follow the link create app. (https://developers.miro.com/docs/getting-started#section-step-1-get-developer-team)
4. Agree to the T&C and click on Create APP.
5. Set App name and Copy access token. (Or you can get an access token later.)
 
 
### API Authentication
 *Miro uses OAuth 2.0 for authentication which provides Access Token to authenticate to the API.*
1. Create a Dev Team and an Application ([follow this link](https://developers.miro.com/docs/getting-started))<br>
2. We need client_id, redirected_uri for token generation, client_id can found in your app under the profile section created in the previous step and localhost URI can be used as redirected_url.
3. Put the required variable in the following link and make a request and it will redirect to provided redirected_uri with code.
Link: https://miro.com/oauth/authorize?response_type=code&client_id={client_ID}&redirect_uri={redirected_uri}/ <br>
3. It will redirect to {redirected_uri}/code=AUTH_CODE
4. Take auth code form here and make a POST request on Postman using following url
https://api.miro.com/v1/oauth/token?grant_type=authorization_code&code={AUTH_CODE}&redirect_uri=REDIRECT_URI&client_id=CLIENT_ID&client_secret=CLIENT_SECRET
Client Id and Secret Id can found under your app section. <br>
5. It will return a JSON response containing non-expiring access_token.
 
 
## Building The Provider
1. Clone the repository, add all the dependencies and create a vendor directory that contains all dependencies. For this, run the following commands: <br>
```
cd terraform-provider-miro
go mod init terraform-provider-miro
go mod tidy
go mod vendor
```
 
## Managing terraform plugins
*For Windows:*
1. Run the following command to create a vendor sub-directory (`%APPDATA%/terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${OS_ARCH}`) which will consist of all terraform plugins. <br> 
Command: 
```bash
mkdir -p %APPDATA%/Roaming/terraform.d/plugins/hashicorp.com/edu/miro/0.3/windows_amd64
```
2. Run `go build -o terraform-provider-miro.exe` to generate the binary in the present working directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-miro.exe %APPDATA%\Roaming\terraform.d\plugins\hashicorp.com\edu\miro\0.3\windows_amd64
 ``` 
<p align="center">[OR]</p>
 
3. Manually move the file from current directory to destination directory (`%APPDATA%\Roaming\terraform.d\plugins\hashicorp.com\edu\miro\0.3\windows_amd64`).<br>
 
 
## Working with terraform
 
### Application Credential Integration in terraform
1. Add `terraform` block and `provider` block as shown in [example usage](#example-usage).
2. Get the credentials: access token and team_id or team_name.
3. Assign the above credentials to the respective field in the `provider` block.
 
### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the changes.
 
### Create User
1. Add the `email` and `role` (role is Optional) in the respective field in the `resource` block as shown in [example usage](#example-usage).
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully invited and an invitation mail has been sent to the user.
5. Set the account using invitation.
 
### Update the user
Update the data of the user in the respective field in the `resource` block as shown in [example usage](#example-usage).
(in this case, the only role can be updated.)
And apply using `terraform apply`
 
### Read the User Data
Add `data` and `output` blocks as shown in the [example usage](#example-usage) and run `terraform apply` to read the user data.
 
### Delete the user
Delete the `resource` block of the user and run `terraform apply`.
 
### Import a User Data
1. Write manually a resource configuration block for the User as shown in [example usage](#example-usage).
2. Run the command `terraform import miro_user.user1 [EMAIL_ID]:[TEAM_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in the resource block.
 
## Example Usage<a id="example-usage"></a>
 
```terraform
terraform {
  required_providers {
    miro = {
      version = "0.3"
      source  = "hashicorp.com/edu/miro"
    }
  }
}
 
provider "miro" {
  miro_token = "_REPLACE_MIRO_TOKEN"
}
 
resource "miro_user" "user1" {
   email      = "demo@domain.com"
   team_id = "_REPLACE_MIRO_TEAM_ID"
   role       = "admin"
}
 
output "user1" {
  value = miro_user.user1
}
 
data "miro_user" "user2" {
  email = "demo@domain.com"
  team_id = "_REPLACE_MIRO_TEAM_ID"
}
 
output "user2" {
  value = data.miro_user.user2
}
```
 
## Argument Reference
* `miro_token` (Required, String) - Miro provides an access token for authentication which we have already seen how to get access_token.
* `miro_team_id` (Required, String) - It's an ID associated with for team, It can get from the profile page.
* `email` (Required, String) - Email is the user's mail id who is a part of the team or will be a part of the team.
* `role` (Optional, string) - Role is the user's role in the team it can be ```admin``` or ```member```.
 
## Exceptions
 
* New user's role will be a member by default but it can be changed later.
* The Only Role can be updated for any team member.
* The role of the last administrator for your team cannot be changed.
* ```non_team``` users can't be managed. <br>
REFERENCE: https://community.miro.com/developer-platform-and-apis-57/how-to-get-email-address-of-non-team-role-member-via-the-api-823 

* If the removed user owns any boards or projects, they also will be removed.
In case you want to save them, you need to reassign ownership first.
* Last team member(which would be admin) can't be removed.
