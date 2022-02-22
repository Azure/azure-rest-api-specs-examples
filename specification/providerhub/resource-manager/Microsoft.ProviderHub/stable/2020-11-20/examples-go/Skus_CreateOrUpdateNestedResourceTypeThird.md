Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fproviderhub%2Farmproviderhub%2Fv0.2.1/sdk/resourcemanager/providerhub/armproviderhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdateNestedResourceTypeThird.json
func ExampleSKUsClient_CreateOrUpdateNestedResourceTypeThird() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateNestedResourceTypeThird(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<nested-resource-type-third>",
		"<sku>",
		armproviderhub.SKUResource{
			Properties: &armproviderhub.SKUResourceProperties{
				SKUSettings: []*armproviderhub.SKUSetting{
					{
						Name: to.StringPtr("<name>"),
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					},
					{
						Name: to.StringPtr("<name>"),
						Costs: []*armproviderhub.SKUCost{
							{
								MeterID: to.StringPtr("<meter-id>"),
							}},
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientCreateOrUpdateNestedResourceTypeThirdResult)
}
```
