package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getMachineLearningServicesUsages.json
func ExampleUsagesClient_NewListPager_quotasListUsagesMachineLearningServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus", nil)
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
		// page.UsagesLimits = armquota.UsagesLimits{
		// 	Value: []*armquota.CurrentUsagesBase{
		// 		{
		// 			Name: to.Ptr("standardDv2Family"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus/providers/Microsoft.Quota/usages/standardDv2Family"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Standard Dv2 Family vCPUs"),
		// 					Value: to.Ptr("standardDv2Family"),
		// 				},
		// 				ResourceType: to.Ptr("dedicated"),
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("totalLowPriorityCores"),
		// 			Type: to.Ptr("Microsoft.Quota/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus/providers/Microsoft.Quota/usages/totalLowPriorityCores"),
		// 			Properties: &armquota.UsagesProperties{
		// 				Name: &armquota.ResourceName{
		// 					LocalizedValue: to.Ptr("Total Regional Low-priority vCPUs"),
		// 					Value: to.Ptr("totalLowPriorityCores"),
		// 				},
		// 				ResourceType: to.Ptr("lowPriority"),
		// 				Unit: to.Ptr("Count"),
		// 				Usages: &armquota.UsagesObject{
		// 					Value: to.Ptr[int32](10),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
