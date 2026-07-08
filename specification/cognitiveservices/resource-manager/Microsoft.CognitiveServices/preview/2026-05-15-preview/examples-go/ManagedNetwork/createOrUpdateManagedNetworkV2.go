package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/ManagedNetwork/createOrUpdateManagedNetworkV2.json
func ExampleManagedNetworkSettingsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedNetworkSettingsClient().BeginPut(ctx, "test-rg", "cognitive-account-name", "default", armcognitiveservices.ManagedNetworkSettingsPropertiesBasicResource{
		Properties: &armcognitiveservices.ManagedNetworkSettingsProperties{
			ManagedNetwork: &armcognitiveservices.ManagedNetworkSettingsEx{
				FirewallSKU:   to.Ptr(armcognitiveservices.FirewallSKUStandard),
				IsolationMode: to.Ptr(armcognitiveservices.IsolationModeAllowOnlyApprovedOutbound),
				OutboundRules: map[string]armcognitiveservices.OutboundRuleClassification{
					"rule_name_1": &armcognitiveservices.FqdnOutboundRule{
						Type:        to.Ptr(armcognitiveservices.RuleTypeFQDN),
						Category:    to.Ptr(armcognitiveservices.RuleCategoryUserDefined),
						Destination: to.Ptr("destination_endpoint"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.ManagedNetworkSettingsClientPutResponse{
	// 	ManagedNetworkSettingsPropertiesBasicResource: armcognitiveservices.ManagedNetworkSettingsPropertiesBasicResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedNetworks"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/cognitive-account-name/managedNetworks/default"),
	// 		Properties: &armcognitiveservices.ManagedNetworkSettingsProperties{
	// 			ManagedNetwork: &armcognitiveservices.ManagedNetworkSettingsEx{
	// 				FirewallPublicIPAddress: to.Ptr("22.22.131.49"),
	// 				FirewallSKU: to.Ptr(armcognitiveservices.FirewallSKUStandard),
	// 				IsolationMode: to.Ptr(armcognitiveservices.IsolationModeAllowOnlyApprovedOutbound),
	// 				NetworkID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				OutboundRules: map[string]armcognitiveservices.OutboundRuleClassification{
	// 					"rule_name_1": &armcognitiveservices.FqdnOutboundRule{
	// 						Type: to.Ptr(armcognitiveservices.RuleTypeFQDN),
	// 						Category: to.Ptr(armcognitiveservices.RuleCategoryUserDefined),
	// 						Destination: to.Ptr("destination_endpoint"),
	// 						Status: to.Ptr(armcognitiveservices.RuleStatusActive),
	// 					},
	// 				},
	// 				ProvisioningState: to.Ptr(armcognitiveservices.ManagedNetworkProvisioningStateSucceeded),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ManagedNetworkProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
