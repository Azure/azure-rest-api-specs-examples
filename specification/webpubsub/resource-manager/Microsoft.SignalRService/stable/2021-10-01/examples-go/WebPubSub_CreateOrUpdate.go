package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_CreateOrUpdate.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armwebpubsub.NewClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myWebPubSubService",
		armwebpubsub.ResourceInfo{
			Location: to.Ptr("eastus"),
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
							Name:    to.Ptr("ConnectivityLogs"),
							Enabled: to.Ptr("true"),
						}},
					Enabled: to.Ptr("false"),
				},
				NetworkACLs: &armwebpubsub.NetworkACLs{
					DefaultAction: to.Ptr(armwebpubsub.ACLActionDeny),
					PrivateEndpoints: []*armwebpubsub.PrivateEndpointACL{
						{
							Allow: []*armwebpubsub.WebPubSubRequestType{
								to.Ptr(armwebpubsub.WebPubSubRequestTypeServerConnection)},
							Name: to.Ptr("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
						}},
					PublicNetwork: &armwebpubsub.NetworkACL{
						Allow: []*armwebpubsub.WebPubSubRequestType{
							to.Ptr(armwebpubsub.WebPubSubRequestTypeClientConnection)},
					},
				},
				PublicNetworkAccess: to.Ptr("Enabled"),
				TLS: &armwebpubsub.TLSSettings{
					ClientCertEnabled: to.Ptr(false),
				},
			},
			SKU: &armwebpubsub.ResourceSKU{
				Name:     to.Ptr("Standard_S1"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr(armwebpubsub.WebPubSubSKUTierStandard),
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
