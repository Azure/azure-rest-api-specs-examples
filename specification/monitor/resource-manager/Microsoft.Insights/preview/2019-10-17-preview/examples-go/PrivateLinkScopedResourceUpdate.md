Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2019-10-17-preview/examples/PrivateLinkScopedResourceUpdate.json
func ExamplePrivateLinkScopedResourcesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewPrivateLinkScopedResourcesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<scope-name>",
		"<name>",
		armmonitor.ScopedResource{
			Properties: &armmonitor.ScopedResourceProperties{
				LinkedResourceID: to.StringPtr("<linked-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkScopedResourcesClientCreateOrUpdateResult)
}
```
