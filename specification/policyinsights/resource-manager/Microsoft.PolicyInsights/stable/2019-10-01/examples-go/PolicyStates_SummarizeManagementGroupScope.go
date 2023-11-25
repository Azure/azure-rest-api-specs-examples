package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeManagementGroupScope.json
func ExamplePolicyStatesClient_SummarizeForManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForManagementGroup(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "myManagementGroup", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](0),
		Filter:    to.Ptr("PolicyDefinitionAction eq 'deny' or PolicyDefinitionAction eq 'audit'"),
		OrderBy:   nil,
		Select:    nil,
		From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00.000Z"); return t }()),
		To:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-06T18:00:00.000Z"); return t }()),
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
	// 	ODataContext: to.Ptr("https://management.azure.com/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 			},
	// 			Results: &armpolicyinsights.SummaryResults{
	// 				NonCompliantPolicies: to.Ptr[int32](68),
	// 				NonCompliantResources: to.Ptr[int32](15410),
	// 				PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("noncompliant"),
	// 						Count: to.Ptr[int32](68),
	// 				}},
	// 				PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("noncompliant"),
	// 						Count: to.Ptr[int32](14),
	// 				}},
	// 				QueryResultsURI: to.Ptr("https://management.azure.com/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2019-10-01&$from=2019-10-05 18:00:00Z&$to=2019-10-06 18:00:00Z&$filter=(PolicyDefinitionAction eq 'deny' or PolicyDefinitionAction eq 'audit') and IsCompliant eq false"),
	// 				ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("compliant"),
	// 						Count: to.Ptr[int32](1400),
	// 					},
	// 					{
	// 						ComplianceState: to.Ptr("noncompliant"),
	// 						Count: to.Ptr[int32](15410),
	// 				}},
	// 			},
	// 	}},
	// }
}
