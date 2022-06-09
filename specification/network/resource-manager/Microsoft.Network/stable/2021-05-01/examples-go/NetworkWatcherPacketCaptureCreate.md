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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherPacketCaptureCreate.json
func ExamplePacketCapturesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewPacketCapturesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<network-watcher-name>",
		"<packet-capture-name>",
		armnetwork.PacketCapture{
			Properties: &armnetwork.PacketCaptureParameters{
				BytesToCapturePerPacket: to.Ptr[int64](10000),
				Filters: []*armnetwork.PacketCaptureFilter{
					{
						LocalIPAddress: to.Ptr("<local-ipaddress>"),
						LocalPort:      to.Ptr("<local-port>"),
						Protocol:       to.Ptr(armnetwork.PcProtocolTCP),
					}},
				StorageLocation: &armnetwork.PacketCaptureStorageLocation{
					FilePath:    to.Ptr("<file-path>"),
					StorageID:   to.Ptr("<storage-id>"),
					StoragePath: to.Ptr("<storage-path>"),
				},
				Target:               to.Ptr("<target>"),
				TimeLimitInSeconds:   to.Ptr[int32](100),
				TotalBytesPerSession: to.Ptr[int64](100000),
			},
		},
		&armnetwork.PacketCapturesClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
