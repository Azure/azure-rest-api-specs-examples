package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceGroupScopeWithFilterAndTop.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForResourceGroupPager_queryAtResourceGroupScopeUsingQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyTrackedResourcesClient().NewListQueryResultsForResourceGroupPager("myResourceGroup", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment' AND TrackedResourceId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource/nestedResourceType/TrackedResource1'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PolicyTrackedResourcesQueryResults = armpolicyinsights.PolicyTrackedResourcesQueryResults{
		// 	Value: []*armpolicyinsights.PolicyTrackedResource{
		// 		{
		// 			CreatedBy: &armpolicyinsights.TrackedResourceModificationDetails{
		// 				DeploymentID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Resources/deployments/deploymentName"),
		// 				DeploymentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-09T00:00:00.000Z"); return t}()),
		// 				PolicyDetails: &armpolicyinsights.PolicyDetails{
		// 					PolicyAssignmentID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/createdByPolicyAssignment"),
		// 					PolicyAssignmentScope: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/"),
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyDefinitions/createdByPolicyDefinition"),
		// 				},
		// 			},
		// 			LastModifiedBy: &armpolicyinsights.TrackedResourceModificationDetails{
		// 				DeploymentID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Resources/deployments/deploymentName"),
		// 				DeploymentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-09T00:00:00.000Z"); return t}()),
		// 				PolicyDetails: &armpolicyinsights.PolicyDetails{
		// 					PolicyAssignmentID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/lastModifiedByPolicyAssignment"),
		// 					PolicyAssignmentScope: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/"),
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyDefinitions/lastModifiedByPolicyDefinition"),
		// 				},
		// 			},
		// 			LastUpdateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-09T20:40:17.358Z"); return t}()),
		// 			PolicyDetails: &armpolicyinsights.PolicyDetails{
		// 				PolicyAssignmentDisplayName: to.Ptr("My Policy Assignment 1 Display name"),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment1"),
		// 				PolicyAssignmentScope: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/"),
		// 				PolicyDefinitionID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyDefinitions/myPolicyDefinition1"),
		// 				PolicyDefinitionReferenceID: to.Ptr("123ABC"),
		// 				PolicySetDefinitionID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/providers/Microsoft.Authorization/policySetDefinitions/mySetDefinition"),
		// 			},
		// 			TrackedResourceID: to.Ptr("/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource/nestedResourceType/TrackedResource1"),
		// 	}},
		// }
	}
}
