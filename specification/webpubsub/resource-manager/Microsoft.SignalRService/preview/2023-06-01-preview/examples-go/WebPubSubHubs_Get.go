package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubHubs_Get.json
func ExampleHubsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHubsClient().Get(ctx, "exampleHub", "myResourceGroup", "myWebPubSubService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Hub = armwebpubsub.Hub{
	// 	Name: to.Ptr("exampleHub"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub/hubs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/hubs/exampleHub"),
	// 	Properties: &armwebpubsub.HubProperties{
	// 		AnonymousConnectPolicy: to.Ptr("allow"),
	// 		EventHandlers: []*armwebpubsub.EventHandler{
	// 			{
	// 				Auth: &armwebpubsub.UpstreamAuthSettings{
	// 					Type: to.Ptr(armwebpubsub.UpstreamAuthTypeManagedIdentity),
	// 					ManagedIdentity: &armwebpubsub.ManagedIdentitySettings{
	// 						Resource: to.Ptr("abc"),
	// 					},
	// 				},
	// 				SystemEvents: []*string{
	// 					to.Ptr("connect"),
	// 					to.Ptr("connected")},
	// 					URLTemplate: to.Ptr("http://host.com"),
	// 					UserEventPattern: to.Ptr("*"),
	// 			}},
	// 			EventListeners: []*armwebpubsub.EventListener{
	// 				{
	// 					Endpoint: &armwebpubsub.EventHubEndpoint{
	// 						Type: to.Ptr(armwebpubsub.EventListenerEndpointDiscriminatorEventHub),
	// 						EventHubName: to.Ptr("eventHubName1"),
	// 						FullyQualifiedNamespace: to.Ptr("example.servicebus.windows.net"),
	// 					},
	// 					Filter: &armwebpubsub.EventNameFilter{
	// 						Type: to.Ptr(armwebpubsub.EventListenerFilterDiscriminatorEventName),
	// 						SystemEvents: []*string{
	// 							to.Ptr("connected"),
	// 							to.Ptr("disconnected")},
	// 							UserEventPattern: to.Ptr("*"),
	// 						},
	// 				}},
	// 			},
	// 		}
}
