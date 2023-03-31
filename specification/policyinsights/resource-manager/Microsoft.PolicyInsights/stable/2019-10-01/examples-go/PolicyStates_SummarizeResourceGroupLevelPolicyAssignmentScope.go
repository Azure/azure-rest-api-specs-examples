package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceGroupLevelPolicyAssignmentScope.json
func ExamplePolicyStatesClient_SummarizeForResourceGroupLevelPolicyAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForResourceGroupLevelPolicyAssignment(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "b7a1ca2596524e3ab19597f2", &armpolicyinsights.QueryOptions{Top: nil,
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/b7a1ca2596524e3ab19597f2/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/b7a1ca2596524e3ab19597f2/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 				{
	// 					PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b7a1ca2596524e3ab19597f2"),
	// 					PolicyDefinitions: []*armpolicyinsights.PolicyDefinitionSummary{
	// 						{
	// 							Effect: to.Ptr("audit"),
	// 							PolicyDefinitionGroupNames: []*string{
	// 								to.Ptr("group1")},
	// 								PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/4a0425e4-97bf-4ad0-ab36-145b94083c60"),
	// 								PolicyDefinitionReferenceID: to.Ptr("2134906828137356512"),
	// 								Results: &armpolicyinsights.SummaryResults{
	// 									NonCompliantResources: to.Ptr[int32](7),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/b7a1ca2596524e3ab19597f2/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-22 23:54:22Z&$to=2019-10-23 23:54:22Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b7a1ca2596524e3ab19597f2' and PolicyDefinitionId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/4a0425e4-97bf-4ad0-ab36-145b94083c60'"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](7),
	// 									}},
	// 								},
	// 						}},
	// 						PolicyGroups: []*armpolicyinsights.PolicyGroupSummary{
	// 							{
	// 								PolicyGroupName: to.Ptr("group1"),
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
	// 									QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-12 19:48:53Z&$to=2019-10-13 19:48:53Z&$filter='group1' IN PolicyDefinitionGroupNames and PolicySetDefinitiontId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad' and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/b7a1ca2596524e3ab19597f2' and 'group1' IN PolicyDefinitionGroupNames"),
	// 									ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 										{
	// 											ComplianceState: to.Ptr("compliant"),
	// 											Count: to.Ptr[int32](140),
	// 										},
	// 										{
	// 											ComplianceState: to.Ptr("noncompliant"),
	// 											Count: to.Ptr[int32](7),
	// 									}},
	// 								},
	// 						}},
	// 						PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad"),
	// 						Results: &armpolicyinsights.SummaryResults{
	// 							NonCompliantPolicies: to.Ptr[int32](1),
	// 							NonCompliantResources: to.Ptr[int32](7),
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
	// 							QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/b7a1ca2596524e3ab19597f2/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-22 23:54:22Z&$to=2019-10-23 23:54:22Z&$filter=IsCompliant eq false and PolicyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.authorization/policyassignments/b7a1ca2596524e3ab19597f2'"),
	// 							ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 								{
	// 									ComplianceState: to.Ptr("compliant"),
	// 									Count: to.Ptr[int32](140),
	// 								},
	// 								{
	// 									ComplianceState: to.Ptr("noncompliant"),
	// 									Count: to.Ptr[int32](7),
	// 							}},
	// 						},
	// 				}},
	// 				Results: &armpolicyinsights.SummaryResults{
	// 					NonCompliantPolicies: to.Ptr[int32](1),
	// 					NonCompliantResources: to.Ptr[int32](7),
	// 					PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 						{
	// 							ComplianceState: to.Ptr("noncompliant"),
	// 							Count: to.Ptr[int32](1),
	// 					}},
	// 					PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 						{
	// 							ComplianceState: to.Ptr("noncompliant"),
	// 							Count: to.Ptr[int32](1),
	// 					}},
	// 					QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/b7a1ca2596524e3ab19597f2/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-22 23:54:22Z&$to=2019-10-23 23:54:22Z&$filter=IsCompliant eq false"),
	// 					ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 						{
	// 							ComplianceState: to.Ptr("compliant"),
	// 							Count: to.Ptr[int32](140),
	// 						},
	// 						{
	// 							ComplianceState: to.Ptr("noncompliant"),
	// 							Count: to.Ptr[int32](7),
	// 					}},
	// 				},
	// 		}},
	// 	}
}
