package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherConnectivityCheck.json
func ExampleWatchersClient_BeginCheckConnectivity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginCheckConnectivity(ctx, "rg1", "nw1", armnetwork.ConnectivityParameters{
		Destination: &armnetwork.ConnectivityDestination{
			Address: to.Ptr("192.168.100.4"),
			Port:    to.Ptr[int32](3389),
		},
		PreferredIPVersion: to.Ptr(armnetwork.IPVersionIPv4),
		Source: &armnetwork.ConnectivitySource{
			ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectivityInformation = armnetwork.ConnectivityInformation{
	// 	AvgLatencyInMs: to.Ptr[int32](1),
	// 	ConnectionStatus: to.Ptr(armnetwork.ConnectionStatusConnected),
	// 	Hops: []*armnetwork.ConnectivityHop{
	// 		{
	// 			Type: to.Ptr("Source"),
	// 			Address: to.Ptr("10.1.1.4"),
	// 			ID: to.Ptr("7dbbe7aa-60ba-4650-831e-63d775d38e9e"),
	// 			Issues: []*armnetwork.ConnectivityIssue{
	// 			},
	// 			NextHopIDs: []*string{
	// 				to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9")},
	// 				ResourceID: to.Ptr("subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkInterfaces/nic0/ipConfigurations/ipconfig1"),
	// 			},
	// 			{
	// 				Type: to.Ptr("VirtualNetwork"),
	// 				Address: to.Ptr("192.168.100.4"),
	// 				ID: to.Ptr("75c8d819-b208-4584-a311-1aa45ce753f9"),
	// 				Issues: []*armnetwork.ConnectivityIssue{
	// 				},
	// 				NextHopIDs: []*string{
	// 				},
	// 				ResourceID: to.Ptr("subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
	// 		}},
	// 		MaxLatencyInMs: to.Ptr[int32](4),
	// 		MinLatencyInMs: to.Ptr[int32](1),
	// 		ProbesFailed: to.Ptr[int32](0),
	// 		ProbesSent: to.Ptr[int32](100),
	// 	}
}
