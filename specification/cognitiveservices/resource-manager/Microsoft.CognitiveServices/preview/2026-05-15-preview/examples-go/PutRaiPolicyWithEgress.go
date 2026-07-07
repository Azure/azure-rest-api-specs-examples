package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/PutRaiPolicyWithEgress.json
func ExampleRaiPoliciesClient_CreateOrUpdate_putRaiPolicyWithEgress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiPoliciesClient().CreateOrUpdate(ctx, "resourceGroupName", "accountName", "egress-baseline", armcognitiveservices.RaiPolicy{
		Properties: &armcognitiveservices.RaiPolicyProperties{
			BasePolicyName: to.Ptr("Microsoft.Default"),
			EgressPolicy: &armcognitiveservices.RaiEgressPolicyConfig{
				Mode:          to.Ptr(armcognitiveservices.RaiEgressModeEnforced),
				DefaultAction: to.Ptr(armcognitiveservices.RaiEgressDefaultActionDeny),
				Description:   to.Ptr("Corporate baseline egress policy for sandboxed agents"),
				Rules: []*armcognitiveservices.RaiEgressRule{
					{
						Name:        to.Ptr("allow-openai"),
						Description: to.Ptr("Allow traffic to OpenAI API"),
						RuleType:    to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
						Match: &armcognitiveservices.RaiEgressRuleMatch{
							Host: to.Ptr("*.openai.com"),
						},
						Action: &armcognitiveservices.RaiEgressRuleAction{
							ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeAllow),
						},
					},
					{
						Name:        to.Ptr("inject-auth-for-internal"),
						Description: to.Ptr("Inject managed identity token for internal services"),
						RuleType:    to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
						Match: &armcognitiveservices.RaiEgressRuleMatch{
							Host: to.Ptr("*.internal.contoso.com"),
						},
						Action: &armcognitiveservices.RaiEgressRuleAction{
							ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeTransform),
							Headers: []*armcognitiveservices.RaiEgressHeaderTransform{
								{
									Operation: to.Ptr(armcognitiveservices.RaiEgressHeaderOperationSet),
									Name:      to.Ptr("Authorization"),
									ValueRef: &armcognitiveservices.RaiEgressHeaderValueRef{
										ManagedIdentityRef: &armcognitiveservices.RaiEgressManagedIdentityRef{
											Resource: to.Ptr("https://internal.contoso.com/.default"),
											Format:   to.Ptr("Bearer {value}"),
										},
									},
								},
							},
						},
					},
					{
						Name:        to.Ptr("rewrite-legacy-api"),
						Description: to.Ptr("Rewrite legacy API hostname to new internal endpoint"),
						RuleType:    to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
						Match: &armcognitiveservices.RaiEgressRuleMatch{
							Host: to.Ptr("legacy-api.contoso.com"),
							Path: to.Ptr("/v1/*"),
						},
						Action: &armcognitiveservices.RaiEgressRuleAction{
							ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeRewrite),
							Rewrite: &armcognitiveservices.RaiEgressRewriteTarget{
								Scheme: to.Ptr(armcognitiveservices.RaiEgressSchemeHTTPS),
								Host:   to.Ptr("api-v2.internal.contoso.com"),
								Path:   to.Ptr("/v2/"),
							},
							Headers: []*armcognitiveservices.RaiEgressHeaderTransform{
								{
									Operation: to.Ptr(armcognitiveservices.RaiEgressHeaderOperationSet),
									Name:      to.Ptr("X-Forwarded-Host"),
									Value:     to.Ptr("legacy-api.contoso.com"),
								},
							},
						},
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
	// res = armcognitiveservices.RaiPoliciesClientCreateOrUpdateResponse{
	// 	RaiPolicy: armcognitiveservices.RaiPolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiPolicies/egress-baseline"),
	// 		Name: to.Ptr("egress-baseline"),
	// 		Properties: &armcognitiveservices.RaiPolicyProperties{
	// 			BasePolicyName: to.Ptr("Microsoft.Default"),
	// 			Type: to.Ptr(armcognitiveservices.RaiPolicyTypeUserManaged),
	// 			EgressPolicy: &armcognitiveservices.RaiEgressPolicyConfig{
	// 				Mode: to.Ptr(armcognitiveservices.RaiEgressModeEnforced),
	// 				DefaultAction: to.Ptr(armcognitiveservices.RaiEgressDefaultActionDeny),
	// 				Description: to.Ptr("Corporate baseline egress policy for sandboxed agents"),
	// 				Rules: []*armcognitiveservices.RaiEgressRule{
	// 					{
	// 						Name: to.Ptr("allow-openai"),
	// 						Description: to.Ptr("Allow traffic to OpenAI API"),
	// 						RuleType: to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
	// 						Match: &armcognitiveservices.RaiEgressRuleMatch{
	// 							Host: to.Ptr("*.openai.com"),
	// 						},
	// 						Action: &armcognitiveservices.RaiEgressRuleAction{
	// 							ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeAllow),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("inject-auth-for-internal"),
	// 						Description: to.Ptr("Inject managed identity token for internal services"),
	// 						RuleType: to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
	// 						Match: &armcognitiveservices.RaiEgressRuleMatch{
	// 							Host: to.Ptr("*.internal.contoso.com"),
	// 						},
	// 						Action: &armcognitiveservices.RaiEgressRuleAction{
	// 							ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeTransform),
	// 							Headers: []*armcognitiveservices.RaiEgressHeaderTransform{
	// 								{
	// 									Operation: to.Ptr(armcognitiveservices.RaiEgressHeaderOperationSet),
	// 									Name: to.Ptr("Authorization"),
	// 									ValueRef: &armcognitiveservices.RaiEgressHeaderValueRef{
	// 										ManagedIdentityRef: &armcognitiveservices.RaiEgressManagedIdentityRef{
	// 											Resource: to.Ptr("https://internal.contoso.com/.default"),
	// 											Format: to.Ptr("Bearer {value}"),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("rewrite-legacy-api"),
	// 						Description: to.Ptr("Rewrite legacy API hostname to new internal endpoint"),
	// 						RuleType: to.Ptr(armcognitiveservices.RaiEgressRuleTypeFqdn),
	// 						Match: &armcognitiveservices.RaiEgressRuleMatch{
	// 							Host: to.Ptr("legacy-api.contoso.com"),
	// 							Path: to.Ptr("/v1/*"),
	// 						},
	// 						Action: &armcognitiveservices.RaiEgressRuleAction{
	// 							ActionType: to.Ptr(armcognitiveservices.RaiEgressRuleActionTypeRewrite),
	// 							Rewrite: &armcognitiveservices.RaiEgressRewriteTarget{
	// 								Scheme: to.Ptr(armcognitiveservices.RaiEgressSchemeHTTPS),
	// 								Host: to.Ptr("api-v2.internal.contoso.com"),
	// 								Path: to.Ptr("/v2/"),
	// 							},
	// 							Headers: []*armcognitiveservices.RaiEgressHeaderTransform{
	// 								{
	// 									Operation: to.Ptr(armcognitiveservices.RaiEgressHeaderOperationSet),
	// 									Name: to.Ptr("X-Forwarded-Host"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiPolicies"),
	// 	},
	// }
}
