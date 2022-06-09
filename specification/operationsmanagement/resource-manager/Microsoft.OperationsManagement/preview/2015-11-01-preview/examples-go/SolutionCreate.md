```go
package armoperationsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionCreate.json
func ExampleSolutionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armoperationsmanagement.NewSolutionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"solution1",
		armoperationsmanagement.Solution{
			Location: to.Ptr("East US"),
			Plan: &armoperationsmanagement.SolutionPlan{
				Name:          to.Ptr("name1"),
				Product:       to.Ptr("product1"),
				PromotionCode: to.Ptr("promocode1"),
				Publisher:     to.Ptr("publisher1"),
			},
			Properties: &armoperationsmanagement.SolutionProperties{
				ContainedResources: []*string{
					to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource1"),
					to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource2")},
				ReferencedResources: []*string{
					to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource2"),
					to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource3")},
				WorkspaceResourceID: to.Ptr("/subscriptions/sub2/resourceGroups/rg2/providers/Microsoft.OperationalInsights/workspaces/ws1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationsmanagement%2Farmoperationsmanagement%2Fv0.6.0/sdk/resourcemanager/operationsmanagement/armoperationsmanagement/README.md) on how to add the SDK to your project and authenticate.
