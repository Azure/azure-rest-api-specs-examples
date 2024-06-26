package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Rules_Update.json
func ExampleRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRulesClient().BeginUpdate(ctx, "RG", "profile1", "ruleSet1", "rule1", armcdn.RuleUpdateParameters{
		Properties: &armcdn.RuleUpdatePropertiesParameters{
			Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
				&armcdn.DeliveryRuleResponseHeaderAction{
					Name: to.Ptr(armcdn.DeliveryRuleActionModifyResponseHeader),
					Parameters: &armcdn.HeaderActionParameters{
						HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
						HeaderName:   to.Ptr("X-CDN"),
						TypeName:     to.Ptr(armcdn.HeaderActionParametersTypeNameDeliveryRuleHeaderActionParameters),
						Value:        to.Ptr("MSFT"),
					},
				}},
			Order: to.Ptr[int32](1),
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
	// res.Rule = armcdn.Rule{
	// 	Name: to.Ptr("rule1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/rulesets/rules"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/rulesets/ruleSet1/rules/rule1"),
	// 	Properties: &armcdn.RuleProperties{
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
	// 			&armcdn.DeliveryRuleResponseHeaderAction{
	// 				Name: to.Ptr(armcdn.DeliveryRuleActionModifyResponseHeader),
	// 				Parameters: &armcdn.HeaderActionParameters{
	// 					HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
	// 					HeaderName: to.Ptr("X-CDN"),
	// 					TypeName: to.Ptr(armcdn.HeaderActionParametersTypeNameDeliveryRuleHeaderActionParameters),
	// 					Value: to.Ptr("MSFT"),
	// 				},
	// 		}},
	// 		Conditions: []armcdn.DeliveryRuleConditionClassification{
	// 			&armcdn.DeliveryRuleRequestMethodCondition{
	// 				Name: to.Ptr(armcdn.MatchVariableRequestMethod),
	// 				Parameters: &armcdn.RequestMethodMatchConditionParameters{
	// 					MatchValues: []*armcdn.RequestMethodMatchConditionParametersMatchValuesItem{
	// 						to.Ptr(armcdn.RequestMethodMatchConditionParametersMatchValuesItemGET)},
	// 						NegateCondition: to.Ptr(false),
	// 						Operator: to.Ptr(armcdn.RequestMethodOperatorEqual),
	// 						Transforms: []*armcdn.Transform{
	// 						},
	// 						TypeName: to.Ptr(armcdn.RequestMethodMatchConditionParametersTypeNameDeliveryRuleRequestMethodConditionParameters),
	// 					},
	// 			}},
	// 			MatchProcessingBehavior: to.Ptr(armcdn.MatchProcessingBehaviorContinue),
	// 			Order: to.Ptr[int32](1),
	// 		},
	// 	}
}
