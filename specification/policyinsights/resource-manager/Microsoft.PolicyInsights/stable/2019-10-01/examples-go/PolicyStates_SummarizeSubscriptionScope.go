package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionScope.json
func ExamplePolicyStatesClient_SummarizeForSubscription_summarizeAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForSubscription(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](5),
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionGroupNames: []*string{
	// 								to.Ptr("group1")},
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
	// 								PolicyDefinitionReferenceID: to.Ptr("1b249ab8b4741b249ab8b474"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](557),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](557),
	// 									}},
	// 								},
	// 						}},
	// 						PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 							{
	// 								PolicyGroupName: to.Ptr("group1"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](531),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01' and 'group1' IN PolicyDefinitionGroupNames"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](557),
	// 									}},
	// 								},
	// 						}},
	// 						PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad"),
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](1),
	// 							NonCompliantResources: to.Ptr[int32](557),
	// 							PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01'"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](557),
	// 							}},
	// 						},
	// 					},
	// 					{
	// 						PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7"),
	// 						PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 							{
	// 								Effect: to.Ptr("audit"),
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 								PolicyDefinitionReferenceID: to.Ptr(""),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](557),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](557),
	// 									}},
	// 								},
	// 						}},
	// 						PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 							{
	// 								PolicyGroupName: to.Ptr(""),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](557),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](14),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](557),
	// 									}},
	// 								},
	// 						}},
	// 						PolicySetDefinitionID: to.Ptr(""),
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](1),
	// 							NonCompliantResources: to.Ptr[int32](557),
	// 							PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7'"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](557),
	// 							}},
	// 						},
	// 					},
	// 					{
	// 						PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/dcda79d769674aea8bfcaa49"),
	// 						PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 							{
	// 								Effect: to.Ptr("audit"),
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
	// 								PolicyDefinitionReferenceID: to.Ptr(""),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](557),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/dcda79d769674aea8bfcaa49' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](557),
	// 									}},
	// 								},
	// 						}},
	// 						PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 							{
	// 								PolicyGroupName: to.Ptr(""),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](557),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/dcda79d769674aea8bfcaa49'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](14),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](557),
	// 									}},
	// 								},
	// 						}},
	// 						PolicySetDefinitionID: to.Ptr(""),
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](1),
	// 							NonCompliantResources: to.Ptr[int32](557),
	// 							PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](1),
	// 							}},
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/dcda79d769674aea8bfcaa49'"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](557),
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
	// 									PolicyDefinitionReferenceID: to.Ptr("2134906828137356512"),
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
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a' and PolicyDefinitionId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policydefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d'"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](140),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](557),
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
	// 												Count: to.Ptr[int32](557),
	// 										}},
	// 									},
	// 							}},
	// 							PolicySetDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policysetdefinitions/335cefd2-ab16-430f-b364-974a170eb1d5"),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantPolicies: to.Ptr[int32](1),
	// 								NonCompliantResources: to.Ptr[int32](557),
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
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/providers/microsoft.management/managementgroups/myManagementGroup/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](557),
	// 								}},
	// 							},
	// 						},
	// 						{
	// 							PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29"),
	// 							PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 								{
	// 									Effect: to.Ptr("audit"),
	// 									PolicyDefinitionGroupNames: []*string{
	// 										to.Ptr("group1")},
	// 										PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 										PolicyDefinitionReferenceID: to.Ptr("5434906828137356512"),
	// 										Results: &armpolicyinsights.SummaryResults{
	// 											NonCompliantResources: to.Ptr[int32](552),
	// 											PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](1),
	// 											}},
	// 											PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](1),
	// 											}},
	// 											QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 											ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("compliant"),
	// 													Count: to.Ptr[int32](140),
	// 												},
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](552),
	// 											}},
	// 										},
	// 									},
	// 									{
	// 										Effect: to.Ptr("audit"),
	// 										PolicyDefinitionGroupNames: []*string{
	// 											to.Ptr("group1")},
	// 											PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/locationauditdefinition"),
	// 											PolicyDefinitionReferenceID: to.Ptr("8724906828137356512"),
	// 											Results: &armpolicyinsights.SummaryResults{
	// 												NonCompliantResources: to.Ptr[int32](29),
	// 												PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 													{
	// 														ComplianceState: to.Ptr("noncompliant"),
	// 														Count: to.Ptr[int32](1),
	// 												}},
	// 												PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 													{
	// 														ComplianceState: to.Ptr("noncompliant"),
	// 														Count: to.Ptr[int32](1),
	// 												}},
	// 												QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/locationauditdefinition'"),
	// 												ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 													{
	// 														ComplianceState: to.Ptr("compliant"),
	// 														Count: to.Ptr[int32](140),
	// 													},
	// 													{
	// 														ComplianceState: to.Ptr("noncompliant"),
	// 														Count: to.Ptr[int32](29),
	// 												}},
	// 											},
	// 										},
	// 										{
	// 											Effect: to.Ptr("audit"),
	// 											PolicyDefinitionGroupNames: []*string{
	// 												to.Ptr("group1")},
	// 												PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e3682"),
	// 												PolicyDefinitionReferenceID: to.Ptr("7254906828137356512"),
	// 												Results: &armpolicyinsights.SummaryResults{
	// 													NonCompliantResources: to.Ptr[int32](2),
	// 													PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 														{
	// 															ComplianceState: to.Ptr("noncompliant"),
	// 															Count: to.Ptr[int32](1),
	// 													}},
	// 													PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 														{
	// 															ComplianceState: to.Ptr("noncompliant"),
	// 															Count: to.Ptr[int32](1),
	// 													}},
	// 													QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e3682'"),
	// 													ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 														{
	// 															ComplianceState: to.Ptr("compliant"),
	// 															Count: to.Ptr[int32](140),
	// 														},
	// 														{
	// 															ComplianceState: to.Ptr("noncompliant"),
	// 															Count: to.Ptr[int32](2),
	// 													}},
	// 												},
	// 										}},
	// 										PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 											{
	// 												PolicyGroupName: to.Ptr("group1"),
	// 												Results: &armpolicyinsights.SummaryResults{
	// 													NonCompliantResources: to.Ptr[int32](557),
	// 													PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 														{
	// 															ComplianceState: to.Ptr("noncompliant"),
	// 															Count: to.Ptr[int32](3),
	// 													}},
	// 													PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 														{
	// 															ComplianceState: to.Ptr("noncompliant"),
	// 															Count: to.Ptr[int32](1),
	// 													}},
	// 													QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29' and 'group1' IN PolicyDefinitionGroupNames"),
	// 													ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 														{
	// 															ComplianceState: to.Ptr("compliant"),
	// 															Count: to.Ptr[int32](140),
	// 														},
	// 														{
	// 															ComplianceState: to.Ptr("noncompliant"),
	// 															Count: to.Ptr[int32](552),
	// 													}},
	// 												},
	// 										}},
	// 										PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c"),
	// 										Results: &armpolicyinsights.SummaryResults{
	// 											NonCompliantPolicies: to.Ptr[int32](3),
	// 											NonCompliantResources: to.Ptr[int32](552),
	// 											PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](3),
	// 											}},
	// 											PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](1),
	// 											}},
	// 											QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29'"),
	// 											ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("compliant"),
	// 													Count: to.Ptr[int32](140),
	// 												},
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](552),
	// 											}},
	// 										},
	// 								}},
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantPolicies: to.Ptr[int32](40),
	// 									NonCompliantResources: to.Ptr[int32](619),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=IsCompliant eq false"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](619),
	// 									}},
	// 								},
	// 						}},
	// 					}
}
