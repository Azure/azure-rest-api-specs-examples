package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalRPrivateEndpointConnections_Update.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Update(ctx, "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e", "myResourceGroup", "mySignalRService", armsignalr.PrivateEndpointConnection{
		Properties: &armsignalr.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armsignalr.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
			},
			PrivateLinkServiceConnectionState: &armsignalr.PrivateLinkServiceConnectionState{
				ActionsRequired: to.Ptr("None"),
				Status:          to.Ptr(armsignalr.PrivateLinkServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armsignalr.PrivateEndpointConnection{
	// 	Name: to.Ptr("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 	Type: to.Ptr("Microsoft.SignalRService/SignalR/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/privateEndpointConnections/mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
	// 	Properties: &armsignalr.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("signalr")},
	// 			PrivateEndpoint: &armsignalr.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armsignalr.PrivateLinkServiceConnectionState{
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armsignalr.PrivateLinkServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armsignalr.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armsignalr.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armsignalr.CreatedByTypeUser),
	// 		},
	// 	}
}
