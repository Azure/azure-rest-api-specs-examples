package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-01-15-preview/ManagedNetwork/createOrUpdateRuleV2.json
func ExampleOutboundRuleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOutboundRuleClient().BeginCreateOrUpdate(ctx, "test-rg", "cognitive-account-name", "default", "rule_name_1", armcognitiveservices.OutboundRuleBasicResource{
		Properties: &armcognitiveservices.FqdnOutboundRule{
			Type:        to.Ptr(armcognitiveservices.RuleTypeFQDN),
			Category:    to.Ptr(armcognitiveservices.RuleCategoryUserDefined),
			Destination: to.Ptr("destination_endpoint"),
			Status:      to.Ptr(armcognitiveservices.RuleStatusActive),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.OutboundRuleClientCreateOrUpdateResponse{
	// 	OutboundRuleBasicResource: &armcognitiveservices.OutboundRuleBasicResource{
	// 		Name: to.Ptr("rule_name_1"),
	// 		Type: to.Ptr("accounts/outboundRules"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/cognitive-account-name/managedNetworks/default/outboundRules/rule_name_1"),
	// 		Properties: &armcognitiveservices.FqdnOutboundRule{
	// 			Type: to.Ptr(armcognitiveservices.RuleTypeFQDN),
	// 			Category: to.Ptr(armcognitiveservices.RuleCategoryUserDefined),
	// 			Destination: to.Ptr("destination_endpoint"),
	// 			Status: to.Ptr(armcognitiveservices.RuleStatusActive),
	// 		},
	// 	},
	// }
}
