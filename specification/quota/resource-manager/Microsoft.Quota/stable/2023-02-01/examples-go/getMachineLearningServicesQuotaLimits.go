package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getMachineLearningServicesQuotaLimits.json
func ExampleClient_NewListPager_quotasListQuotaLimitsMachineLearningServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus", nil)
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
		// page.Limits = armquota.Limits{
		// 	Value: []*armquota.CurrentQuotaLimitBase{
		// 		{
		// 			Name: to.Ptr("standardDv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus/providers/Microsoft.Quota/Quotas/standardDv2Family"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard Dv2 Family vCPUs"),
		// 					Value: to.Ptr("standardDv2Family"),
		// 				},
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				ResourceType: to.Ptr("dedicated"),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("totalLowPriorityCores"),
		// 			Type: to.Ptr("Microsoft.Quota/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus/providers/Microsoft.Quota/Quotas/totalLowPriorityCores"),
		// 			Properties: &armquota.Properties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Total Regional Low-priority vCPUs"),
		// 					Value: to.Ptr("totalLowPriorityCores"),
		// 				},
		// 				Limit: &armquota.LimitObject{
		// 					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 				ResourceType: to.Ptr("lowPriority"),
		// 				Unit: to.Ptr("Count"),
		// 			},
		// 	}},
		// }
	}
}
