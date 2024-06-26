package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRoutePortCreate.json
func ExampleExpressRoutePortsClient_BeginCreateOrUpdate_expressRoutePortCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRoutePortsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "portName", armnetwork.ExpressRoutePort{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
			BandwidthInGbps: to.Ptr[int32](100),
			BillingType:     to.Ptr(armnetwork.ExpressRoutePortsBillingTypeUnlimitedData),
			Encapsulation:   to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
			PeeringLocation: to.Ptr("peeringLocationName"),
		},
	}, nil)
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
