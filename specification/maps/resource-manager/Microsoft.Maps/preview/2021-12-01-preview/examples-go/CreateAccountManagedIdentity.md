Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaps%2Farmmaps%2Fv0.2.1/sdk/resourcemanager/maps/armmaps/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/CreateAccountManagedIdentity.json
func ExampleAccountsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaps.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmaps.Account{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"test": to.StringPtr("true"),
			},
			Identity: &armmaps.ManagedServiceIdentity{
				Type: armmaps.ResourceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armmaps.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName": {},
				},
			},
			Kind: armmaps.Kind("Gen2").ToPtr(),
			Properties: &armmaps.AccountProperties{
				DisableLocalAuth: to.BoolPtr(false),
				LinkedResources: []*armmaps.LinkedResource{
					{
						ID:         to.StringPtr("<id>"),
						UniqueName: to.StringPtr("<unique-name>"),
					},
					{
						ID:         to.StringPtr("<id>"),
						UniqueName: to.StringPtr("<unique-name>"),
					}},
			},
			SKU: &armmaps.SKU{
				Name: armmaps.Name("G2").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientCreateOrUpdateResult)
}
```
