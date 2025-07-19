package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceGroupScopeGroupByComponentTypeWithAggregate.json
func ExampleComponentPolicyStatesClient_ListQueryResultsForResourceGroup_queryLatestComponentPolicyComplianceStateCountGroupedByComponentTypeAtResourceGroupScopeFilteredByGivenAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentPolicyStatesClient().ListQueryResultsForResourceGroup(ctx, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", armpolicyinsights.ComponentPolicyStatesResourceLatest, &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions{Top: nil,
		OrderBy: nil,
		Select:  nil,
		From:    nil,
		To:      nil,
		Filter:  to.Ptr("policyAssignmentId eq '/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'"),
		Apply:   to.Ptr("groupby((type,complianceState),aggregate($count as count))"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentPolicyStatesQueryResults = armpolicyinsights.ComponentPolicyStatesQueryResults{
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest"),
	// 	ODataCount: to.Ptr[int32](2),
	// 	Value: []*armpolicyinsights.ComponentPolicyState{
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"count": float64(26),
	// 			},
	// 			ComplianceState: to.Ptr("NonCompliant"),
	// 			ComponentType: to.Ptr("Certificate"),
	// 		},
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"count": float64(10),
	// 			},
	// 			ComplianceState: to.Ptr("Compliant"),
	// 			ComponentType: to.Ptr("Certificate"),
	// 	}},
	// }
}
