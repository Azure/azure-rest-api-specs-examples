package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceScopeExpandPolicyEvaluationDetails.json
func ExampleComponentPolicyStatesClient_ListQueryResultsForResource_queryLatestComponentPolicyStatesAtResourceScopeAndExpandPolicyEvaluationDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentPolicyStatesClient().ListQueryResultsForResource(ctx, "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster", armpolicyinsights.ComponentPolicyStatesResourceLatest, &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceOptions{Top: nil,
		OrderBy: nil,
		Select:  nil,
		From:    nil,
		To:      nil,
		Filter:  to.Ptr("componentType eq 'pod' AND componentId eq 'default/test-pod' AND componentName eq 'test-pod'"),
		Apply:   nil,
		Expand:  to.Ptr("PolicyEvaluationDetails"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentPolicyStatesQueryResults = armpolicyinsights.ComponentPolicyStatesQueryResults{
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest"),
	// 	ODataCount: to.Ptr[int32](1),
	// 	Value: []*armpolicyinsights.ComponentPolicyState{
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"complianceReasonCode": "tbd",
	// 			},
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providersMicrosoft.ContainerService/managedClusters/myCluster/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest/$entity"),
	// 			ComplianceState: to.Ptr("NonCompliant"),
	// 			ComponentID: to.Ptr("default/test-pod"),
	// 			ComponentName: to.Ptr("test-pod"),
	// 			ComponentType: to.Ptr("pod"),
	// 			PolicyAssignmentID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyAssignments/test"),
	// 			PolicyAssignmentName: to.Ptr("test"),
	// 			PolicyAssignmentOwner: to.Ptr("tbd"),
	// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 			PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 			PolicyDefinitionAction: to.Ptr("audit"),
	// 			PolicyDefinitionCategory: to.Ptr("tbd"),
	// 			PolicyDefinitionGroupNames: []*string{
	// 				to.Ptr("myGroup")},
	// 				PolicyDefinitionID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyDefinitions/ab108bc4-32df-4677-8b38-fa8b2905df56"),
	// 				PolicyDefinitionName: to.Ptr("Audit Kubernetes policy"),
	// 				PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 				PolicyEvaluationDetails: &armpolicyinsights.ComponentPolicyEvaluationDetails{
	// 					Reason: to.Ptr("tbd reason for evaluation outcome"),
	// 				},
	// 				ResourceGroup: to.Ptr("myResourceGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster"),
	// 				ResourceLocation: to.Ptr("eastus"),
	// 				ResourceType: to.Ptr("/Microsoft.ContainerService/managedClusters"),
	// 				SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			},
	// 			{
	// 				AdditionalProperties: map[string]any{
	// 					"complianceReasonCode": "tbd",
	// 				},
	// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providersMicrosoft.ContainerService/managedClusters/myCluster/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest/$entity"),
	// 				ComplianceState: to.Ptr("NonCompliant"),
	// 				ComponentID: to.Ptr("default/test-pod"),
	// 				ComponentName: to.Ptr("test-pod"),
	// 				ComponentType: to.Ptr("pod"),
	// 				PolicyAssignmentID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyAssignments/test2"),
	// 				PolicyAssignmentName: to.Ptr("test"),
	// 				PolicyAssignmentOwner: to.Ptr("tbd"),
	// 				PolicyAssignmentScope: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 				PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 				PolicyDefinitionAction: to.Ptr("audit"),
	// 				PolicyDefinitionCategory: to.Ptr("tbd"),
	// 				PolicyDefinitionGroupNames: []*string{
	// 					to.Ptr("myGroup")},
	// 					PolicyDefinitionID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyDefinitions/ab108bc4-32df-4677-8b38-fa8b2905df56"),
	// 					PolicyDefinitionName: to.Ptr("Audit Kubernetes policy"),
	// 					PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 					PolicyEvaluationDetails: &armpolicyinsights.ComponentPolicyEvaluationDetails{
	// 						Reason: to.Ptr("tbd reason for evaluation outcome"),
	// 					},
	// 					ResourceGroup: to.Ptr("myResourceGroup"),
	// 					ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster"),
	// 					ResourceLocation: to.Ptr("eastus"),
	// 					ResourceType: to.Ptr("/Microsoft.ContainerService/managedClusters"),
	// 					SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			}},
	// 		}
}
