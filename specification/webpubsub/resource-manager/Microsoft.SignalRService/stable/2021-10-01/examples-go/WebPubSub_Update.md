Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwebpubsub%2Farmwebpubsub%2Fv0.4.0/sdk/resourcemanager/webpubsub/armwebpubsub/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_Update.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armwebpubsub.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armwebpubsub.ResourceInfo{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Identity: &armwebpubsub.ManagedIdentity{
				Type: to.Ptr(armwebpubsub.ManagedIdentityTypeSystemAssigned),
			},
			Properties: &armwebpubsub.Properties{
				DisableAADAuth:   to.Ptr(false),
				DisableLocalAuth: to.Ptr(false),
				LiveTraceConfiguration: &armwebpubsub.LiveTraceConfiguration{
					Categories: []*armwebpubsub.LiveTraceCategory{
						{
							Name:    to.Ptr("<name>"),
							Enabled: to.Ptr("<enabled>"),
						}},
					Enabled: to.Ptr("<enabled>"),
				},
				NetworkACLs: &armwebpubsub.NetworkACLs{
					DefaultAction: to.Ptr(armwebpubsub.ACLActionDeny),
					PrivateEndpoints: []*armwebpubsub.PrivateEndpointACL{
						{
							Allow: []*armwebpubsub.WebPubSubRequestType{
								to.Ptr(armwebpubsub.WebPubSubRequestTypeServerConnection)},
							Name: to.Ptr("<name>"),
						}},
					PublicNetwork: &armwebpubsub.NetworkACL{
						Allow: []*armwebpubsub.WebPubSubRequestType{
							to.Ptr(armwebpubsub.WebPubSubRequestTypeClientConnection)},
					},
				},
				PublicNetworkAccess: to.Ptr("<public-network-access>"),
				TLS: &armwebpubsub.TLSSettings{
					ClientCertEnabled: to.Ptr(false),
				},
			},
			SKU: &armwebpubsub.ResourceSKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(armwebpubsub.WebPubSubSKUTierStandard),
			},
		},
		&armwebpubsub.ClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
