Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Faad%2Farmaad%2Fv0.4.0/sdk/resourcemanager/aad/armaad/README.md) on how to add the SDK to your project and authenticate.

```go
package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateLinkResourceGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateLinkResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<policy-name>",
		"<group-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
