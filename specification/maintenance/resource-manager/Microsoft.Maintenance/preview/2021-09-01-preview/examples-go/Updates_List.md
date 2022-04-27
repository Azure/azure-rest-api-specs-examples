Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaintenance%2Farmmaintenance%2Fv0.4.0/sdk/resourcemanager/maintenance/armmaintenance/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/Updates_List.json
func ExampleUpdatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmaintenance.NewUpdatesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager("<resource-group-name>",
		"<provider-name>",
		"<resource-type>",
		"<resource-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
