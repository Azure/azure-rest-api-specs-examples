package armdevopsinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/SubscriptionUsages_ListByLocation.json
func ExampleSubscriptionUsagesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionUsagesClient().NewListByLocationPager("eastus", nil)
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
		// page.QuotaListResult = armdevopsinfrastructure.QuotaListResult{
		// 	Value: []*armdevopsinfrastructure.Quota{
		// 		{
		// 			Name: to.Ptr("standardDADSv5Family"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.DevOpsInfrastructure/Usages/standardDADSv5Family"),
		// 			Properties: &armdevopsinfrastructure.QuotaProperties{
		// 				Name: &armdevopsinfrastructure.QuotaName{
		// 					LocalizedValue: to.Ptr("Standard DADSv5 Family vCPUs (PME VMSS)"),
		// 					Value: to.Ptr("standardDADSv5Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int64](0),
		// 				Limit: to.Ptr[int64](212),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("standardDPLDSv5Family"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.DevOpsInfrastructure/Usages/standardDPLDSv5Family"),
		// 			Properties: &armdevopsinfrastructure.QuotaProperties{
		// 				Name: &armdevopsinfrastructure.QuotaName{
		// 					LocalizedValue: to.Ptr("Standard DPLDSv5 Family vCPUs (PME VMSS)"),
		// 					Value: to.Ptr("standardDPLDSv5Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int64](0),
		// 				Limit: to.Ptr[int64](100),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 	}},
		// }
	}
}
