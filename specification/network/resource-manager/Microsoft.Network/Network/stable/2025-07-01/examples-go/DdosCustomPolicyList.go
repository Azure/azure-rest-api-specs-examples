package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/DdosCustomPolicyList.json
func ExampleDdosCustomPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDdosCustomPoliciesClient().NewListPager("rg1", nil)
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
		// page = armnetwork.DdosCustomPoliciesClientListResponse{
		// 	DdosCustomPolicyListResult: armnetwork.DdosCustomPolicyListResult{
		// 		Value: []*armnetwork.DdosCustomPolicy{
		// 			{
		// 				Name: to.Ptr("test-ddos-custom-policy"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy"),
		// 				Type: to.Ptr("Microsoft.Network/ddosCustomPolicies"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.DdosCustomPolicyPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					DetectionRules: []*armnetwork.DdosDetectionRule{
		// 						{
		// 							Name: to.Ptr("detectionRuleTcp"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy/ddosDetectionRules/detectionRuleTcp"),
		// 							Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 							Type: to.Ptr("Microsoft.Network/ddosCustomPolicies/ddosDetectionRules"),
		// 							Properties: &armnetwork.DdosDetectionRulePropertiesFormat{
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								DetectionMode: to.Ptr(armnetwork.DdosDetectionModeTrafficThreshold),
		// 								TrafficDetectionRule: &armnetwork.TrafficDetectionRule{
		// 									TrafficType: to.Ptr(armnetwork.DdosTrafficTypeTCP),
		// 									PacketsPerSecond: to.Ptr[int32](1000000),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("test-ddos-custom-policy-2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy-2"),
		// 				Type: to.Ptr("Microsoft.Network/ddosCustomPolicies"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetwork.DdosCustomPolicyPropertiesFormat{
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
