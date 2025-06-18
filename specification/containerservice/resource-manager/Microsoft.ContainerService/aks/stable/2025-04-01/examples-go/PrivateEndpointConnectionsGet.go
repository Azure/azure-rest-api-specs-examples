package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/87bd051c295d94fffa28a4fa6b18f8b4b71c50ec/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-04-01/examples/PrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "rg1", "clustername1", "privateendpointconnection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armcontainerservice.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateendpointconnection1"),
	// 	Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
	// 	Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
	// 			Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
