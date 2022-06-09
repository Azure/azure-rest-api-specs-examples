```go
package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/UpdateAccountManagedIdentity.json
func ExampleAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmaps.NewAccountsClient("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"myResourceGroup",
		"myMapsAccount",
		armmaps.AccountUpdateParameters{
			Identity: &armmaps.ManagedServiceIdentity{
				Type: to.Ptr(armmaps.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armmaps.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName": {},
				},
			},
			Kind: to.Ptr(armmaps.KindGen2),
			Properties: &armmaps.AccountProperties{
				LinkedResources: []*armmaps.LinkedResource{
					{
						ID:         to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/accounts/{storageName}"),
						UniqueName: to.Ptr("myBatchStorageAccount"),
					}},
			},
			SKU: &armmaps.SKU{
				Name: to.Ptr(armmaps.NameG2),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaps%2Farmmaps%2Fv0.5.0/sdk/resourcemanager/maps/armmaps/README.md) on how to add the SDK to your project and authenticate.
