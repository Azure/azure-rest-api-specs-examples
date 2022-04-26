Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaps%2Farmmaps%2Fv0.4.0/sdk/resourcemanager/maps/armmaps/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/CreateAccountManagedIdentity.json
func ExampleAccountsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmaps.NewAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmaps.Account{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"test": to.Ptr("true"),
			},
			Identity: &armmaps.ManagedServiceIdentity{
				Type: to.Ptr(armmaps.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armmaps.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName": {},
				},
			},
			Kind: to.Ptr(armmaps.KindGen2),
			Properties: &armmaps.AccountProperties{
				DisableLocalAuth: to.Ptr(false),
				LinkedResources: []*armmaps.LinkedResource{
					{
						ID:         to.Ptr("<id>"),
						UniqueName: to.Ptr("<unique-name>"),
					},
					{
						ID:         to.Ptr("<id>"),
						UniqueName: to.Ptr("<unique-name>"),
					}},
			},
			SKU: &armmaps.SKU{
				Name: to.Ptr(armmaps.NameG2),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
