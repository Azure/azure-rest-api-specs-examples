Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatacatalog%2Farmdatacatalog%2Fv0.1.0/sdk/resourcemanager/datacatalog/armdatacatalog/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatacatalog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog"
)

// x-ms-original-file: specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/ListADCCatalogsByResourceGroup.json
func ExampleADCCatalogsClient_ListtByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatacatalog.NewADCCatalogsClient("<subscription-id>",
		"<catalog-name>", cred, nil)
	_, err = client.ListtByResourceGroup(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
