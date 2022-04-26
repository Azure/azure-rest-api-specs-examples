Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicelinker%2Farmservicelinker%2Fv0.4.0/sdk/resourcemanager/servicelinker/armservicelinker/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/Link.json
func ExampleLinkerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armservicelinker.NewLinkerClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-uri>",
		"<linker-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
