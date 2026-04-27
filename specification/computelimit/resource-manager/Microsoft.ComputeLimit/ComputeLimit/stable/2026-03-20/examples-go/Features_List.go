package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-03-20/Features_List.json
func ExampleFeaturesClient_NewListBySubscriptionLocationResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("74219ad7-63fc-442f-8037-4b43c627c07d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFeaturesClient().NewListBySubscriptionLocationResourcePager("eastus", nil)
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
		// page = armcomputelimit.FeaturesClientListBySubscriptionLocationResourceResponse{
		// 	FeatureListResult: armcomputelimit.FeatureListResult{
		// 		Value: []*armcomputelimit.Feature{
		// 			{
		// 				ID: to.Ptr("/subscriptions/74219ad7-63fc-442f-8037-4b43c627c07d/providers/Microsoft.ComputeLimit/locations/eastus/features/VmCategoryQuota"),
		// 				Name: to.Ptr("VmCategoryQuota"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/features"),
		// 				Properties: &armcomputelimit.FeatureProperties{
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 					State: to.Ptr(armcomputelimit.FeatureStateEnabled),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
