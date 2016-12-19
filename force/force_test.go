package force

import (
	"github.com/nimajalali/go-force/sobjects"
	"testing"
)

func TestCreateWithAccessToken(t *testing.T) {

	// Manually grab an OAuth token, so that we can pass it into CreateWithAccessToken
	oauth := &ForceOauth{
		clientId:      testClientId,
		clientSecret:  testClientSecret,
		userName:      testUserName,
		password:      testPassword,
		securityToken: testSecurityToken,
		environment:   testEnvironment,
	}

	forceApi := &ForceApi{
		apiResources:           make(map[string]string),
		apiSObjects:            make(map[string]*SObjectMetaData),
		apiSObjectDescriptions: make(map[string]*SObjectDescription),
		apiVersion:             version,
		OAuth:                  oauth,
	}

	err := forceApi.OAuth.Authenticate()
	if err != nil {
		t.Fatalf("Unable to authenticate: %#v", err)
	}
	if err := forceApi.OAuth.Validate(); err != nil {
		t.Fatalf("Oauth object is invlaid: %#v", err)
	}

	// We shouldn't hit any errors creating a new force instance and manually passing in these oauth details now.
	newForceApi, err := CreateWithAccessToken(testVersion, testClientId, forceApi.OAuth.clientSecret, forceApi.OAuth.AccessToken, forceApi.OAuth.RefreshToken, forceApi.OAuth.InstanceUrl)
	if err != nil {
		t.Fatalf("Unable to create new force api instance using pre-defined oauth details: %#v", err)
	}
	if err := newForceApi.OAuth.Validate(); err != nil {
		t.Fatalf("Oauth object is invlaid: %#v", err)
	}

	// We should be able to make a basic query now with the newly created object (i.e. the oauth details should be correctly usable).
	_, err = newForceApi.DescribeSObject(sobjects.Account{})
	if err != nil {
		t.Fatalf("Failed to retrieve description of sobject: %v", err)
	}
}
