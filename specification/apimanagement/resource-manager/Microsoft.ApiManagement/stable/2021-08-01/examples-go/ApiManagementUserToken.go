package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserToken.json
func ExampleUserClient_GetSharedAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewUserClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetSharedAccessToken(ctx,
		"rg1",
		"apimService1",
		"userId1718",
		armapimanagement.UserTokenParameters{
			Properties: &armapimanagement.UserTokenParameterProperties{
				Expiry:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T00:44:24.2845269Z"); return t }()),
				KeyType: to.Ptr(armapimanagement.KeyTypePrimary),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
