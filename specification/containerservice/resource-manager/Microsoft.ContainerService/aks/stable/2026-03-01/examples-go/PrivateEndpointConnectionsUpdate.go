package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-03-01/PrivateEndpointConnectionsUpdate.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Update(ctx, "rg1", "clustername1", "privateendpointconnection1", armcontainerservice.PrivateEndpointConnection{
		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservice.PrivateEndpointConnectionsClientUpdateResponse{
	// 	PrivateEndpointConnection: armcontainerservice.PrivateEndpointConnection{
	// 		Name: to.Ptr("privateendpointconnection1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/managedClusters/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/privateEndpointConnections/privateendpointconnection1"),
	// 		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
	// 			PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
