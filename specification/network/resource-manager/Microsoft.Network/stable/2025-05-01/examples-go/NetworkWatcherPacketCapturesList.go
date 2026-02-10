package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherPacketCapturesList.json
func ExamplePacketCapturesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPacketCapturesClient().NewListPager("rg1", "nw1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PacketCaptureListResult = armnetwork.PacketCaptureListResult{
		// 	Value: []*armnetwork.PacketCaptureResult{
		// 		{
		// 			Name: to.Ptr("pc1"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/packetCaptures/pc1"),
		// 			Properties: &armnetwork.PacketCaptureResultProperties{
		// 				BytesToCapturePerPacket: to.Ptr[int64](10000),
		// 				Filters: []*armnetwork.PacketCaptureFilter{
		// 					{
		// 						LocalIPAddress: to.Ptr("10.0.0.4"),
		// 						LocalPort: to.Ptr("80"),
		// 						Protocol: to.Ptr(armnetwork.PcProtocolTCP),
		// 				}},
		// 				StorageLocation: &armnetwork.PacketCaptureStorageLocation{
		// 					FilePath: to.Ptr("D:\\capture\\pc1.cap"),
		// 					StorageID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore"),
		// 					StoragePath: to.Ptr("https://mytestaccountname.blob.core.windows.net/capture/pc1.cap"),
		// 				},
		// 				Target: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 				TimeLimitInSeconds: to.Ptr[int32](100),
		// 				TotalBytesPerSession: to.Ptr[int64](100000),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pc2"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/packetCaptures/pc2"),
		// 			Properties: &armnetwork.PacketCaptureResultProperties{
		// 				BytesToCapturePerPacket: to.Ptr[int64](10000),
		// 				Filters: []*armnetwork.PacketCaptureFilter{
		// 				},
		// 				StorageLocation: &armnetwork.PacketCaptureStorageLocation{
		// 					FilePath: to.Ptr("D:\\capture\\pc2.cap"),
		// 					StorageID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore"),
		// 					StoragePath: to.Ptr("https://mytestaccountname.blob.core.windows.net/capture/pc2.cap"),
		// 				},
		// 				Target: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 				TimeLimitInSeconds: to.Ptr[int32](100),
		// 				TotalBytesPerSession: to.Ptr[int64](100000),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
