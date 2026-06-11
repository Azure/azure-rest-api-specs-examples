package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/ManagedNetwork/listManagedNetworkV2.json
func ExampleManagedNetworkSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedNetworkSettingsClient().NewListPager("test-rg", "aml-workspace-name", nil)
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
		// page = armmachinelearning.ManagedNetworkSettingsClientListResponse{
		// 	ManagedNetworkListResult: armmachinelearning.ManagedNetworkListResult{
		// 		Value: []*armmachinelearning.ManagedNetworkSettingsPropertiesBasicResource{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/managedNetworks"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.MachineLearningServices/workspaces/aml-workspace-name/managedNetworks/default"),
		// 				Properties: &armmachinelearning.ManagedNetworkSettingsProperties{
		// 					ManagedNetwork: &armmachinelearning.ManagedNetworkSettingsEx{
		// 						FirewallPublicIPAddress: to.Ptr("22.22.131.49"),
		// 						FirewallSKU: to.Ptr(armmachinelearning.FirewallSKUStandard),
		// 						IsolationMode: to.Ptr(armmachinelearning.IsolationModeAllowOnlyApprovedOutbound),
		// 						NetworkID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						OutboundRules: map[string]armmachinelearning.OutboundRuleClassification{
		// 							"rule_name_1": &armmachinelearning.FqdnOutboundRule{
		// 								Type: to.Ptr(armmachinelearning.RuleTypeFQDN),
		// 								Category: to.Ptr(armmachinelearning.RuleCategoryUserDefined),
		// 								Destination: to.Ptr("destination_endpoint"),
		// 								Status: to.Ptr(armmachinelearning.RuleStatusActive),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armmachinelearning.ManagedNetworkProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
