package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShields_ListBySubscription.json
func ExampleNeuroShieldsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNeuroShieldsClient().NewListBySubscriptionPager(nil)
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
		// page = armnetworksecurity.NeuroShieldsClientListBySubscriptionResponse{
		// 	NeuroShieldResourceListResult: armnetworksecurity.NeuroShieldResourceListResult{
		// 		Value: []*armnetworksecurity.NeuroShieldResource{
		// 			{
		// 				Name: to.Ptr("myNeuroShield1"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/microsoft.networksecurity/neuroShields/myNeuroShield1"),
		// 				Type: to.Ptr("microsoft.networksecurity/neuroShields"),
		// 				Location: to.Ptr("eastus2"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("test"),
		// 				},
		// 				Properties: &armnetworksecurity.NeuroShieldProperties{
		// 					ProvisioningState: to.Ptr(armnetworksecurity.ProvisioningStateNeuroShieldSucceeded),
		// 					ResourceGUID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					ResponseMode: to.Ptr(armnetworksecurity.NeuroShieldResponseModeManual),
		// 					SeverityThreshold: to.Ptr(armnetworksecurity.NeuroShieldSeverityThresholdMedium),
		// 					Description: to.Ptr("First NeuroShield resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myNeuroShield2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/microsoft.networksecurity/neuroShields/myNeuroShield2"),
		// 				Type: to.Ptr("microsoft.networksecurity/neuroShields"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("prod"),
		// 				},
		// 				Properties: &armnetworksecurity.NeuroShieldProperties{
		// 					ProvisioningState: to.Ptr(armnetworksecurity.ProvisioningStateNeuroShieldSucceeded),
		// 					ResourceGUID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 					ResponseMode: to.Ptr(armnetworksecurity.NeuroShieldResponseModeAuto),
		// 					SeverityThreshold: to.Ptr(armnetworksecurity.NeuroShieldSeverityThresholdHigh),
		// 					Description: to.Ptr("Second NeuroShield resource"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
