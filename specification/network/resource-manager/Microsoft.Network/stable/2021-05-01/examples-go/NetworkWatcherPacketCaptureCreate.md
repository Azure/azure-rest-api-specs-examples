Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherPacketCaptureCreate.json
func ExamplePacketCapturesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPacketCapturesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<network-watcher-name>",
		"<packet-capture-name>",
		armnetwork.PacketCapture{
			Properties: &armnetwork.PacketCaptureParameters{
				BytesToCapturePerPacket: to.Int64Ptr(10000),
				Filters: []*armnetwork.PacketCaptureFilter{
					{
						LocalIPAddress: to.StringPtr("<local-ipaddress>"),
						LocalPort:      to.StringPtr("<local-port>"),
						Protocol:       armnetwork.PcProtocol("TCP").ToPtr(),
					}},
				StorageLocation: &armnetwork.PacketCaptureStorageLocation{
					FilePath:    to.StringPtr("<file-path>"),
					StorageID:   to.StringPtr("<storage-id>"),
					StoragePath: to.StringPtr("<storage-path>"),
				},
				Target:               to.StringPtr("<target>"),
				TimeLimitInSeconds:   to.Int32Ptr(100),
				TotalBytesPerSession: to.Int64Ptr(100000),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
