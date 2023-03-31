package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceGroupScope.json
func ExamplePolicyStatesClient_SummarizeForResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForResourceGroup(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", &armpolicyinsights.QueryOptions{Top: nil,
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/186044306c044a1d8c0ff76c"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionGroupNames: []*string{
	// 								to.Ptr("group1")},
	// 								PolicyDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
	// 								PolicyDefinitionReferenceID: to.Ptr(""),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](100),
	// 									PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](1),
	// 									}},
	// 									PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](1),
	// 									}},
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:55:09Z&$to=2019-10-13 19:55:09Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/186044306c044a1d8c0ff76c' and PolicyDefinitionId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](100),
	// 									}},
	// 								},
	// 						}},
	// 						PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 							{
	// 								PolicyGroupName: to.Ptr("group1"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](100),
	// 									PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](1),
	// 									}},
	// 									PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](1),
	// 									}},
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/186044306c044a1d8c0ff76c' and 'group1' IN PolicyDefinitionGroupNames"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](100),
	// 									}},
	// 								},
	// 						}},
	// 						PolicySetDefinitionID: to.Ptr(""),
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](1),
	// 							NonCompliantResources: to.Ptr[int32](55),
	// 							PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](2),
	// 							}},
	// 							PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:55:09Z&$to=2019-10-13 19:55:09Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/186044306c044a1d8c0ff76c'"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](55),
	// 							}},
	// 						},
	// 					},
	// 					{
	// 						PolicyAssignmentID: to.Ptr("/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a"),
	// 						PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 							{
	// 								Effect: to.Ptr("audit"),
	// 								PolicyDefinitionGroupNames: []*string{
	// 									to.Ptr("group1")},
	// 									PolicyDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
	// 									PolicyDefinitionReferenceID: to.Ptr(""),
	// 									Results: &armpolicyinsights.SummaryResults{
	// 										NonCompliantResources: to.Ptr[int32](55),
	// 										PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](1),
	// 										}},
	// 										PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](1),
	// 										}},
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:55:09Z&$to=2019-10-13 19:55:09Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a' and PolicyDefinitionId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d'"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](140),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](55),
	// 										}},
	// 									},
	// 							}},
	// 							PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 								{
	// 									PolicyGroupName: to.Ptr("group1"),
	// 									Results: &armpolicyinsights.SummaryResults{
	// 										NonCompliantResources: to.Ptr[int32](557),
	// 										PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](1),
	// 										}},
	// 										PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](1),
	// 										}},
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/335cefd2-ab16-430f-b364-974a170eb1d5' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a' and 'group1' IN PolicyDefinitionGroupNames"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](140),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](55),
	// 										}},
	// 									},
	// 							}},
	// 							PolicySetDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policysetdefinitions/335cefd2-ab16-430f-b364-974a170eb1d5"),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantPolicies: to.Ptr[int32](1),
	// 								NonCompliantResources: to.Ptr[int32](55),
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
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:55:09Z&$to=2019-10-13 19:55:09Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](55),
	// 								}},
	// 							},
	// 					}},
	// 					Results: &armpolicyinsights.SummaryResults{
	// 						NonCompliantPolicies: to.Ptr[int32](20),
	// 						NonCompliantResources: to.Ptr[int32](55),
	// 						PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](2),
	// 						}},
	// 						PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](1),
	// 						}},
	// 						QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/d0610b27-9663-4c05-89f8-5b4be01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:55:09Z&$to=2019-10-13 19:55:09Z&$filter=IsCompliant eq false"),
	// 						ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("compliant"),
	// 								Count: to.Ptr[int32](140),
	// 							},
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](55),
	// 						}},
	// 					},
	// 			}},
	// 		}
}
