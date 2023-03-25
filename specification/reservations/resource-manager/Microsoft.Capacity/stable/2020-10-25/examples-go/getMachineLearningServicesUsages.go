package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getMachineLearningServicesUsages.json
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
		// page.QuotaLimits = armreservations.QuotaLimits{
		// 	Value: []*armreservations.CurrentQuotaLimitBase{
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard Dv2 Family vCPUs"),
		// 					Value: to.Ptr("standardDv2Family"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](15),
		// 				Limit: to.Ptr[int32](20),
		// 				ResourceType: to.Ptr(armreservations.ResourceTypeDedicated),
		// 			},
		// 		},
		// 		{
		// 			Properties: &armreservations.QuotaProperties{
		// 				Name: &armreservations.ResourceName{
		// 					LocalizedValue: to.Ptr("Total Regional Low-priority vCPUs"),
		// 					Value: to.Ptr("totalLowPriorityCores"),
		// 				},
		// 				CurrentValue: to.Ptr[int32](49),
		// 				Limit: to.Ptr[int32](600),
		// 				ResourceType: to.Ptr(armreservations.ResourceTypeLowPriority),
		// 			},
		// 	}},
		// }
	}
}
