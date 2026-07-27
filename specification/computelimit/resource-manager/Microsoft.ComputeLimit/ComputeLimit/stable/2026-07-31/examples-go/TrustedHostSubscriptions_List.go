package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-07-31/TrustedHostSubscriptions_List.json
func ExampleTrustedHostSubscriptionsClient_NewListBySubscriptionLocationResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTrustedHostSubscriptionsClient().NewListBySubscriptionLocationResourcePager("eastus", nil)
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
		// page = armcomputelimit.TrustedHostSubscriptionsClientListBySubscriptionLocationResourceResponse{
		// 	TrustedHostSubscriptionListResult: armcomputelimit.TrustedHostSubscriptionListResult{
		// 		Value: []*armcomputelimit.TrustedHostSubscription{
		// 			{
		// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.ComputeLimit/locations/eastus/trustedHostSubscriptions/22222222-2222-2222-2222-222222222222"),
		// 				Name: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/trustedHostSubscriptions"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/providers/Microsoft.ComputeLimit/locations/eastus/trustedHostSubscriptions/33333333-3333-3333-3333-333333333333"),
		// 				Name: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/trustedHostSubscriptions"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
