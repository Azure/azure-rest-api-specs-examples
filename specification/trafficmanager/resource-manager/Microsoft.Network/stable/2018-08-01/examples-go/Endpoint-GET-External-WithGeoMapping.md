```go
package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-GET-External-WithGeoMapping.json
func ExampleEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"azuresdkfornetautoresttrafficmanager2191",
		"azuresdkfornetautoresttrafficmanager8224",
		armtrafficmanager.EndpointTypeExternalEndpoints,
		"My%20external%20endpoint",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftrafficmanager%2Farmtrafficmanager%2Fv1.0.0/sdk/resourcemanager/trafficmanager/armtrafficmanager/README.md) on how to add the SDK to your project and authenticate.
