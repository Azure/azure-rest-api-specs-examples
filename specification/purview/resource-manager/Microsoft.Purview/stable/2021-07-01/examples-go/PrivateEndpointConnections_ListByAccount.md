Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpurview%2Farmpurview%2Fv0.2.1/sdk/resourcemanager/purview/armpurview/README.md) on how to add the SDK to your project and authenticate.

```go
package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateEndpointConnections_ListByAccount.json
func ExamplePrivateEndpointConnectionsClient_ListByAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpurview.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	pager := client.ListByAccount("<resource-group-name>",
		"<account-name>",
		&armpurview.PrivateEndpointConnectionsClientListByAccountOptions{SkipToken: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
