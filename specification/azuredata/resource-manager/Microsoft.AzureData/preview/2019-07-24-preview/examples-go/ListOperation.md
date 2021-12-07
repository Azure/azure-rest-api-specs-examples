Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazuredata%2Farmazuredata%2Fv0.1.0/sdk/resourcemanager/azuredata/armazuredata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// x-ms-original-file: specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/ListOperation.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazuredata.NewOperationsClient(cred, nil)
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
```
