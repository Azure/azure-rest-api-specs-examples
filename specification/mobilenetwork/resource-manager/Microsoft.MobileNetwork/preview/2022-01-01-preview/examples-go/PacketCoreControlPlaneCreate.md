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

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/PacketCoreControlPlaneCreate.json
func ExamplePacketCoreControlPlanesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewPacketCoreControlPlanesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<packet-core-control-plane-name>",
		armmobilenetwork.PacketCoreControlPlane{
			Location: to.StringPtr("<location>"),
			Properties: &armmobilenetwork.PacketCoreControlPlanePropertiesFormat{
				ControlPlaneAccessInterface: &armmobilenetwork.InterfaceProperties{
					Name: to.StringPtr("<name>"),
				},
				CoreNetworkTechnology: armmobilenetwork.CoreNetworkType("5GC").ToPtr(),
				CustomLocation: &armmobilenetwork.CustomLocationResourceID{
					ID: to.StringPtr("<id>"),
				},
				MobileNetwork: &armmobilenetwork.ResourceID{
					ID: to.StringPtr("<id>"),
				},
				Version: to.StringPtr("<version>"),
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
	log.Printf("Response result: %#v\n", res.PacketCoreControlPlanesClientCreateOrUpdateResult)
}
```
