Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsignalr%2Farmsignalr%2Fv0.1.0/sdk/resourcemanager/signalr/armsignalr/README.md) on how to add the SDK to your project and authenticate.

```go
package armsignalr_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_CreateOrUpdate.json
func ExampleSignalRClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.SignalRResource{
			TrackedResource: armsignalr.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
			},
			Identity: &armsignalr.ManagedIdentity{
				Type: armsignalr.ManagedIdentityTypeSystemAssigned.ToPtr(),
			},
			Kind: armsignalr.ServiceKindSignalR.ToPtr(),
			Properties: &armsignalr.SignalRProperties{
				Cors: &armsignalr.SignalRCorsSettings{
					AllowedOrigins: []*string{
						to.StringPtr("https://foo.com"),
						to.StringPtr("https://bar.com")},
				},
				DisableAADAuth:   to.BoolPtr(false),
				DisableLocalAuth: to.BoolPtr(false),
				Features: []*armsignalr.SignalRFeature{
					{
						Flag:       armsignalr.FeatureFlagsServiceMode.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableConnectivityLogs.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableMessagingLogs.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableLiveTrace.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					}},
				NetworkACLs: &armsignalr.SignalRNetworkACLs{
					DefaultAction: armsignalr.ACLActionDeny.ToPtr(),
					PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
						{
							NetworkACL: armsignalr.NetworkACL{
								Allow: []*armsignalr.SignalRRequestType{
									armsignalr.SignalRRequestTypeServerConnection.ToPtr()},
							},
							Name: to.StringPtr("<name>"),
						}},
					PublicNetwork: &armsignalr.NetworkACL{
						Allow: []*armsignalr.SignalRRequestType{
							armsignalr.SignalRRequestTypeClientConnection.ToPtr()},
					},
				},
				PublicNetworkAccess: to.StringPtr("<public-network-access>"),
				TLS: &armsignalr.SignalRTLSSettings{
					ClientCertEnabled: to.BoolPtr(false),
				},
				Upstream: &armsignalr.ServerlessUpstreamSettings{
					Templates: []*armsignalr.UpstreamTemplate{
						{
							Auth: &armsignalr.UpstreamAuthSettings{
								Type: armsignalr.UpstreamAuthTypeManagedIdentity.ToPtr(),
								ManagedIdentity: &armsignalr.ManagedIdentitySettings{
									Resource: to.StringPtr("<resource>"),
								},
							},
							CategoryPattern: to.StringPtr("<category-pattern>"),
							EventPattern:    to.StringPtr("<event-pattern>"),
							HubPattern:      to.StringPtr("<hub-pattern>"),
							URLTemplate:     to.StringPtr("<urltemplate>"),
						}},
				},
			},
			SKU: &armsignalr.ResourceSKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(1),
				Tier:     armsignalr.SignalRSKUTierStandard.ToPtr(),
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
	log.Printf("SignalRResource.ID: %s\n", *res.ID)
}
```
