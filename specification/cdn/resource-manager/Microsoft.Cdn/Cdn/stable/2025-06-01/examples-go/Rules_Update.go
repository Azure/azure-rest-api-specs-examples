package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/Rules_Update.json
func ExampleRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRulesClient().BeginUpdate(ctx, "RG", "profile1", "ruleSet1", "rule1", armcdn.RuleUpdateParameters{
		Properties: &armcdn.RuleUpdatePropertiesParameters{
			Actions: []armcdn.DeliveryRuleActionClassification{
				&armcdn.DeliveryRuleResponseHeaderAction{
					Name: to.Ptr(armcdn.DeliveryRuleActionNameModifyResponseHeader),
					Parameters: &armcdn.HeaderActionParameters{
						HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
						HeaderName:   to.Ptr("X-CDN"),
						TypeName:     to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleHeaderActionParameters),
						Value:        to.Ptr("MSFT"),
					},
				},
			},
			Order: to.Ptr[int32](1),
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
	// res = armcdn.RulesClientUpdateResponse{
	// 	Rule: armcdn.Rule{
	// 		Name: to.Ptr("rule1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/rulesets/rules"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/rulesets/ruleSet1/rules/rule1"),
	// 		Properties: &armcdn.RuleProperties{
	// 			Actions: []armcdn.DeliveryRuleActionClassification{
	// 				&armcdn.DeliveryRuleResponseHeaderAction{
	// 					Name: to.Ptr(armcdn.DeliveryRuleActionNameModifyResponseHeader),
	// 					Parameters: &armcdn.HeaderActionParameters{
	// 						HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
	// 						HeaderName: to.Ptr("X-CDN"),
	// 						TypeName: to.Ptr(armcdn.DeliveryRuleActionParametersTypeDeliveryRuleHeaderActionParameters),
	// 						Value: to.Ptr("MSFT"),
	// 					},
	// 				},
	// 			},
	// 			Conditions: []armcdn.DeliveryRuleConditionClassification{
	// 				&armcdn.DeliveryRuleRequestMethodCondition{
	// 					Name: to.Ptr(armcdn.MatchVariableRequestMethod),
	// 					Parameters: &armcdn.RequestMethodMatchConditionParameters{
	// 						MatchValues: []*armcdn.RequestMethodMatchValue{
	// 							to.Ptr(armcdn.RequestMethodMatchValueGET),
	// 						},
	// 						NegateCondition: to.Ptr(false),
	// 						Operator: to.Ptr(armcdn.RequestMethodOperatorEqual),
	// 						Transforms: []*armcdn.Transform{
	// 						},
	// 						TypeName: to.Ptr(armcdn.DeliveryRuleConditionParametersTypeDeliveryRuleRequestMethodConditionParameters),
	// 					},
	// 				},
	// 			},
	// 			DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 			MatchProcessingBehavior: to.Ptr(armcdn.MatchProcessingBehaviorContinue),
	// 			Order: to.Ptr[int32](1),
	// 			ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
