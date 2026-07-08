package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/ManagedNetwork/listManagedNetworkV2.json
func ExampleManagedNetworkSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedNetworkSettingsClient().NewListPager("test-rg", "cognitive-account-name", nil)
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
		// page = armcognitiveservices.ManagedNetworkSettingsClientListResponse{
		// 	ManagedNetworkListResult: armcognitiveservices.ManagedNetworkListResult{
		// 		Value: []*armcognitiveservices.ManagedNetworkSettingsPropertiesBasicResource{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/managedNetworks"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/cognitive-account-name/managedNetworks/default"),
		// 				Properties: &armcognitiveservices.ManagedNetworkSettingsProperties{
		// 					ManagedNetwork: &armcognitiveservices.ManagedNetworkSettingsEx{
		// 						FirewallPublicIPAddress: to.Ptr("22.22.131.49"),
		// 						FirewallSKU: to.Ptr(armcognitiveservices.FirewallSKUStandard),
		// 						IsolationMode: to.Ptr(armcognitiveservices.IsolationModeAllowOnlyApprovedOutbound),
		// 						NetworkID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						OutboundRules: map[string]armcognitiveservices.OutboundRuleClassification{
		// 							"rule_name_1": &armcognitiveservices.FqdnOutboundRule{
		// 								Type: to.Ptr(armcognitiveservices.RuleTypeFQDN),
		// 								Category: to.Ptr(armcognitiveservices.RuleCategoryUserDefined),
		// 								Destination: to.Ptr("destination_endpoint"),
		// 								Status: to.Ptr(armcognitiveservices.RuleStatusActive),
		// 							},
		// 						},
		// 						ProvisioningState: to.Ptr(armcognitiveservices.ManagedNetworkProvisioningStateSucceeded),
		// 					},
		// 					ProvisioningState: to.Ptr(armcognitiveservices.ManagedNetworkProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
