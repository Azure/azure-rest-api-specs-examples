package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewPrivateEndpointConnectionsClient("01234567-89ab-4def-0123-456789abcdef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "ExampleResourceGroup", "ExampleWorkspace", "ExamplePrivateEndpointConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armsynapse.PrivateEndpointConnection{
	// 	Name: to.Ptr("sql"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/privateEndpointConnections/ExamplePrivateEndpointConnection"),
	// 	Properties: &armsynapse.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armsynapse.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/privateEndpoints/ExamplePrivateEndpoint"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armsynapse.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
