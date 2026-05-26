package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v4"
)

// Generated from example definition: 2020-10-25/getMachineLearningServicesUsages.json
func ExampleQuotaClient_NewListPager_quotasListUsagesMachineLearningServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotaClient().NewListPager("00000000-0000-0000-0000-000000000000", "Microsoft.MachineLearningServices", "eastus", nil)
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
		// page = armreservations.QuotaClientListResponse{
		// 	QuotaLimits: armreservations.QuotaLimits{
		// 		Value: []*armreservations.CurrentQuotaLimitBase{
		// 			{
		// 				Name: to.Ptr("standardDv2Family"),
		// 				Type: to.Ptr("Microsoft.Capacity/ServiceLimits"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Capacity/resourceProviders/Microsoft.MachineLearningServices/locations/eastus/serviceLimits/standardDv2Family"),
		// 				Properties: &armreservations.QuotaProperties{
		// 					Name: &armreservations.ResourceName{
		// 						LocalizedValue: to.Ptr("Standard Dv2 Family vCPUs"),
		// 						Value: to.Ptr("standardDv2Family"),
		// 					},
		// 					CurrentValue: to.Ptr[int32](15),
		// 					Limit: to.Ptr[int32](20),
		// 					ResourceType: to.Ptr(armreservations.ResourceTypeDedicated),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("totalLowPriorityCores"),
		// 				Type: to.Ptr("Microsoft.Capacity/ServiceLimits"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Capacity/resourceProviders/Microsoft.MachineLearningServices/locations/eastus/serviceLimits/totalLowPriorityCores"),
		// 				Properties: &armreservations.QuotaProperties{
		// 					Name: &armreservations.ResourceName{
		// 						LocalizedValue: to.Ptr("Total Regional Low-priority vCPUs"),
		// 						Value: to.Ptr("totalLowPriorityCores"),
		// 					},
		// 					CurrentValue: to.Ptr[int32](49),
		// 					Limit: to.Ptr[int32](600),
		// 					ResourceType: to.Ptr(armreservations.ResourceTypeLowPriority),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
