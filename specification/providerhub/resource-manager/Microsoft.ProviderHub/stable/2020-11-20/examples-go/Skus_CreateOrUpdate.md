```go
package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdate.json
func ExampleSKUsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armproviderhub.NewSKUsClient("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"Microsoft.Contoso",
		"testResourceType",
		"testSku",
		armproviderhub.SKUResource{
			Properties: &armproviderhub.SKUResourceProperties{
				SKUSettings: []*armproviderhub.SKUSetting{
					{
						Name: to.Ptr("freeSku"),
						Kind: to.Ptr("Standard"),
						Tier: to.Ptr("Tier1"),
					},
					{
						Name: to.Ptr("premiumSku"),
						Costs: []*armproviderhub.SKUCost{
							{
								MeterID: to.Ptr("xxx"),
							}},
						Kind: to.Ptr("Premium"),
						Tier: to.Ptr("Tier2"),
					}},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv1.0.0/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.
