package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-07-01/SharedLimitCaps_List.json
func ExampleSharedLimitCapsClient_NewListBySubscriptionLocationResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharedLimitCapsClient().NewListBySubscriptionLocationResourcePager("eastus", nil)
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
		// page = armcomputelimit.SharedLimitCapsClientListBySubscriptionLocationResourceResponse{
		// 	SharedLimitCapListResult: armcomputelimit.SharedLimitCapListResult{
		// 		Value: []*armcomputelimit.SharedLimitCap{
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimitCaps/StandardDSv3Family"),
		// 				Name: to.Ptr("StandardDSv3Family"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimitCaps"),
		// 				Properties: &armcomputelimit.SharedLimitCapProperties{
		// 					DefaultMemberCap: to.Ptr[int32](500),
		// 					IsBoundedCap: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimitCaps/StandardFSv2Family"),
		// 				Name: to.Ptr("StandardFSv2Family"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimitCaps"),
		// 				Properties: &armcomputelimit.SharedLimitCapProperties{
		// 					DefaultMemberCap: to.Ptr[int32](600),
		// 					IsBoundedCap: to.Ptr(false),
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
