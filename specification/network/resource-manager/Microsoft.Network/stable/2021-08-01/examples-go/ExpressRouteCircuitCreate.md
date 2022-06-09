```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitCreate.json
func ExampleExpressRouteCircuitsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteCircuitsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"circuitName",
		armnetwork.ExpressRouteCircuit{
			Location: to.Ptr("Brazil South"),
			Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
				AllowClassicOperations: to.Ptr(false),
				Authorizations:         []*armnetwork.ExpressRouteCircuitAuthorization{},
				Peerings:               []*armnetwork.ExpressRouteCircuitPeering{},
				ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
					BandwidthInMbps:     to.Ptr[int32](200),
					PeeringLocation:     to.Ptr("Silicon Valley"),
					ServiceProviderName: to.Ptr("Equinix"),
				},
			},
			SKU: &armnetwork.ExpressRouteCircuitSKU{
				Name:   to.Ptr("Standard_MeteredData"),
				Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
				Tier:   to.Ptr(armnetwork.ExpressRouteCircuitSKUTierStandard),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
