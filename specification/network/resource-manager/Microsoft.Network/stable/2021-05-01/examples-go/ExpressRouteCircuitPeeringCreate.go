package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitPeeringCreate.json
func ExampleExpressRouteCircuitPeeringsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteCircuitPeeringsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<circuit-name>",
		"<peering-name>",
		armnetwork.ExpressRouteCircuitPeering{
			Properties: &armnetwork.ExpressRouteCircuitPeeringPropertiesFormat{
				PeerASN:                    to.Ptr[int64](200),
				PrimaryPeerAddressPrefix:   to.Ptr("<primary-peer-address-prefix>"),
				SecondaryPeerAddressPrefix: to.Ptr("<secondary-peer-address-prefix>"),
				VlanID:                     to.Ptr[int32](200),
			},
		},
		&armnetwork.ExpressRouteCircuitPeeringsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
