Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbaremetalinfrastructure%2Farmbaremetalinfrastructure%2Fv0.1.0/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/README.md) on how to add the SDK to your project and authenticate.

```go
package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure"
)

// x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalOperations_List.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbaremetalinfrastructure.NewOperationsClient(cred, nil)
	_, err = client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
