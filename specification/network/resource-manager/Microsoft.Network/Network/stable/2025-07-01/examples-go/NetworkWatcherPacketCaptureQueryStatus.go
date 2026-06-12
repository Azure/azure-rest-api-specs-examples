package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkWatcherPacketCaptureQueryStatus.json
func ExamplePacketCapturesClient_BeginGetStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPacketCapturesClient().BeginGetStatus(ctx, "rg1", "nw1", "pc1", nil)
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
	// res = armnetwork.PacketCapturesClientGetStatusResponse{
	// 	PacketCaptureQueryStatusResult: armnetwork.PacketCaptureQueryStatusResult{
	// 		Name: to.Ptr("pc1"),
	// 		CaptureStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-07T12:35:24Z"); return t}()),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/packetCaptures/pc1"),
	// 		PacketCaptureError: []*armnetwork.PcError{
	// 		},
	// 		PacketCaptureStatus: to.Ptr(armnetwork.PcStatusStopped),
	// 		StopReason: to.Ptr("TimeExceeded"),
	// 	},
	// }
}
