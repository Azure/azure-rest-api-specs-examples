package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/VirtualNetworkMoveIpConfigurations.json
func ExampleVirtualNetworksClient_BeginMoveIPConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginMoveIPConfigurations(ctx, "rg1", "test-vnet", armnetwork.MoveIPConfigurationsRequest{
		MoveIPConfigurationItems: []*armnetwork.MoveIPConfigurationItem{
			{
				SourceIPConfiguration: &armnetwork.MoveIPConfigurationResourceReference{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
				},
				TargetIPConfiguration: &armnetwork.MoveIPConfigurationResourceReference{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic2/ipConfigurations/ipconfig2"),
				},
			},
			{
				SourceIPConfiguration: &armnetwork.MoveIPConfigurationResourceReference{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic3/ipConfigurations/ipconfig3"),
				},
				TargetIPConfiguration: &armnetwork.MoveIPConfigurationResourceReference{
					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic4/ipConfigurations/ipconfig4"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworksClientMoveIPConfigurationsResponse{
	// }
}
