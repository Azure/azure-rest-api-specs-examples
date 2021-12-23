Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.1.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeObjectMetadata_List.json
func ExampleIntegrationRuntimeObjectMetadataClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewIntegrationRuntimeObjectMetadataClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<integration-runtime-name>",
		&armsynapse.IntegrationRuntimeObjectMetadataListOptions{GetMetadataRequest: &armsynapse.GetSsisObjectMetadataRequest{
			MetadataPath: to.StringPtr("<metadata-path>"),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
```
