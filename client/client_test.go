package client

import(
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
)

func init(){
	token := "{MIRO_TOKEN}"
	os.Setenv("MIRO_TOKEN", token)
}

func TestClient_GetUser(t *testing.T) {
	testCases := []struct {
		testName     string
		userName     string
		team_id	     string
		seedData     map[string]getUserStruct
		expectErr    bool
		expectedResp getUserStruct
	}{
		{
			testName: "user exists",
			userName: "{USER's EMAIL}",
			team_id:  "{MIRO_TEAM_ID}",
			seedData: map[string]getUserStruct{
				"user1": {
					Type:		"user",
					ID:			"{USER ID}",
					Name:		"{USER NAME}",
					Industry:	"{INDUSTRY}",
					CreatedAt:	"{CREATED AT}",
					Company:	"{COMPANY}",
					Role:		"{ROLE}",
					TeamName:	"{TEAM NAME}",
					Email:		"{EMAIL}",
					State:		"{STATE}",
				},
			},
			expectErr: false,
			expectedResp: getUserStruct{
					Type:		"user",
					ID:			"{USER ID}",
					Name:		"{USER NAME}",
					Industry:	"{INDUSTRY}",
					CreatedAt:	"{CREATED AT}",
					Company:	"{COMPANY}",
					Role:		"{ROLE}",
					TeamName:	"{TEAM NAME}",
					Email:		"{EMAIL}",
					State:		"{STATE}",
			},
		},
			{
			testName:     "user does not exist",
			userName:     "{USER's EMAIL ID}",
			team_id:      "{TEAM ID}",
			seedData:     nil,
			expectErr:    true,
			expectedResp: getUserStruct{},
			},
		}
	

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client 		:= NewClient(os.Getenv("MIRO_TOKEN"))
			item, err 	:= client.GetUser(tc.userName, tc.team_id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}
func TestClient_CreateUser(t *testing.T) {
	testCases :=  []struct {
		testName  string
		newUser   string
		team_id   string
		seedData  map[string]string
		expectErr bool
	}{
		{
			testName: "success",
			newUser:  "{EMAIL}",
			team_id:  "{TEAM ID}",
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "user already exists",
			newUser:  "{EMAIL}",
			team_id:  "{TEAM ID}",
			seedData:  map[string]string {"user1": "{EMAIL}"},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client 	:= NewClient( os.Getenv("MIRO_TOKEN"))
			err 	:= client.CreateUser(tc.newUser, tc.team_id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName    string
		role		string
		team_id		string
		seedData    map[string]string
		expectErr   bool
		email		string
	}{
		{
			testName: "user exists",
			role:	  "{ROLE}",
			team_id:  "{MIRO TEAM ID}",
			email:	  "{EMAIL}",
			seedData: map[string]string{
				"user1": "{USER ID}",
			},
			expectErr: false,
		},
		{
			testName: "user does not exist",
			role:	  "admin",
			team_id:  "{MIRO TEAM ID}",
			email:	  "{EMAIL}",
			seedData:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("MIRO_TOKEN"))
			err    := client.UpdateUser(tc.email, tc.role, tc.team_id)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetUser(tc.email, tc.team_id)
			assert.NoError(t, err)
			assert.Equal(t, tc.role, user.Role)
		})
	}
}

func TestClient_DeleteUser(t *testing.T) {
	testCases := []struct {
		testName  string
		user	  string
		team_id   string
		seedData  map[string]string
		expectErr bool
	}{
		{
			testName: "user exists",
			user:     "{EMAIL ID}",
			team_id:  "{TEAM ID}",
			seedData: map[string]string{
				"user1": "{USER ID}",
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client 		:= NewClient(os.Getenv("MIRO_TOKEN"))
			err := client.DeleteUser(tc.user, tc.team_id)
			if err != nil {
				assert.NoError(t,err)
			}
		})
	}
}
