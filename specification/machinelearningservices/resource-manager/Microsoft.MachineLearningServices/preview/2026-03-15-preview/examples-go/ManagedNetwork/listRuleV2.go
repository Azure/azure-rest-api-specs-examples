package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/ManagedNetwork/listRuleV2.json
func ExampleOutboundRuleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOutboundRuleClient().NewListPager("test-rg", "aml-workspace-name", "default", nil)
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
		// page = armmachinelearning.OutboundRuleClientListResponse{
		// 	OutboundRuleListResult: armmachinelearning.OutboundRuleListResult{
		// 		Value: []*armmachinelearning.OutboundRuleBasicResource{
		// 			{
		// 				Name: to.Ptr("rule_name_1"),
		// 				Type: to.Ptr("workspace/outboundRules"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.MachineLearningServices/workspaces/aml-workspace-name/managedNetworks/default/outboundRules/rule_name_1"),
		// 				Properties: &armmachinelearning.FqdnOutboundRule{
		// 					Type: to.Ptr(armmachinelearning.RuleTypeFQDN),
		// 					Category: to.Ptr(armmachinelearning.RuleCategoryRequired),
		// 					Destination: to.Ptr("destination_of_the_fqdn_rule"),
		// 					ErrorInformation: to.Ptr("Error message"),
		// 					Status: to.Ptr(armmachinelearning.RuleStatusInactive),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("rule_name_2"),
		// 				Type: to.Ptr("workspace/outboundRules"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.MachineLearningServices/workspaces/aml-workspace-name/managedNetworks/default/outboundRules/rule_name_2"),
		// 				Properties: &armmachinelearning.FqdnOutboundRule{
		// 					Type: to.Ptr(armmachinelearning.RuleTypeFQDN),
		// 					Category: to.Ptr(armmachinelearning.RuleCategoryRequired),
		// 					Destination: to.Ptr("destination_of_the_fqdn_rule"),
		// 					ErrorInformation: to.Ptr("Error message"),
		// 					Status: to.Ptr(armmachinelearning.RuleStatusInactive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
