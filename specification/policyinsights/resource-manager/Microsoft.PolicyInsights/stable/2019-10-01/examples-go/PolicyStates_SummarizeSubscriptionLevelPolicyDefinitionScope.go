package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicyDefinitionScope.json
func ExamplePolicyStatesClient_SummarizeForPolicyDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForPolicyDefinition(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "24813039-7534-408a-9842-eb99f45721b1", &armpolicyinsights.QueryOptions{Top: nil,
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 							PolicyDefinitionReferenceID: to.Ptr(""),
	// 							Results: &armpolicyinsights.SummaryResults{
	// 								NonCompliantResources: to.Ptr[int32](558),
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
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](558),
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
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter=PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](140),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](558),
	// 								}},
	// 							},
	// 					}},
	// 					PolicySetDefinitionID: to.Ptr(""),
	// 					Results: &armpolicyinsights.SummaryResults{
	// 						NonCompliantPolicies: to.Ptr[int32](1),
	// 						NonCompliantResources: to.Ptr[int32](558),
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
	// 						QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7'"),
	// 						ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 							{
	// 								ComplianceState: to.Ptr("compliant"),
	// 								Count: to.Ptr[int32](140),
	// 							},
	// 							{
	// 								ComplianceState: to.Ptr("noncompliant"),
	// 								Count: to.Ptr[int32](558),
	// 						}},
	// 					},
	// 				},
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionGroupNames: []*string{
	// 								to.Ptr("group1")},
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 								PolicyDefinitionReferenceID: to.Ptr("2481303924813039"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](553),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](553),
	// 									}},
	// 								},
	// 						}},
	// 						PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 							{
	// 								PolicyGroupName: to.Ptr("group1"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](553),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29' and 'group1' IN PolicyDefinitionGroupNames"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](14),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](553),
	// 									}},
	// 								},
	// 						}},
	// 						PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c"),
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](1),
	// 							NonCompliantResources: to.Ptr[int32](553),
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
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29'"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](553),
	// 							}},
	// 						},
	// 					},
	// 					{
	// 						PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1"),
	// 						PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 							{
	// 								Effect: to.Ptr("audit"),
	// 								PolicyDefinitionGroupNames: []*string{
	// 									to.Ptr("group1")},
	// 									PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 									PolicyDefinitionReferenceID: to.Ptr(""),
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
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](14),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](531),
	// 										}},
	// 									},
	// 							}},
	// 							PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 								{
	// 									PolicyGroupName: to.Ptr("group1"),
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
	// 										QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1' and 'group1' IN PolicyDefinitionGroupNames"),
	// 										ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 											{
	// 												ComplianceState: to.Ptr("compliant"),
	// 												Count: to.Ptr[int32](14),
	// 											},
	// 											{
	// 												ComplianceState: to.Ptr("noncompliant"),
	// 												Count: to.Ptr[int32](531),
	// 										}},
	// 									},
	// 							}},
	// 							PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
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
	// 								QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1'"),
	// 								ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 									{
	// 										ComplianceState: to.Ptr("compliant"),
	// 										Count: to.Ptr[int32](14),
	// 									},
	// 									{
	// 										ComplianceState: to.Ptr("noncompliant"),
	// 										Count: to.Ptr[int32](531),
	// 								}},
	// 							},
	// 						},
	// 						{
	// 							PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531"),
	// 							PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 								{
	// 									Effect: to.Ptr("audit"),
	// 									PolicyDefinitionGroupNames: []*string{
	// 										to.Ptr("group1")},
	// 										PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 										PolicyDefinitionReferenceID: to.Ptr(""),
	// 										Results: &armpolicyinsights.SummaryResults{
	// 											NonCompliantResources: to.Ptr[int32](531),
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
	// 											QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1'"),
	// 											ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("compliant"),
	// 													Count: to.Ptr[int32](140),
	// 												},
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](531),
	// 											}},
	// 										},
	// 								}},
	// 								PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 									{
	// 										PolicyGroupName: to.Ptr("group1"),
	// 										Results: &armpolicyinsights.SummaryResults{
	// 											NonCompliantResources: to.Ptr[int32](531),
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
	// 											QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531' and 'group1' IN PolicyDefinitionGroupNames"),
	// 											ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 												{
	// 													ComplianceState: to.Ptr("compliant"),
	// 													Count: to.Ptr[int32](14),
	// 												},
	// 												{
	// 													ComplianceState: to.Ptr("noncompliant"),
	// 													Count: to.Ptr[int32](531),
	// 											}},
	// 										},
	// 								}},
	// 								PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantPolicies: to.Ptr[int32](1),
	// 									NonCompliantResources: to.Ptr[int32](531),
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531'"),
	// 								},
	// 						}},
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](4),
	// 							NonCompliantResources: to.Ptr[int32](561),
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
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 20:07:10Z&$to=2019-10-13 20:07:10Z&$filter=IsCompliant eq false"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](561),
	// 							}},
	// 						},
	// 				}},
	// 			}
}
