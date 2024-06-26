package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrations.json
func ExampleSKUsClient_NewListByResourceTypeRegistrationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListByResourceTypeRegistrationsPager("Microsoft.Contoso", "testResourceType", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SKUResourceArrayResponseWithContinuation = armproviderhub.SKUResourceArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.SKUResource{
		// 		{
		// 			Name: to.Ptr("testSku"),
		// 			Properties: &armproviderhub.SKUResourceProperties{
		// 				SKUSettings: []*armproviderhub.SKUSetting{
		// 					{
		// 						Name: to.Ptr("freeSku"),
		// 						Kind: to.Ptr("Standard"),
		// 						Tier: to.Ptr("Tier1"),
		// 					},
		// 					{
		// 						Name: to.Ptr("premiumSku"),
		// 						Costs: []*armproviderhub.SKUCost{
		// 							{
		// 								MeterID: to.Ptr("xxx"),
		// 						}},
		// 						Kind: to.Ptr("Premium"),
		// 						Tier: to.Ptr("Tier2"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
