Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.3.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesDeleteSavedSearches.json
func ExampleSavedSearchesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationalinsights.NewSavedSearchesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<saved-search-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
