Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv0.5.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAuthorizationServerClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<authsid>",
		armapimanagement.AuthorizationServerContract{
			Properties: &armapimanagement.AuthorizationServerContractProperties{
				Description: to.Ptr("<description>"),
				AuthorizationMethods: []*armapimanagement.AuthorizationMethod{
					to.Ptr(armapimanagement.AuthorizationMethodGET)},
				BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethod{
					to.Ptr(armapimanagement.BearerTokenSendingMethodAuthorizationHeader)},
				DefaultScope:               to.Ptr("<default-scope>"),
				ResourceOwnerPassword:      to.Ptr("<resource-owner-password>"),
				ResourceOwnerUsername:      to.Ptr("<resource-owner-username>"),
				SupportState:               to.Ptr(true),
				TokenEndpoint:              to.Ptr("<token-endpoint>"),
				AuthorizationEndpoint:      to.Ptr("<authorization-endpoint>"),
				ClientID:                   to.Ptr("<client-id>"),
				ClientRegistrationEndpoint: to.Ptr("<client-registration-endpoint>"),
				ClientSecret:               to.Ptr("<client-secret>"),
				DisplayName:                to.Ptr("<display-name>"),
				GrantTypes: []*armapimanagement.GrantType{
					to.Ptr(armapimanagement.GrantTypeAuthorizationCode),
					to.Ptr(armapimanagement.GrantTypeImplicit)},
			},
		},
		&armapimanagement.AuthorizationServerClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
