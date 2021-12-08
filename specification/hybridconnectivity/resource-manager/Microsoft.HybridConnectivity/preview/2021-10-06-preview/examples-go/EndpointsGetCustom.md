Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridconnectivity%2Farmhybridconnectivity%2Fv0.1.0/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2021-10-06-preview/examples/EndpointsGetCustom.json
func ExampleEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridconnectivity.NewEndpointsClient(cred, nil)
	res, err := client.Get(ctx,
		"<resource-uri>",
		"<endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("EndpointResource.ID: %s\n", *res.ID)
}
```
