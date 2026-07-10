package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Skus_ListBySubscription.json
func ExampleSKUsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListBySubscriptionPager(nil)
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
		// page = armdevcenter.SKUsClientListBySubscriptionResponse{
		// 	SKUListResult: armdevcenter.SKUListResult{
		// 		Value: []*armdevcenter.SKUInfo{
		// 			{
		// 				Name: to.Ptr("Large"),
		// 				Locations: []*string{
		// 					to.Ptr("CentralUS"),
		// 				},
		// 				ResourceType: to.Ptr("projects/pools"),
		// 				Tier: to.Ptr(armdevcenter.SKUTierPremium),
		// 			},
		// 			{
		// 				Name: to.Ptr("Medium"),
		// 				Locations: []*string{
		// 					to.Ptr("CentralUS"),
		// 				},
		// 				ResourceType: to.Ptr("projects/pools"),
		// 				Tier: to.Ptr(armdevcenter.SKUTierStandard),
		// 			},
		// 		},
		// 	},
		// }
	}
}
