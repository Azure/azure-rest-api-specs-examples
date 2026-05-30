package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/EvaluateDeploymentPolicies.json
func ExampleAccountsClient_EvaluateDeploymentPolicies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().EvaluateDeploymentPolicies(ctx, "resourceGroupName", "accountName", armcognitiveservices.EvaluateDeploymentPoliciesRequest{
		Deployments: []*armcognitiveservices.EvaluateDeploymentPoliciesDeployment{
			{
				Name: to.Ptr("gpt4o-deployment"),
				Properties: &armcognitiveservices.EvaluateDeploymentPoliciesDeploymentProperties{
					Model: &armcognitiveservices.DeploymentModel{
						Format:  to.Ptr("OpenAI"),
						Name:    to.Ptr("gpt-4o"),
						Version: to.Ptr("2024-11-20"),
					},
					RaiPolicyName: to.Ptr("Microsoft.DefaultV2"),
				},
			},
			{
				Name: to.Ptr("ada-embedding"),
				Properties: &armcognitiveservices.EvaluateDeploymentPoliciesDeploymentProperties{
					Model: &armcognitiveservices.DeploymentModel{
						Format:  to.Ptr("OpenAI"),
						Name:    to.Ptr("text-embedding-ada-002"),
						Version: to.Ptr("2"),
					},
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
	// res = armcognitiveservices.AccountsClientEvaluateDeploymentPoliciesResponse{
	// 	EvaluateDeploymentPoliciesResponse: armcognitiveservices.EvaluateDeploymentPoliciesResponse{
	// 		Results: map[string]*armcognitiveservices.DeploymentPolicyEvaluationResult{
	// 			"gpt4o-deployment": &armcognitiveservices.DeploymentPolicyEvaluationResult{
	// 				EvaluationOutcome: to.Ptr(armcognitiveservices.PolicyEvaluationOutcomeCompliant),
	// 				NonCompliantAssignments: []*armcognitiveservices.PolicyAssignmentEvaluationDetails{
	// 				},
	// 			},
	// 			"ada-embedding": &armcognitiveservices.DeploymentPolicyEvaluationResult{
	// 				EvaluationOutcome: to.Ptr(armcognitiveservices.PolicyEvaluationOutcomeNonCompliant),
	// 				NonCompliantAssignments: []*armcognitiveservices.PolicyAssignmentEvaluationDetails{
	// 					{
	// 						AssignmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policyAssignments/deny-gpt-models"),
	// 						PolicyDefinitionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policyDefinitions/deny-embedding-models"),
	// 						EvaluationOutcome: to.Ptr(armcognitiveservices.PolicyEvaluationOutcomeNonCompliant),
	// 						NonComplianceReason: to.Ptr("Model text-embedding-ada-002 is not allowed by policy."),
	// 						Effect: to.Ptr("Deny"),
	// 						ExpressionEvaluations: []*armcognitiveservices.PolicyExpressionEvaluationDetails{
	// 							{
	// 								Expression: to.Ptr("Microsoft.CognitiveServices/accounts/deployments/model.name"),
	// 								ExpressionKind: to.Ptr("Field"),
	// 								Operator: to.Ptr("notIn"),
	// 								Result: to.Ptr("True"),
	// 								TargetValue: to.Ptr("[\"gpt-4o\",\"gpt-4\"]"),
	// 								ExpressionValue: to.Ptr("text-embedding-ada-002"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
