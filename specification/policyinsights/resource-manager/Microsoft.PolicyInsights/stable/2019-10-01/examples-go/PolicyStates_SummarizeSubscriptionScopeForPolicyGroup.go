package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionScopeForPolicyGroup.json
func ExamplePolicyStatesClient_SummarizeForSubscription_summarizeAtSubscriptionScopeForAPolicyDefinitionGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForSubscription(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("'group1' IN PolicyDefinitionGroupNames"),
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
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/ab379b49-a579-4045-984e-1b249ab8b474"),
	// 								PolicyDefinitionReferenceID: to.Ptr("1595906828137356523"),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/ab379b49-a579-4045-984e-1b249ab8b474'"),
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
	// 							},
	// 							{
	// 								Effect: to.Ptr("audit"),
	// 								PolicyDefinitionGroupNames: []*string{
	// 									to.Ptr("group1")},
	// 									PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 									PolicyDefinitionReferenceID: to.Ptr("2134906828137356512"),
	// 									Results: &armpolicyinsights.SummaryResults{
	// 										NonCompliantResources: to.Ptr[int32](34),
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
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](510),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](34),
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
	// 												Count: to.Ptr[int32](2),
	// 										}},
	// 										PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](1),
	// 										}},
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01' and 'group1' IN PolicyDefinitionGroupNames"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](14),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](557),
	// 										}},
	// 									},
	// 							}},
	// 							PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad"),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantPolicies: to.Ptr[int32](1),
	// 								NonCompliantResources: to.Ptr[int32](557),
	// 								PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](2),
	// 								}},
	// 								PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](1),
	// 								}},
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](17),
	// 								}},
	// 							},
	// 					}},
	// 					Results: &armpolicyinsights.SummaryResults{
	// 						NonCompliantPolicies: to.Ptr[int32](1),
	// 						NonCompliantResources: to.Ptr[int32](557),
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
	// 						QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames"),
	// 						ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("compliant"),
	// 								Count: to.Ptr[int32](140),
	// 							},
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](17),
	// 						}},
	// 					},
	// 			}},
	// 		}
}
