Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatashare%2Farmdatashare%2Fv0.4.0/sdk/resourcemanager/datashare/armdatashare/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_Create.json
func ExampleDataSetsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatashare.NewDataSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		"<data-set-name>",
		&armdatashare.BlobDataSet{
			Kind: to.Ptr(armdatashare.DataSetKindBlob),
			Properties: &armdatashare.BlobProperties{
				ContainerName:      to.Ptr("<container-name>"),
				FilePath:           to.Ptr("<file-path>"),
				ResourceGroup:      to.Ptr("<resource-group>"),
				StorageAccountName: to.Ptr("<storage-account-name>"),
				SubscriptionID:     to.Ptr("<subscription-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
