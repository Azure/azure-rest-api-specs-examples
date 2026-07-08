package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-03-01-preview/PrivateEndpointConnectionCreateOrUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myRegistry", "myConnection", armcontainerregistry.PrivateEndpointConnection{
		Properties: &armcontainerregistry.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcontainerregistry.PrivateLinkServiceConnectionState{
				Status:      to.Ptr(armcontainerregistry.ConnectionStatusApproved),
				Description: to.Ptr("Auto-Approved"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.PrivateEndpointConnectionsClientCreateOrUpdateResponse{
	// 	PrivateEndpointConnection: armcontainerregistry.PrivateEndpointConnection{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/privateEndpointConnections/myConnection"),
	// 		Name: to.Ptr("myConnection"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/privateEndpointConnections"),
	// 		Properties: &armcontainerregistry.PrivateEndpointConnectionProperties{
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			PrivateEndpoint: &armcontainerregistry.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armcontainerregistry.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armcontainerregistry.ConnectionStatusApproved),
	// 				Description: to.Ptr("Auto-Approved"),
	// 				ActionsRequired: to.Ptr(armcontainerregistry.ActionsRequiredNone),
	// 			},
	// 		},
	// 	},
	// }
}
