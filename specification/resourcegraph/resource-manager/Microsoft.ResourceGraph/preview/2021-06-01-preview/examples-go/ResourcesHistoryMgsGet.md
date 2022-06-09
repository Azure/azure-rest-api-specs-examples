```go
package armresourcegraph_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesHistoryMgsGet.json
func ExampleClient_ResourcesHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresourcegraph.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ResourcesHistory(ctx,
		armresourcegraph.ResourcesHistoryRequest{
			ManagementGroups: []*string{
				to.Ptr("e927f598-c1d4-4f72-8541-95d83a6a4ac8"),
				to.Ptr("ProductionMG")},
			Options: &armresourcegraph.ResourcesHistoryRequestOptions{
				Interval: &armresourcegraph.DateTimeInterval{
					End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T01:25:00.0000000Z"); return t }()),
					Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T01:00:00.0000000Z"); return t }()),
				},
			},
			Query: to.Ptr("where name =~ 'cpu-utilization' | project id, name, properties"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresourcegraph%2Farmresourcegraph%2Fv0.6.0/sdk/resourcemanager/resourcegraph/armresourcegraph/README.md) on how to add the SDK to your project and authenticate.
