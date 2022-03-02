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

// x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimCreate.json
func ExampleSimsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmobilenetwork.NewSimsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sim-name>",
		armmobilenetwork.Sim{
			Location: to.StringPtr("<location>"),
			Properties: &armmobilenetwork.SimPropertiesFormat{
				AuthenticationKey:                     to.StringPtr("<authentication-key>"),
				DeviceType:                            to.StringPtr("<device-type>"),
				IntegratedCircuitCardIdentifier:       to.StringPtr("<integrated-circuit-card-identifier>"),
				InternationalMobileSubscriberIdentity: to.StringPtr("<international-mobile-subscriber-identity>"),
				MobileNetwork: &armmobilenetwork.ResourceID{
					ID: to.StringPtr("<id>"),
				},
				OperatorKeyCode: to.StringPtr("<operator-key-code>"),
				SimPolicy: &armmobilenetwork.SimPolicyResourceID{
					ID: to.StringPtr("<id>"),
				},
				StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
					{
						AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
							ID: to.StringPtr("<id>"),
						},
						Slice: &armmobilenetwork.SliceResourceID{
							ID: to.StringPtr("<id>"),
						},
						StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
							IPv4Address: to.StringPtr("<ipv4address>"),
						},
					}},
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
	log.Printf("Response result: %#v\n", res.SimsClientCreateOrUpdateResult)
}
```
