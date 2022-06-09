```go
package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApi.json
func ExampleAPIClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"tempgroup",
		armapimanagement.APICreateOrUpdateParameter{
			Properties: &armapimanagement.APICreateOrUpdateProperties{
				Description: to.Ptr("apidescription5200"),
				AuthenticationSettings: &armapimanagement.AuthenticationSettingsContract{
					OAuth2: &armapimanagement.OAuth2AuthenticationSettingsContract{
						AuthorizationServerID: to.Ptr("authorizationServerId2283"),
						Scope:                 to.Ptr("oauth2scope2580"),
					},
				},
				SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
					Header: to.Ptr("header4520"),
					Query:  to.Ptr("query3037"),
				},
				Path:        to.Ptr("newapiPath"),
				DisplayName: to.Ptr("apiname1463"),
				Protocols: []*armapimanagement.Protocol{
					to.Ptr(armapimanagement.ProtocolHTTPS),
					to.Ptr(armapimanagement.ProtocolHTTP)},
				ServiceURL: to.Ptr("http://newechoapi.cloudapp.net/api"),
			},
		},
		&armapimanagement.APIClientBeginCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv1.0.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.
