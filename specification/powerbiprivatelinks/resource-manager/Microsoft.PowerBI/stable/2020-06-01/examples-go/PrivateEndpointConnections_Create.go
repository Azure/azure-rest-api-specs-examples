package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_Create.json
func ExamplePrivateEndpointConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbiprivatelinks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Create(ctx, "resourceGroup", "azureResourceName", "myPrivateEndpointName", armpowerbiprivatelinks.PrivateEndpointConnection{
		Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointName"),
			},
			PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
				Description:     to.Ptr(""),
				ActionsRequired: to.Ptr("None"),
				Status:          to.Ptr(armpowerbiprivatelinks.PersistedConnectionStatus("Approved ")),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armpowerbiprivatelinks.PrivateEndpointConnection{
	// 	Name: to.Ptr("myPrivateEndpointName.58ffb8de-89ad-41eb-9f8f-de0a7db9d721"),
	// 	Type: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnections/myPrivateEndpointName.58ffb8de-89ad-41eb-9f8f-de0a7db9d721"),
	// 	Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointName"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
	// 			Description: to.Ptr(""),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armpowerbiprivatelinks.PersistedConnectionStatus("Approved ")),
	// 		},
	// 	},
	// }
}
