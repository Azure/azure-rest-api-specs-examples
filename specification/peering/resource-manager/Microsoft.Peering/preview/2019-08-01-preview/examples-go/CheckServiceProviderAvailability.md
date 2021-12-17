Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpeering%2Farmpeering%2Fv0.1.0/sdk/resourcemanager/peering/armpeering/README.md) on how to add the SDK to your project and authenticate.

```go
package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/CheckServiceProviderAvailability.json
func ExamplePeeringManagementClient_CheckServiceProviderAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewPeeringManagementClient("<subscription-id>", cred, nil)
	_, err = client.CheckServiceProviderAvailability(ctx,
		armpeering.CheckServiceProviderAvailabilityInput{
			PeeringServiceLocation: to.StringPtr("<peering-service-location>"),
			PeeringServiceProvider: to.StringPtr("<peering-service-provider>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
