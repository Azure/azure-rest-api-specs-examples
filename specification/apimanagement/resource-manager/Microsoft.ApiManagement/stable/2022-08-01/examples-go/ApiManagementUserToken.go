package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUserToken.json
func ExampleUserClient_GetSharedAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().GetSharedAccessToken(ctx, "rg1", "apimService1", "userId1718", armapimanagement.UserTokenParameters{
		Properties: &armapimanagement.UserTokenParameterProperties{
			Expiry:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T00:44:24.284Z"); return t }()),
			KeyType: to.Ptr(armapimanagement.KeyTypePrimary),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserTokenResult = armapimanagement.UserTokenResult{
	// 	Value: to.Ptr("userId1718&201904210044&9A1GR1f5WIhFvFmzQG+xxxxxxxxxxx/kBeu87DWad3tkasUXuvPL+MgzlwUHyg=="),
	// }
}
