```go
package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_Update.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsignalr.NewClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"mySignalRService",
		armsignalr.ResourceInfo{
			Location: to.Ptr("eastus"),
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
						Value:      to.Ptr("Serverless"),
					},
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsEnableConnectivityLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("True"),
					},
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsEnableMessagingLogs),
						Properties: map[string]*string{},
						Value:      to.Ptr("False"),
					},
					{
						Flag:       to.Ptr(armsignalr.FeatureFlagsEnableLiveTrace),
						Properties: map[string]*string{},
						Value:      to.Ptr("False"),
					}},
				LiveTraceConfiguration: &armsignalr.LiveTraceConfiguration{
					Categories: []*armsignalr.LiveTraceCategory{
						{
							Name:    to.Ptr("ConnectivityLogs"),
							Enabled: to.Ptr("true"),
						}},
					Enabled: to.Ptr("false"),
				},
				NetworkACLs: &armsignalr.NetworkACLs{
					DefaultAction: to.Ptr(armsignalr.ACLActionDeny),
					PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
						{
							Allow: []*armsignalr.SignalRRequestType{
								to.Ptr(armsignalr.SignalRRequestTypeServerConnection)},
							Name: to.Ptr("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
						}},
					PublicNetwork: &armsignalr.NetworkACL{
						Allow: []*armsignalr.SignalRRequestType{
							to.Ptr(armsignalr.SignalRRequestTypeClientConnection)},
					},
				},
				PublicNetworkAccess: to.Ptr("Enabled"),
				TLS: &armsignalr.TLSSettings{
					ClientCertEnabled: to.Ptr(false),
				},
				Upstream: &armsignalr.ServerlessUpstreamSettings{
					Templates: []*armsignalr.UpstreamTemplate{
						{
							Auth: &armsignalr.UpstreamAuthSettings{
								Type: to.Ptr(armsignalr.UpstreamAuthTypeManagedIdentity),
								ManagedIdentity: &armsignalr.ManagedIdentitySettings{
									Resource: to.Ptr("api://example"),
								},
							},
							CategoryPattern: to.Ptr("*"),
							EventPattern:    to.Ptr("connect,disconnect"),
							HubPattern:      to.Ptr("*"),
							URLTemplate:     to.Ptr("https://example.com/chat/api/connect"),
						}},
				},
			},
			SKU: &armsignalr.ResourceSKU{
				Name:     to.Ptr("Standard_S1"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(armsignalr.SignalRSKUTierStandard),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsignalr%2Farmsignalr%2Fv1.0.0/sdk/resourcemanager/signalr/armsignalr/README.md) on how to add the SDK to your project and authenticate.
