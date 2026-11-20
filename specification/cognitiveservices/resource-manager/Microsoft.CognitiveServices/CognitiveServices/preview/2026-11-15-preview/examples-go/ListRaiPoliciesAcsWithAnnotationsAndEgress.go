package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/ListRaiPoliciesAcsWithAnnotationsAndEgress.json
func ExampleRaiPoliciesClient_NewListPager_listAcsPoliciesWithOptionalAnnotationsAndEgressCompanions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRaiPoliciesClient().NewListPager("resource-group", "safety-account", nil)
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
		// page = armcognitiveservices.RaiPoliciesClientListResponse{
		// 	RaiPolicyListResult: armcognitiveservices.RaiPolicyListResult{
		// 		Value: []*armcognitiveservices.RaiPolicy{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiPolicies/annotated-egress-guard"),
		// 				Name: to.Ptr("annotated-egress-guard"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiPolicies"),
		// 				Etag: to.Ptr("\"00000000-0000-0000-0000-000000000006\""),
		// 				Properties: &armcognitiveservices.RaiPolicyProperties{
		// 					Type: to.Ptr(armcognitiveservices.RaiPolicyTypeUserManaged),
		// 					Format: to.Ptr(armcognitiveservices.RaiPolicyFormatACS),
		// 					Mode: to.Ptr(armcognitiveservices.RaiPolicyModeAsynchronousFilter),
		// 					Acs: &armcognitiveservices.RaiAcsManifest{
		// 						AgentControlSpecificationVersion: to.Ptr("0.4.0-alpha.1"),
		// 						Metadata: map[string]any{
		// 							"description": "Illustrative annotation authoring; supported annotators and configuration depend on the evaluation host.",
		// 						},
		// 						Policies: map[string]*armcognitiveservices.RaiAcsRegoPolicyDefinition{
		// 							"agent-guard": &armcognitiveservices.RaiAcsRegoPolicyDefinition{
		// 								Type: to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
		// 								Query: to.Ptr("data.agent_guard.verdict"),
		// 							},
		// 						},
		// 						Annotators: map[string]*armcognitiveservices.RaiAcsAnnotatorDefinition{
		// 							"risk": &armcognitiveservices.RaiAcsAnnotatorDefinition{
		// 								Type: to.Ptr(armcognitiveservices.RaiAcsAnnotatorTypeLlm),
		// 								AdditionalProperties: map[string]any{
		// 								"provider": "azure_openai",
		// 								"endpoint": "https://synthetic-judge.openai.azure.com/",
		// 								"deployment": "judge",
		// 								"api_version": "2024-10-21",
		// 								"result_schema": "acs.judge.observation/v2",
		// 								"system_prompt": "Classify the complete subject for harmful intent. Score from 0 to 10: 0 is benign, 5 is ambiguous or potentially harmful, and 10 is clear harmful intent. Use label clear for scores below 4, review for scores from 4 to below 7, and high_risk for scores of 7 or higher. Report observations only; Rego decides the action.",
		// 								"timeout_ms": 10000,
		// 							},
		// 							},
		// 						},
		// 						InterventionPoints: &armcognitiveservices.RaiAcsInterventionPoints{
		// 							AgentStartup: &armcognitiveservices.RaiAcsAgentStartupInterventionPoint{
		// 								PolicyTarget: to.Ptr("$snap.agent_init"),
		// 								PolicyTargetKind: to.Ptr("agent_init"),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 							},
		// 							Input: &armcognitiveservices.RaiAcsInterventionPoint{
		// 								PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetInput),
		// 								PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindUserInput),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 								Annotations: map[string]*armcognitiveservices.RaiAcsAnnotationBinding{
		// 									"risk": &armcognitiveservices.RaiAcsAnnotationBinding{
		// 										From: to.Ptr("$target.content"),
		// 									},
		// 								},
		// 							},
		// 							PreModelCall: &armcognitiveservices.RaiAcsPreModelCallInterventionPoint{
		// 								PolicyTarget: to.Ptr("$snap.messages"),
		// 								PolicyTargetKind: to.Ptr("messages"),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 							},
		// 							PostModelCall: &armcognitiveservices.RaiAcsPostModelCallInterventionPoint{
		// 								PolicyTarget: to.Ptr("$snap.response"),
		// 								PolicyTargetKind: to.Ptr("model_response"),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 							},
		// 							PreToolCall: &armcognitiveservices.RaiAcsToolInterventionPoint{
		// 								PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetToolArguments),
		// 								PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindToolArguments),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 							},
		// 							PostToolCall: &armcognitiveservices.RaiAcsToolInterventionPoint{
		// 								PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetToolResult),
		// 								PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindToolResult),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 							},
		// 							Output: &armcognitiveservices.RaiAcsInterventionPoint{
		// 								PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetOutput),
		// 								PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindAssistantOutput),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 								Annotations: map[string]*armcognitiveservices.RaiAcsAnnotationBinding{
		// 									"risk": &armcognitiveservices.RaiAcsAnnotationBinding{
		// 										From: to.Ptr("$target.content"),
		// 									},
		// 								},
		// 							},
		// 							AgentShutdown: &armcognitiveservices.RaiAcsAgentShutdownInterventionPoint{
		// 								PolicyTarget: to.Ptr("$snap.summary"),
		// 								PolicyTargetKind: to.Ptr("agent_summary"),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("agent-guard"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					AcsRegos: []*armcognitiveservices.RaiRegoReference{
		// 						{
		// 							RegoName: to.Ptr("agent-guard"),
		// 						},
		// 					},
		// 					EgressPolicy: &armcognitiveservices.RaiEgressPolicyConfig{
		// 						Mode: to.Ptr(armcognitiveservices.RaiEgressModeAudit),
		// 						DefaultAction: to.Ptr(armcognitiveservices.RaiEgressDefaultActionDeny),
		// 						Rules: []*armcognitiveservices.RaiEgressRule{
		// 							{
		// 								Name: to.Ptr("allow-docs"),
		// 								RuleType: to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
		// 								Match: &armcognitiveservices.RaiEgressRuleMatch{
		// 									Host: to.Ptr("learn.microsoft.com"),
		// 									Path: to.Ptr("/azure/*"),
		// 								},
		// 								Action: &armcognitiveservices.RaiEgressRuleAction{
		// 									ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeAllow),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("label-api"),
		// 								RuleType: to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
		// 								Match: &armcognitiveservices.RaiEgressRuleMatch{
		// 									Host: to.Ptr("api.contoso.com"),
		// 									Path: to.Ptr("/v1/*"),
		// 								},
		// 								Action: &armcognitiveservices.RaiEgressRuleAction{
		// 									ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeTransform),
		// 									Headers: []*armcognitiveservices.RaiEgressHeaderTransform{
		// 										{
		// 											Operation: to.Ptr(armcognitiveservices.RaiEgressHeaderOperationSet),
		// 											Name: to.Ptr("X-Workload-Class"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiPolicies/input-guard"),
		// 				Name: to.Ptr("input-guard"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiPolicies"),
		// 				Etag: to.Ptr("\"00000000-0000-0000-0000-000000000005\""),
		// 				Properties: &armcognitiveservices.RaiPolicyProperties{
		// 					Type: to.Ptr(armcognitiveservices.RaiPolicyTypeUserManaged),
		// 					Format: to.Ptr(armcognitiveservices.RaiPolicyFormatACS),
		// 					Acs: &armcognitiveservices.RaiAcsManifest{
		// 						AgentControlSpecificationVersion: to.Ptr("0.4.0-alpha.1"),
		// 						Policies: map[string]*armcognitiveservices.RaiAcsRegoPolicyDefinition{
		// 							"input-guard": &armcognitiveservices.RaiAcsRegoPolicyDefinition{
		// 								Type: to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
		// 								Query: to.Ptr("data.input_guard.verdict"),
		// 							},
		// 						},
		// 						InterventionPoints: &armcognitiveservices.RaiAcsInterventionPoints{
		// 							Input: &armcognitiveservices.RaiAcsInterventionPoint{
		// 								PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetInput),
		// 								PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindUserInput),
		// 								Policy: &armcognitiveservices.RaiAcsPolicyBinding{
		// 									ID: to.Ptr("input-guard"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					AcsRegos: []*armcognitiveservices.RaiRegoReference{
		// 						{
		// 							RegoName: to.Ptr("input-guard"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
