package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyRestrictions_CheckAtResourceGroupScope.json
func ExamplePolicyRestrictionsClient_CheckAtResourceGroupScope_checkPolicyRestrictionsAtResourceGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyRestrictionsClient().CheckAtResourceGroupScope(ctx, "vmRg", armpolicyinsights.CheckRestrictionsRequest{
		PendingFields: []*armpolicyinsights.PendingField{
			{
				Field: to.Ptr("name"),
				Values: []*string{
					to.Ptr("myVMName")},
			},
			{
				Field: to.Ptr("location"),
				Values: []*string{
					to.Ptr("eastus"),
					to.Ptr("westus"),
					to.Ptr("westus2"),
					to.Ptr("westeurope")},
			},
			{
				Field: to.Ptr("tags"),
			}},
		ResourceDetails: &armpolicyinsights.CheckRestrictionsResourceDetails{
			APIVersion: to.Ptr("2019-12-01"),
			ResourceContent: map[string]any{
				"type": "Microsoft.Compute/virtualMachines",
				"properties": map[string]any{
					"priority": "Spot",
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckRestrictionsResult = armpolicyinsights.CheckRestrictionsResult{
	// 	ContentEvaluationResult: &armpolicyinsights.CheckRestrictionsResultContentEvaluationResult{
	// 		PolicyEvaluations: []*armpolicyinsights.PolicyEvaluationResult{
	// 			{
	// 				EffectDetails: &armpolicyinsights.PolicyEffectDetails{
	// 					PolicyEffect: to.Ptr("Deny"),
	// 				},
	// 				EvaluationDetails: &armpolicyinsights.CheckRestrictionEvaluationDetails{
	// 					EvaluatedExpressions: []*armpolicyinsights.ExpressionEvaluationDetails{
	// 						{
	// 							Path: to.Ptr("type"),
	// 							Expression: to.Ptr("type"),
	// 							ExpressionKind: to.Ptr("field"),
	// 							ExpressionValue: "microsoft.compute/virtualmachines",
	// 							Operator: to.Ptr("equals"),
	// 							Result: to.Ptr("True"),
	// 							TargetValue: "microsoft.compute/virtualmachines",
	// 					}},
	// 				},
	// 				EvaluationResult: to.Ptr("NonCompliant"),
	// 				PolicyInfo: &armpolicyinsights.PolicyReference{
	// 					PolicyAssignmentID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyAssignments/2FF66C37"),
	// 					PolicyDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyDefinitions/435CAE41"),
	// 					PolicyDefinitionReferenceID: to.Ptr("defref222"),
	// 					PolicySetDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policySetDefinitions/2162358E"),
	// 				},
	// 		}},
	// 	},
	// 	FieldRestrictions: []*armpolicyinsights.FieldRestrictions{
	// 		{
	// 			Field: to.Ptr("tags.newtag"),
	// 			Restrictions: []*armpolicyinsights.FieldRestriction{
	// 				{
	// 					DefaultValue: to.Ptr("defaultVal"),
	// 					Policy: &armpolicyinsights.PolicyReference{
	// 						PolicyAssignmentID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyAssignments/57DAC8A0"),
	// 						PolicyDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyDefinitions/1D0906C3"),
	// 						PolicyDefinitionReferenceID: to.Ptr("DefRef"),
	// 						PolicySetDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policySetDefinitions/05D92080"),
	// 					},
	// 					PolicyEffect: to.Ptr("Deny"),
	// 					Reason: to.Ptr("tags.newtag is required"),
	// 					Result: to.Ptr(armpolicyinsights.FieldRestrictionResultRequired),
	// 			}},
	// 		},
	// 		{
	// 			Field: to.Ptr("tags.environment"),
	// 			Restrictions: []*armpolicyinsights.FieldRestriction{
	// 				{
	// 					Policy: &armpolicyinsights.PolicyReference{
	// 						PolicyAssignmentID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyAssignments/7EB1508A"),
	// 						PolicyDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyDefinitions/30BD79F6"),
	// 						PolicyDefinitionReferenceID: to.Ptr("DefRef"),
	// 						PolicySetDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policySetDefinitions/735551F1"),
	// 					},
	// 					PolicyEffect: to.Ptr("Deny"),
	// 					Reason: to.Ptr("tags.environment is required"),
	// 					Result: to.Ptr(armpolicyinsights.FieldRestrictionResultRequired),
	// 					Values: []*string{
	// 						to.Ptr("Prod"),
	// 						to.Ptr("Int"),
	// 						to.Ptr("Test")},
	// 				}},
	// 			},
	// 			{
	// 				Field: to.Ptr("location"),
	// 				Restrictions: []*armpolicyinsights.FieldRestriction{
	// 					{
	// 						Policy: &armpolicyinsights.PolicyReference{
	// 							PolicyAssignmentID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyAssignments/1563EBD3"),
	// 							PolicyDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyDefinitions/0711CCC1"),
	// 							PolicyDefinitionReferenceID: to.Ptr("DefRef"),
	// 							PolicySetDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policySetDefinitions/1E17783A"),
	// 						},
	// 						PolicyEffect: to.Ptr("Deny"),
	// 						Reason: to.Ptr("location must be one of the following: eastus, westus, westus2"),
	// 						Result: to.Ptr(armpolicyinsights.FieldRestrictionResultDeny),
	// 						Values: []*string{
	// 							to.Ptr("west europe")},
	// 						},
	// 						{
	// 							Policy: &armpolicyinsights.PolicyReference{
	// 								PolicyAssignmentID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyAssignments/5382A69D"),
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policyDefinitions/25C9F66B"),
	// 								PolicyDefinitionReferenceID: to.Ptr("DefRef"),
	// 								PolicySetDefinitionID: to.Ptr("/subscriptions/d8db6de6-2b96-46af-b825-07aef2033c0b/providers/microsoft.authorization/policySetDefinitions/392D107B"),
	// 							},
	// 							PolicyEffect: to.Ptr("Deny"),
	// 							Reason: to.Ptr("location must be one of the following: westus2"),
	// 							Result: to.Ptr(armpolicyinsights.FieldRestrictionResultDeny),
	// 							Values: []*string{
	// 								to.Ptr("eastus"),
	// 								to.Ptr("westus")},
	// 						}},
	// 				}},
	// 			}
}
