package armsignalr_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_Update.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.ResourceInfo{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Identity: &armsignalr.ManagedIdentity{
				Type: armsignalr.ManagedIdentityType("SystemAssigned").ToPtr(),
			},
			Kind: armsignalr.ServiceKind("SignalR").ToPtr(),
			Properties: &armsignalr.Properties{
				Cors: &armsignalr.CorsSettings{
					AllowedOrigins: []*string{
						to.StringPtr("https://foo.com"),
						to.StringPtr("https://bar.com")},
				},
				DisableAADAuth:   to.BoolPtr(false),
				DisableLocalAuth: to.BoolPtr(false),
				Features: []*armsignalr.Feature{
					{
						Flag:       armsignalr.FeatureFlags("ServiceMode").ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlags("EnableConnectivityLogs").ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlags("EnableMessagingLogs").ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlags("EnableLiveTrace").ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					}},
				NetworkACLs: &armsignalr.NetworkACLs{
					DefaultAction: armsignalr.ACLAction("Deny").ToPtr(),
					PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
						{
							Allow: []*armsignalr.SignalRRequestType{
								armsignalr.SignalRRequestType("ServerConnection").ToPtr()},
							Name: to.StringPtr("<name>"),
						}},
					PublicNetwork: &armsignalr.NetworkACL{
						Allow: []*armsignalr.SignalRRequestType{
							armsignalr.SignalRRequestType("ClientConnection").ToPtr()},
					},
				},
				PublicNetworkAccess: to.StringPtr("<public-network-access>"),
				TLS: &armsignalr.TLSSettings{
					ClientCertEnabled: to.BoolPtr(false),
				},
				Upstream: &armsignalr.ServerlessUpstreamSettings{
					Templates: []*armsignalr.UpstreamTemplate{
						{
							Auth: &armsignalr.UpstreamAuthSettings{
								Type: armsignalr.UpstreamAuthType("ManagedIdentity").ToPtr(),
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
				Tier:     armsignalr.SignalRSKUTier("Standard").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ClientUpdateResult)
}
