Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv1.0.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateAuthorizationServer.json
func ExampleAuthorizationServerClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAuthorizationServerClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"newauthServer",
		armapimanagement.AuthorizationServerContract{
			Properties: &armapimanagement.AuthorizationServerContractProperties{
				Description: to.Ptr("test server"),
				AuthorizationMethods: []*armapimanagement.AuthorizationMethod{
					to.Ptr(armapimanagement.AuthorizationMethodGET)},
				BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethod{
					to.Ptr(armapimanagement.BearerTokenSendingMethodAuthorizationHeader)},
				DefaultScope:               to.Ptr("read write"),
				ResourceOwnerPassword:      to.Ptr("pwd"),
				ResourceOwnerUsername:      to.Ptr("un"),
				SupportState:               to.Ptr(true),
				TokenEndpoint:              to.Ptr("https://www.contoso.com/oauth2/token"),
				AuthorizationEndpoint:      to.Ptr("https://www.contoso.com/oauth2/auth"),
				ClientID:                   to.Ptr("1"),
				ClientRegistrationEndpoint: to.Ptr("https://www.contoso.com/apps"),
				ClientSecret:               to.Ptr("2"),
				DisplayName:                to.Ptr("test2"),
				GrantTypes: []*armapimanagement.GrantType{
					to.Ptr(armapimanagement.GrantTypeAuthorizationCode),
					to.Ptr(armapimanagement.GrantTypeImplicit)},
			},
		},
		&armapimanagement.AuthorizationServerClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
