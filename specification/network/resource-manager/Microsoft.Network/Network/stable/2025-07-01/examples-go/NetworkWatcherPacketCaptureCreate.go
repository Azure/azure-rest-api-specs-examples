package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkWatcherPacketCaptureCreate.json
func ExamplePacketCapturesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPacketCapturesClient().BeginCreate(ctx, "rg1", "nw1", "pc1", armnetwork.PacketCapture{
		Properties: &armnetwork.PacketCaptureParameters{
			BytesToCapturePerPacket: to.Ptr[int64](10000),
			Filters: []*armnetwork.PacketCaptureFilter{
				{
					LocalIPAddress: to.Ptr("10.0.0.4"),
					LocalPort:      to.Ptr("80"),
					Protocol:       to.Ptr(armnetwork.PcProtocolTCP),
				},
			},
			StorageLocation: &armnetwork.PacketCaptureStorageLocation{
				FilePath:    to.Ptr("D:\\capture\\pc1.cap"),
				StorageID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore"),
				StoragePath: to.Ptr("https://mytestaccountname.blob.core.windows.net/capture/pc1.cap"),
			},
			Target:               to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
			TimeLimitInSeconds:   to.Ptr[int32](100),
			TotalBytesPerSession: to.Ptr[int64](100000),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
