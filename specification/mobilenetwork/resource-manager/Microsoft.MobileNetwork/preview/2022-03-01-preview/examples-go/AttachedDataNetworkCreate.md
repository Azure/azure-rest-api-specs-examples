```go
package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/AttachedDataNetworkCreate.json
func ExampleAttachedDataNetworksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewAttachedDataNetworksClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"TestPacketCoreCP",
		"TestPacketCoreDP",
		"TestAttachedDataNetwork",
		armmobilenetwork.AttachedDataNetwork{
			Location: to.Ptr("eastus"),
			Properties: &armmobilenetwork.AttachedDataNetworkPropertiesFormat{
				NaptConfiguration: &armmobilenetwork.NaptConfiguration{
					Enabled:       to.Ptr(armmobilenetwork.NaptEnabledEnabled),
					PinholeLimits: to.Ptr[int32](65536),
					PinholeTimeouts: &armmobilenetwork.PinholeTimeouts{
						Icmp: to.Ptr[int32](60),
						TCP:  to.Ptr[int32](7440),
						UDP:  to.Ptr[int32](300),
					},
					PortRange: &armmobilenetwork.PortRange{
						MaxPort: to.Ptr[int32](65535),
						MinPort: to.Ptr[int32](1024),
					},
					PortReuseHoldTime: &armmobilenetwork.PortReuseHoldTimes{
						TCP: to.Ptr[int32](120),
						UDP: to.Ptr[int32](60),
					},
				},
				UserEquipmentAddressPoolPrefix: []*string{
					to.Ptr("2.2.0.0/16")},
				UserEquipmentStaticAddressPoolPrefix: []*string{
					to.Ptr("2.4.0.0/16")},
				UserPlaneDataInterface: &armmobilenetwork.InterfaceProperties{
					Name: to.Ptr("N6"),
				},
			},
		},
		nil)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmobilenetwork%2Farmmobilenetwork%2Fv0.5.0/sdk/resourcemanager/mobilenetwork/armmobilenetwork/README.md) on how to add the SDK to your project and authenticate.
