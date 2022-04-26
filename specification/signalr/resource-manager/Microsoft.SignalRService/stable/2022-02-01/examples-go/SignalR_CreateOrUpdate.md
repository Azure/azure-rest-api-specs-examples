Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsignalr%2Farmsignalr%2Fv0.5.0/sdk/resourcemanager/signalr/armsignalr/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_CreateOrUpdate.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsignalr.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.ResourceInfo{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Identity: &armsignalr.ManagedIdentity{
				Type: to.Ptr(armsignalr.ManagedIdentityTypeSystemAssigned),
			},
			Kind: to.Ptr(armsignalr.ServiceKindSignalR),
			Properties: &armsignalr.Properties{
				Cors: &armsignalr.CorsSettings{
					AllowedOrigins: []*string{
						to.Ptr("https://foo.com"),
						to.Ptr("https://bar.com")},
				},
				DisableAADAuth:   to.Ptr(false),
				DisableLocalAuth: to.Ptr(false),
				Features: []*armsignalr.Feature{
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsServiceMode),
						Properties: map[string]*string{},
						Value:      to.Ptr("<value>"),
					},
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsEnableConnectivityLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("<value>"),
					},
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsEnableMessagingLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("<value>"),
					},
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsEnableLiveTrace),
						Properties: map[string]*string{},
						Value:      to.Ptr("<value>"),
					}},
				LiveTraceConfiguration: &armsignalr.LiveTraceConfiguration{
					Categories: []*armsignalr.LiveTraceCategory{
						{
							Name:    to.Ptr("<name>"),
							Enabled: to.Ptr("<enabled>"),
						}},
					Enabled: to.Ptr("<enabled>"),
				},
				NetworkACLs: &armsignalr.NetworkACLs{
					DefaultAction: to.Ptr(armsignalr.ACLActionDeny),
					PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
						{
							Allow: []*armsignalr.SignalRRequestType{
								to.Ptr(armsignalr.SignalRRequestTypeServerConnection)},
							Name: to.Ptr("<name>"),
						}},
					PublicNetwork: &armsignalr.NetworkACL{
						Allow: []*armsignalr.SignalRRequestType{
							to.Ptr(armsignalr.SignalRRequestTypeClientConnection)},
					},
				},
				PublicNetworkAccess: to.Ptr("<public-network-access>"),
				TLS: &armsignalr.TLSSettings{
					ClientCertEnabled: to.Ptr(false),
				},
				Upstream: &armsignalr.ServerlessUpstreamSettings{
					Templates: []*armsignalr.UpstreamTemplate{
						{
							Auth: &armsignalr.UpstreamAuthSettings{
								Type: to.Ptr(armsignalr.UpstreamAuthTypeManagedIdentity),
								ManagedIdentity: &armsignalr.ManagedIdentitySettings{
									Resource: to.Ptr("<resource>"),
								},
							},
							CategoryPattern: to.Ptr("<category-pattern>"),
							EventPattern:    to.Ptr("<event-pattern>"),
							HubPattern:      to.Ptr("<hub-pattern>"),
							URLTemplate:     to.Ptr("<urltemplate>"),
						}},
				},
			},
			SKU: &armsignalr.ResourceSKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(armsignalr.SignalRSKUTierStandard),
			},
		},
		&armsignalr.ClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
