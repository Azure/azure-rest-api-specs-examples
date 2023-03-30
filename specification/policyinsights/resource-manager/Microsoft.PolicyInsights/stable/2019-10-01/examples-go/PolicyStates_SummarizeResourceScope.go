package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceScope.json
func ExamplePolicyStatesClient_SummarizeForResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForResource(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](2),
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SummarizeResults = armpolicyinsights.SummarizeResults{
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/8174043a1e2849179635b874"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policydefinitions/72c0c41a-c752-4bc0-9c61-0d6adc567066"),
	// 							PolicyDefinitionReferenceID: to.Ptr(""),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantResources: to.Ptr[int32](1),
	// 								PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:59:17Z&$to=2019-10-13 19:59:17Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/8174043a1e2849179635b874' and PolicyDefinitionId eq '/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policydefinitions/72c0c41a-c752-4bc0-9c61-0d6adc567066'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 							},
	// 					}},
	// 					PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 						{
	// 							PolicyGroupName: to.Ptr(""),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantResources: to.Ptr[int32](7),
	// 								PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/8174043a1e2849179635b874'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 							},
	// 					}},
	// 					PolicySetDefinitionID: to.Ptr(""),
	// 					Results: &armpolicyinsights.SummaryResults{
	// 						NonCompliantPolicies: to.Ptr[int32](1),
	// 						NonCompliantResources: to.Ptr[int32](1),
	// 						PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 						PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 						QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:59:17Z&$to=2019-10-13 19:59:17Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/8174043a1e2849179635b874'"),
	// 						ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("compliant"),
	// 								Count: to.Ptr[int32](140),
	// 							},
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 					},
	// 				},
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
	// 							PolicyDefinitionReferenceID: to.Ptr(""),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantResources: to.Ptr[int32](1),
	// 								PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:59:17Z&$to=2019-10-13 19:59:17Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a' and PolicyDefinitionId eq '/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 							},
	// 					}},
	// 					PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 						{
	// 							PolicyGroupName: to.Ptr("group1"),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantResources: to.Ptr[int32](100),
	// 								PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a' and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/335cefd2-ab16-430f-b364-974a170eb1d5' and 'group1' IN PolicyDefinitionGroupNames"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 							},
	// 					}},
	// 					PolicySetDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policysetdefinitions/335cefd2-ab16-430f-b364-974a170eb1d5"),
	// 					Results: &armpolicyinsights.SummaryResults{
	// 						NonCompliantPolicies: to.Ptr[int32](1),
	// 						NonCompliantResources: to.Ptr[int32](1),
	// 						PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 						PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 						QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:59:17Z&$to=2019-10-13 19:59:17Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a'"),
	// 						ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("compliant"),
	// 								Count: to.Ptr[int32](140),
	// 							},
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 					},
	// 			}},
	// 			Results: &armpolicyinsights.SummaryResults{
	// 				NonCompliantPolicies: to.Ptr[int32](14),
	// 				NonCompliantResources: to.Ptr[int32](1),
	// 				PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("noncompliant"),
	// 						Count: to.Ptr[int32](2),
	// 				}},
	// 				PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("noncompliant"),
	// 						Count: to.Ptr[int32](1),
	// 				}},
	// 				QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/my-vault/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:59:17Z&$to=2019-10-13 19:59:17Z&$filter=IsCompliant eq false"),
	// 				ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("compliant"),
	// 						Count: to.Ptr[int32](140),
	// 					},
	// 					{
	// 						ComplianceState: to.Ptr("noncompliant"),
	// 						Count: to.Ptr[int32](1),
	// 				}},
	// 			},
	// 	}},
	// }
}
