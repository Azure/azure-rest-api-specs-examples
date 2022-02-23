Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.3.1/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/SchemaRegistry/SchemaRegistryCreate.json
func ExampleSchemaRegistryClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventhub.NewSchemaRegistryClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<schema-group-name>",
		armeventhub.SchemaGroup{
			Properties: &armeventhub.SchemaGroupProperties{
				GroupProperties:     map[string]*string{},
				SchemaCompatibility: armeventhub.SchemaCompatibility("Forward").ToPtr(),
				SchemaType:          armeventhub.SchemaType("Avro").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SchemaRegistryClientCreateOrUpdateResult)
}
```
