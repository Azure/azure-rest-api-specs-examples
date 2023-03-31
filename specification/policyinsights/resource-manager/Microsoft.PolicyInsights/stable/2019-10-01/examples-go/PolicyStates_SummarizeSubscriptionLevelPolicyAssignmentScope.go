package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicyAssignmentScope.json
func ExamplePolicyStatesClient_SummarizeForSubscriptionLevelPolicyAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForSubscriptionLevelPolicyAssignment(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "ec8f9645-8ecb-4abb-9c0b-5292f19d4003", &armpolicyinsights.QueryOptions{Top: nil,
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionGroupNames: []*string{
	// 								to.Ptr("group1")},
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
	// 								PolicyDefinitionReferenceID: to.Ptr("2134906828137356512"),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:09:24Z&$to=2019-10-13 20:09:24Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](531),
	// 									}},
	// 								},
	// 							},
	// 							{
	// 								Effect: to.Ptr("audit"),
	// 								PolicyDefinitionGroupNames: []*string{
	// 									to.Ptr("group1")},
	// 									PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 									PolicyDefinitionReferenceID: to.Ptr("3424906828137356512"),
	// 									Results: &armpolicyinsights.SummaryResults{
	// 										NonCompliantResources: to.Ptr[int32](531),
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
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:09:24Z&$to=2019-10-13 20:09:24Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](140),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](531),
	// 										}},
	// 									},
	// 								},
	// 								{
	// 									Effect: to.Ptr("audit"),
	// 									PolicyDefinitionGroupNames: []*string{
	// 										to.Ptr("group1")},
	// 										PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/locationauditdefinition"),
	// 										PolicyDefinitionReferenceID: to.Ptr("7943906828137356512"),
	// 										Results: &armpolicyinsights.SummaryResults{
	// 											NonCompliantResources: to.Ptr[int32](220),
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
	// 											QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:09:24Z&$to=2019-10-13 20:09:24Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/locationauditdefinition'"),
	// 											ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("compliant"),
	// 													Count: to.Ptr[int32](140),
	// 												},
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](220),
	// 											}},
	// 										},
	// 									},
	// 									{
	// 										Effect: to.Ptr("audit"),
	// 										PolicyDefinitionGroupNames: []*string{
	// 											to.Ptr("group1")},
	// 											PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e3682"),
	// 											PolicyDefinitionReferenceID: to.Ptr("1234906828137356512"),
	// 											Results: &armpolicyinsights.SummaryResults{
	// 												NonCompliantResources: to.Ptr[int32](54),
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
	// 												QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:09:24Z&$to=2019-10-13 20:09:24Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e3682'"),
	// 												ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 													{
	// 														ComplianceState: to.Ptr("compliant"),
	// 														Count: to.Ptr[int32](140),
	// 													},
	// 													{
	// 														ComplianceState: to.Ptr("noncompliant"),
	// 														Count: to.Ptr[int32](54),
	// 												}},
	// 											},
	// 									}},
	// 									PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 										{
	// 											PolicyGroupName: to.Ptr("group1"),
	// 											Results: &armpolicyinsights.SummaryResults{
	// 												NonCompliantResources: to.Ptr[int32](100),
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
	// 												QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/1ef5d536aec743a0aa801c1a' and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/335cefd2-ab16-430f-b364-974a170eb1d5' and 'group1' IN PolicyDefinitionGroupNames"),
	// 												ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 													{
	// 														ComplianceState: to.Ptr("compliant"),
	// 														Count: to.Ptr[int32](140),
	// 													},
	// 													{
	// 														ComplianceState: to.Ptr("noncompliant"),
	// 														Count: to.Ptr[int32](100),
	// 												}},
	// 											},
	// 									}},
	// 									PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
	// 									Results: &armpolicyinsights.SummaryResults{
	// 										NonCompliantPolicies: to.Ptr[int32](4),
	// 										NonCompliantResources: to.Ptr[int32](531),
	// 										PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](4),
	// 										}},
	// 										PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](1),
	// 										}},
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:09:24Z&$to=2019-10-13 20:09:24Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531'"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](140),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](531),
	// 										}},
	// 									},
	// 							}},
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantPolicies: to.Ptr[int32](1),
	// 								NonCompliantResources: to.Ptr[int32](531),
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
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:09:24Z&$to=2019-10-13 20:09:24Z&$filter=IsCompliant eq false"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](531),
	// 								}},
	// 							},
	// 					}},
	// 				}
}
