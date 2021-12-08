Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatashare%2Farmdatashare%2Fv0.1.0/sdk/resourcemanager/datashare/armdatashare/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_Create.json
func ExampleDataSetMappingsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewDataSetMappingsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-subscription-name>",
		"<data-set-mapping-name>",
		&armdatashare.BlobDataSetMapping{
			DataSetMapping: armdatashare.DataSetMapping{
				Kind: armdatashare.DataSetMappingKindBlob.ToPtr(),
			},
			Properties: &armdatashare.BlobMappingProperties{
				ContainerName:      to.StringPtr("<container-name>"),
				DataSetID:          to.StringPtr("<data-set-id>"),
				FilePath:           to.StringPtr("<file-path>"),
				ResourceGroup:      to.StringPtr("<resource-group>"),
				StorageAccountName: to.StringPtr("<storage-account-name>"),
				SubscriptionID:     to.StringPtr("<subscription-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataSetMappingClassification.GetDataSetMapping().ID: %s\n", *res.GetDataSetMapping().ID)
}
```
