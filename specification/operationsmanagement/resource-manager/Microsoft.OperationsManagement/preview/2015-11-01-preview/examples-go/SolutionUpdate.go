package armoperationsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionUpdate.json
func ExampleSolutionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSolutionsClient().BeginUpdate(ctx, "rg1", "solution1", armoperationsmanagement.SolutionPatch{
		Tags: map[string]*string{
			"Dept":        to.Ptr("IT"),
			"Environment": to.Ptr("Test"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Solution = armoperationsmanagement.Solution{
	// 	Name: to.Ptr("solution1"),
	// 	Type: to.Ptr("Microsoft.OperationsManagement/solution"),
	// 	ID: to.Ptr("subscriptions/subid/resourcegroups/rg1/providers/Microsoft.OperationsManagement/solutions/solution1"),
	// 	Location: to.Ptr("East US"),
	// 	Plan: &armoperationsmanagement.SolutionPlan{
	// 		Name: to.Ptr("name1"),
	// 		Product: to.Ptr("product1"),
	// 		PromotionCode: to.Ptr("promocode1"),
	// 		Publisher: to.Ptr("publisher1"),
	// 	},
	// 	Properties: &armoperationsmanagement.SolutionProperties{
	// 		ContainedResources: []*string{
	// 			to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource1"),
	// 			to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource2")},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ReferencedResources: []*string{
	// 				to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource2"),
	// 				to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource3")},
	// 				WorkspaceResourceID: to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/Microsoft.OperationalInsights/workspaces/ws1"),
	// 			},
	// 			Tags: map[string]*string{
	// 				"Dept": to.Ptr("IT"),
	// 				"Environment": to.Ptr("Test"),
	// 			},
	// 		}
}
