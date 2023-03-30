package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdateNestedResourceTypeThird.json
func ExampleSKUsClient_CreateOrUpdateNestedResourceTypeThird() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSKUsClient().CreateOrUpdateNestedResourceTypeThird(ctx, "Microsoft.Contoso", "testResourceType", "nestedResourceTypeFirst", "nestedResourceTypeSecond", "nestedResourceTypeThird", "testSku", armproviderhub.SKUResource{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUResource = armproviderhub.SKUResource{
	// 	Name: to.Ptr("Microsoft.Contoso/"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/"),
	// }
}
