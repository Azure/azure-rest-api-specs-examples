package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherIpFlowVerify.json
func ExampleWatchersClient_BeginVerifyIPFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginVerifyIPFlow(ctx, "rg1", "nw1", armnetwork.VerificationIPFlowParameters{
		Direction:        to.Ptr(armnetwork.DirectionOutbound),
		LocalIPAddress:   to.Ptr("10.2.0.4"),
		LocalPort:        to.Ptr("80"),
		RemoteIPAddress:  to.Ptr("121.10.1.1"),
		RemotePort:       to.Ptr("80"),
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
		Protocol:         to.Ptr(armnetwork.IPFlowProtocolTCP),
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
