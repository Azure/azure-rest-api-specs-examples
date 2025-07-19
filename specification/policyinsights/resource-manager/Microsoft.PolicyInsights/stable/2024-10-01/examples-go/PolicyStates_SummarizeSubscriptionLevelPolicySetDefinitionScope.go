package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicySetDefinitionScope.json
func ExamplePolicyStatesClient_SummarizeForPolicySetDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyStatesClient().SummarizeForPolicySetDefinition(ctx, armpolicyinsights.PolicyStatesSummaryResourceTypeLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "3e3807c1-65c9-49e0-a406-82d8ae3e338c", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyDefinitionAction eq 'deny'"),
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policySetDefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.Summary{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policySetDefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c/providers/Microsoft.PolicyInsights/policyStates/$metadata#summary/$entity"),
	// 			PolicyAssignments: []*armpolicyinsights.PolicyAssignmentSummary{
	// 			},
	// 			Results: &armpolicyinsights.SummaryResults{
	// 				NonCompliantPolicies: to.Ptr[int32](0),
	// 				NonCompliantResources: to.Ptr[int32](0),
	// 				PolicyDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("compliant"),
	// 						Count: to.Ptr[int32](1),
	// 				}},
	// 				PolicyGroupDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("compliant"),
	// 						Count: to.Ptr[int32](1),
	// 				}},
	// 				QueryResultsURI: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policySetDefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c/providers/Microsoft.PolicyInsights/policyStates/latest/queryResults?api-version=2024-10-01&$from=2019-10-05 18:00:00Z&$to=2019-10-06 18:00:00Z&$filter=(PolicyDefinitionAction eq 'deny') and IsCompliant eq false"),
	// 				ResourceDetails: []*armpolicyinsights.ComplianceDetail{
	// 					{
	// 						ComplianceState: to.Ptr("compliant"),
	// 						Count: to.Ptr[int32](140),
	// 				}},
	// 			},
	// 	}},
	// }
}
