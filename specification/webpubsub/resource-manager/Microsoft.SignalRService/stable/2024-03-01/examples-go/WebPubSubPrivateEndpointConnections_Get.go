package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7189fb57f69468c56df76f9a4d68dd9ff04ab100/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/WebPubSubPrivateEndpointConnections_Get.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e", "myResourceGroup", "myWebPubSubService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armwebpubsub.PrivateEndpointConnection{
	// 	Name: to.Ptr("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/privateEndpointConnections/mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 	SystemData: &armwebpubsub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armwebpubsub.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("webpubsub")},
	// 			PrivateEndpoint: &armwebpubsub.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armwebpubsub.PrivateLinkServiceConnectionState{
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armwebpubsub.PrivateLinkServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
