package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/UpdateRaiPolicyAcs.json
func ExampleRaiPoliciesClient_CreateOrUpdate_replaceAnAcsPolicyConditionally() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiPoliciesClient().CreateOrUpdate(ctx, "resource-group", "safety-account", "agent-guard", armcognitiveservices.RaiPolicy{
		Properties: &armcognitiveservices.RaiPolicyProperties{
			Format: to.Ptr(armcognitiveservices.RaiPolicyFormatACS),
			Acs: &armcognitiveservices.RaiAcsManifest{
				AgentControlSpecificationVersion: to.Ptr("0.4.0-alpha.1"),
				Policies: map[string]*armcognitiveservices.RaiAcsRegoPolicyDefinition{
					"input-guard": {
						Type:  to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
						Query: to.Ptr("data.input_guard.verdict"),
					},
					"tool-guard": {
						Type:  to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
						Query: to.Ptr("data.tool_guard.verdict"),
					},
				},
				InterventionPoints: &armcognitiveservices.RaiAcsInterventionPoints{
					Input: &armcognitiveservices.RaiAcsInterventionPoint{
						PolicyTarget:     to.Ptr(armcognitiveservices.RaiAcsPolicyTargetInput),
						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindUserInput),
						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
							ID: to.Ptr("input-guard"),
							AacsModeration: &armcognitiveservices.RaiAcsModerationBindingExtension{
								SubjectFormat: to.Ptr(armcognitiveservices.RaiAcsModerationSubjectFormatText),
								HarmConfigs: []*armcognitiveservices.RaiAcsHarmConfiguration{
									{
										Category:     to.Ptr(armcognitiveservices.RaiAcsHarmCategoryHate),
										HarmConfigID: to.Ptr("Hate_Text_MultiSev"),
									},
									{
										Category: to.Ptr(armcognitiveservices.RaiAcsHarmCategoryPromptInjection),
									},
								},
							},
						},
					},
					PreToolCall: &armcognitiveservices.RaiAcsToolInterventionPoint{
						PolicyTarget:     to.Ptr(armcognitiveservices.RaiAcsPolicyTargetToolArguments),
						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindToolArguments),
						ToolNameFrom:     to.Ptr(armcognitiveservices.RaiAcsToolNameSelectorToolCallName),
						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
							ID: to.Ptr("tool-guard"),
						},
					},
					PostToolCall: &armcognitiveservices.RaiAcsToolInterventionPoint{
						PolicyTarget:     to.Ptr(armcognitiveservices.RaiAcsPolicyTargetToolResult),
						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindToolResult),
						ToolNameFrom:     to.Ptr(armcognitiveservices.RaiAcsToolNameSelectorToolCallName),
						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
							ID: to.Ptr("tool-guard"),
						},
					},
				},
				Tools: map[string]*armcognitiveservices.RaiAcsToolDefinition{
					"web_search": {
						ID:          to.Ptr("web_search"),
						Type:        to.Ptr("retrieval"),
						Description: to.Ptr("Search approved public documentation and approved partner sites"),
						SecurityLabels: []*string{
							to.Ptr("network_egress"),
							to.Ptr("untrusted_content"),
						},
						Clearance: to.Ptr("public"),
						AdditionalProperties: map[string]any{
							"allowed_domains": []any{
								"learn.microsoft.com",
								"support.microsoft.com",
							},
						},
					},
					"wire_transfer": {
						ID:   to.Ptr("wire_transfer"),
						Type: to.Ptr("financial_action"),
						SecurityLabels: []*string{
							to.Ptr("financial_write"),
						},
						Clearance: to.Ptr("confidential"),
					},
				},
			},
			AcsRegos: []*armcognitiveservices.RaiRegoReference{
				{
					RegoName: to.Ptr("input-guard"),
				},
				{
					RegoName: to.Ptr("tool-guard"),
				},
			},
			CustomBlocklists:              []*armcognitiveservices.CustomBlocklistConfig{},
			CustomExternalSafetyProviders: []*armcognitiveservices.RaiPolicyCustomExternalSafetyProviderReference{},
		},
	}, &armcognitiveservices.RaiPoliciesClientCreateOrUpdateOptions{
		IfMatch: to.Ptr("\"00000000-0000-0000-0000-000000000003\"")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiPoliciesClientCreateOrUpdateResponse{
	// 	EtagHeader: to.Ptr("\"00000000-0000-0000-0000-000000000004\""),
	// 	RaiPolicy: armcognitiveservices.RaiPolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiPolicies/agent-guard"),
	// 		Name: to.Ptr("agent-guard"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiPolicies"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000004\""),
	// 		Properties: &armcognitiveservices.RaiPolicyProperties{
	// 			Type: to.Ptr(armcognitiveservices.RaiPolicyTypeUserManaged),
	// 			Format: to.Ptr(armcognitiveservices.RaiPolicyFormatACS),
	// 			Acs: &armcognitiveservices.RaiAcsManifest{
	// 				AgentControlSpecificationVersion: to.Ptr("0.4.0-alpha.1"),
	// 				Policies: map[string]*armcognitiveservices.RaiAcsRegoPolicyDefinition{
	// 					"input-guard": &armcognitiveservices.RaiAcsRegoPolicyDefinition{
	// 						Type: to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
	// 						Query: to.Ptr("data.input_guard.verdict"),
	// 					},
	// 					"tool-guard": &armcognitiveservices.RaiAcsRegoPolicyDefinition{
	// 						Type: to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
	// 						Query: to.Ptr("data.tool_guard.verdict"),
	// 					},
	// 				},
	// 				InterventionPoints: &armcognitiveservices.RaiAcsInterventionPoints{
	// 					Input: &armcognitiveservices.RaiAcsInterventionPoint{
	// 						PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetInput),
	// 						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindUserInput),
	// 						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
	// 							ID: to.Ptr("input-guard"),
	// 							AacsModeration: &armcognitiveservices.RaiAcsModerationBindingExtension{
	// 								SubjectFormat: to.Ptr(armcognitiveservices.RaiAcsModerationSubjectFormatText),
	// 								HarmConfigs: []*armcognitiveservices.RaiAcsHarmConfiguration{
	// 									{
	// 										Category: to.Ptr(armcognitiveservices.RaiAcsHarmCategoryHate),
	// 										HarmConfigID: to.Ptr("Hate_Text_MultiSev"),
	// 									},
	// 									{
	// 										Category: to.Ptr(armcognitiveservices.RaiAcsHarmCategoryPromptInjection),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					PreToolCall: &armcognitiveservices.RaiAcsToolInterventionPoint{
	// 						PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetToolArguments),
	// 						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindToolArguments),
	// 						ToolNameFrom: to.Ptr(armcognitiveservices.RaiAcsToolNameSelectorToolCallName),
	// 						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
	// 							ID: to.Ptr("tool-guard"),
	// 						},
	// 					},
	// 					PostToolCall: &armcognitiveservices.RaiAcsToolInterventionPoint{
	// 						PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetToolResult),
	// 						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindToolResult),
	// 						ToolNameFrom: to.Ptr(armcognitiveservices.RaiAcsToolNameSelectorToolCallName),
	// 						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
	// 							ID: to.Ptr("tool-guard"),
	// 						},
	// 					},
	// 				},
	// 				Tools: map[string]*armcognitiveservices.RaiAcsToolDefinition{
	// 					"web_search": &armcognitiveservices.RaiAcsToolDefinition{
	// 						ID: to.Ptr("web_search"),
	// 						Type: to.Ptr("retrieval"),
	// 						Description: to.Ptr("Search approved public documentation and approved partner sites"),
	// 						SecurityLabels: []*string{
	// 							to.Ptr("network_egress"),
	// 							to.Ptr("untrusted_content"),
	// 						},
	// 						Clearance: to.Ptr("public"),
	// 						AdditionalProperties: map[string]any{
	// 						"allowed_domains": []any{
	// 							"learn.microsoft.com",
	// 							"support.microsoft.com",
	// 						},
	// 					},
	// 					},
	// 					"wire_transfer": &armcognitiveservices.RaiAcsToolDefinition{
	// 						ID: to.Ptr("wire_transfer"),
	// 						Type: to.Ptr("financial_action"),
	// 						SecurityLabels: []*string{
	// 							to.Ptr("financial_write"),
	// 						},
	// 						Clearance: to.Ptr("confidential"),
	// 					},
	// 				},
	// 			},
	// 			AcsRegos: []*armcognitiveservices.RaiRegoReference{
	// 				{
	// 					RegoName: to.Ptr("input-guard"),
	// 				},
	// 				{
	// 					RegoName: to.Ptr("tool-guard"),
	// 				},
	// 			},
	// 			CustomBlocklists: []*armcognitiveservices.CustomBlocklistConfig{
	// 			},
	// 			CustomExternalSafetyProviders: []*armcognitiveservices.RaiPolicyCustomExternalSafetyProviderReference{
	// 			},
	// 		},
	// 	},
	// }
}
