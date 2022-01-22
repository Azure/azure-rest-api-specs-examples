Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwebpubsub%2Farmwebpubsub%2Fv0.2.0/sdk/resourcemanager/webpubsub/armwebpubsub/README.md) on how to add the SDK to your project and authenticate.

```go
package armwebpubsub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_CreateOrUpdate.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwebpubsub.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armwebpubsub.ResourceInfo{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Identity: &armwebpubsub.ManagedIdentity{
				Type: armwebpubsub.ManagedIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armwebpubsub.Properties{
				DisableAADAuth:   to.BoolPtr(false),
				DisableLocalAuth: to.BoolPtr(false),
				LiveTraceConfiguration: &armwebpubsub.LiveTraceConfiguration{
					Categories: []*armwebpubsub.LiveTraceCategory{
						{
							Name:    to.StringPtr("<name>"),
							Enabled: to.StringPtr("<enabled>"),
						}},
					Enabled: to.StringPtr("<enabled>"),
				},
				NetworkACLs: &armwebpubsub.NetworkACLs{
					DefaultAction: armwebpubsub.ACLAction("Deny").ToPtr(),
					PrivateEndpoints: []*armwebpubsub.PrivateEndpointACL{
						{
							Allow: []*armwebpubsub.WebPubSubRequestType{
								armwebpubsub.WebPubSubRequestType("ServerConnection").ToPtr()},
							Name: to.StringPtr("<name>"),
						}},
					PublicNetwork: &armwebpubsub.NetworkACL{
						Allow: []*armwebpubsub.WebPubSubRequestType{
							armwebpubsub.WebPubSubRequestType("ClientConnection").ToPtr()},
					},
				},
				PublicNetworkAccess: to.StringPtr("<public-network-access>"),
				TLS: &armwebpubsub.TLSSettings{
					ClientCertEnabled: to.BoolPtr(false),
				},
			},
			SKU: &armwebpubsub.ResourceSKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(1),
				Tier:     armwebpubsub.WebPubSubSKUTier("Standard").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ClientCreateOrUpdateResult)
}
```
