Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmobilenetwork%2Farmmobilenetwork%2Fv0.1.0/sdk/resourcemanager/mobilenetwork/armmobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/AttachedDataNetworkCreate.json
func ExampleAttachedDataNetworksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewAttachedDataNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		"<packet-core-data-plane-name>",
		"<attached-data-network-name>",
		armmobilenetwork.AttachedDataNetwork{
			Location: to.StringPtr("<location>"),
			Properties: &armmobilenetwork.AttachedDataNetworkPropertiesFormat{
				NaptConfiguration: &armmobilenetwork.NaptConfiguration{
					Enabled:       armmobilenetwork.NaptEnabled("Enabled").ToPtr(),
					PinholeLimits: to.Int32Ptr(65536),
					PinholeTimeouts: &armmobilenetwork.PinholeTimeouts{
						Icmp: to.Int32Ptr(60),
						TCP:  to.Int32Ptr(7440),
						UDP:  to.Int32Ptr(300),
					},
					PortRange: &armmobilenetwork.PortRange{
						MaxPort: to.Int32Ptr(65535),
						MinPort: to.Int32Ptr(1024),
					},
					PortReuseHoldTime: &armmobilenetwork.PortReuseHoldTimes{
						TCP: to.Int32Ptr(120),
						UDP: to.Int32Ptr(60),
					},
				},
				UserEquipmentAddressPoolPrefix: []*string{
					to.StringPtr("2.2.0.0/16")},
				UserEquipmentStaticAddressPoolPrefix: []*string{
					to.StringPtr("2.4.0.0/16")},
				UserPlaneDataInterface: &armmobilenetwork.InterfaceProperties{
					Name: to.StringPtr("<name>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttachedDataNetworksClientCreateOrUpdateResult)
}
```
