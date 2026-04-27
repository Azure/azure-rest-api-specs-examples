package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-01-15-preview/ManagedNetwork/listRuleV2.json
func ExampleOutboundRuleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOutboundRuleClient().NewListPager("test-rg", "cognitive-account-name", "default", nil)
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
		// page = armcognitiveservices.OutboundRuleClientListResponse{
		// 	OutboundRuleListResult: armcognitiveservices.OutboundRuleListResult{
		// 		Value: []*armcognitiveservices.OutboundRuleBasicResource{
		// 			{
		// 				Name: to.Ptr("rule_name_1"),
		// 				Type: to.Ptr("accounts/outboundRules"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/cognitive-account-name/managedNetworks/default/outboundRules/rule_name_1"),
		// 				Properties: &armcognitiveservices.FqdnOutboundRule{
		// 					Type: to.Ptr(armcognitiveservices.RuleTypeFQDN),
		// 					Category: to.Ptr(armcognitiveservices.RuleCategoryRequired),
		// 					Destination: to.Ptr("destination_of_the_fqdn_rule"),
		// 					ErrorInformation: to.Ptr("Error message"),
		// 					Status: to.Ptr(armcognitiveservices.RuleStatusInactive),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("rule_name_2"),
		// 				Type: to.Ptr("accounts/outboundRules"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/cognitive-account-name/managedNetworks/default/outboundRules/rule_name_2"),
		// 				Properties: &armcognitiveservices.FqdnOutboundRule{
		// 					Type: to.Ptr(armcognitiveservices.RuleTypeFQDN),
		// 					Category: to.Ptr(armcognitiveservices.RuleCategoryRequired),
		// 					Destination: to.Ptr("destination_of_the_fqdn_rule"),
		// 					ErrorInformation: to.Ptr("Error message"),
		// 					Status: to.Ptr(armcognitiveservices.RuleStatusInactive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
