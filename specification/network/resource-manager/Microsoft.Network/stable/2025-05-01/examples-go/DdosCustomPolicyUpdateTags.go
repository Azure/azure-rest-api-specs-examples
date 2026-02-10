package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DdosCustomPolicyUpdateTags.json
func ExampleDdosCustomPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDdosCustomPoliciesClient().UpdateTags(ctx, "rg1", "test-ddos-custom-policy", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DdosCustomPolicy = armnetwork.DdosCustomPolicy{
	// 	Name: to.Ptr("test-ddos-custom-policy"),
	// 	Type: to.Ptr("Microsoft.Network/ddosCustomPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.DdosCustomPolicyPropertiesFormat{
	// 		DetectionRules: []*armnetwork.DdosDetectionRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosCustomPolicies/test-ddos-custom-policy/ddosDetectionRules/detectionRuleTcp"),
	// 				Name: to.Ptr("detectionRuleTcp"),
	// 				Type: to.Ptr("Microsoft.Network/ddosCustomPolicies/ddosDetectionRules"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.DdosDetectionRulePropertiesFormat{
	// 					DetectionMode: to.Ptr(armnetwork.DdosDetectionModeTrafficThreshold),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					TrafficDetectionRule: &armnetwork.TrafficDetectionRule{
	// 						PacketsPerSecond: to.Ptr[int32](1000000),
	// 						TrafficType: to.Ptr(armnetwork.DdosTrafficTypeTCP),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// }
}
