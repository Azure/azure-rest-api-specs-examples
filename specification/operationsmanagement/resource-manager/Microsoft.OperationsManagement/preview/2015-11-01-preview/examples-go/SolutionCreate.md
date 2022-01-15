Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationsmanagement%2Farmoperationsmanagement%2Fv0.3.0/sdk/resourcemanager/operationsmanagement/armoperationsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationsmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"
)

// x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionCreate.json
func ExampleSolutionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationsmanagement.NewSolutionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<solution-name>",
		armoperationsmanagement.Solution{
			Location: to.StringPtr("<location>"),
			Plan: &armoperationsmanagement.SolutionPlan{
				Name:          to.StringPtr("<name>"),
				Product:       to.StringPtr("<product>"),
				PromotionCode: to.StringPtr("<promotion-code>"),
				Publisher:     to.StringPtr("<publisher>"),
			},
			Properties: &armoperationsmanagement.SolutionProperties{
				ContainedResources: []*string{
					to.StringPtr("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource1"),
					to.StringPtr("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource2")},
				ReferencedResources: []*string{
					to.StringPtr("/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource2"),
					to.StringPtr("/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource3")},
				WorkspaceResourceID: to.StringPtr("<workspace-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
