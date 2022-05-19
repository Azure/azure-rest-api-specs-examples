Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpeering%2Farmpeering%2Fv1.0.0/sdk/resourcemanager/peering/armpeering/README.md) on how to add the SDK to your project and authenticate.

```go
package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreatePeeringService.json
func ExampleServicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpeering.NewServicesClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rgName",
		"peeringServiceName",
		armpeering.Service{
			Location: to.Ptr("eastus"),
			Properties: &armpeering.ServiceProperties{
				PeeringServiceLocation:         to.Ptr("state1"),
				PeeringServiceProvider:         to.Ptr("serviceProvider1"),
				ProviderBackupPeeringLocation:  to.Ptr("peeringLocation2"),
				ProviderPrimaryPeeringLocation: to.Ptr("peeringLocation1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
