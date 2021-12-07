Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Faad%2Farmaad%2Fv0.1.0/sdk/resourcemanager/aad/armaad/README.md) on how to add the SDK to your project and authenticate.

```go
package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// x-ms-original-file: specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/preview/2017-04-01-preview/examples/listOperations.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armaad.NewOperationsClient(cred, nil)
	_, err = client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
