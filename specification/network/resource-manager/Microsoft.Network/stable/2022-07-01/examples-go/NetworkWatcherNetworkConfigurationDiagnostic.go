package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherNetworkConfigurationDiagnostic.json
func ExampleWatchersClient_BeginGetNetworkConfigurationDiagnostic() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetNetworkConfigurationDiagnostic(ctx, "rg1", "nw1", armnetwork.ConfigurationDiagnosticParameters{
		Profiles: []*armnetwork.ConfigurationDiagnosticProfile{
			{
				Destination:     to.Ptr("12.11.12.14"),
				DestinationPort: to.Ptr("12100"),
				Direction:       to.Ptr(armnetwork.DirectionInbound),
				Source:          to.Ptr("10.1.0.4"),
				Protocol:        to.Ptr("TCP"),
			}},
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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
