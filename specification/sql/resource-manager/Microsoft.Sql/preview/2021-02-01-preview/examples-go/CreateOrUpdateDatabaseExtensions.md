Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CreateOrUpdateDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewDatabaseExtensionsClient("a1c0814d-3c18-4e1e-a247-c128c12b1677", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg_20cbe0f0-c2d9-4522-9177-5469aad53029",
		"srv_1ffd1cf8-9951-47fb-807d-a9c384763849",
		"878e303f-1ea0-4f17-aa3d-a22ac5e9da08",
		"polybaseimport",
		armsql.DatabaseExtensions{
			Properties: &armsql.DatabaseExtensionsProperties{
				OperationMode:  to.Ptr(armsql.OperationModePolybaseImport),
				StorageKey:     to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				StorageKeyType: to.Ptr(armsql.StorageKeyTypeStorageAccessKey),
				StorageURI:     to.Ptr("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
