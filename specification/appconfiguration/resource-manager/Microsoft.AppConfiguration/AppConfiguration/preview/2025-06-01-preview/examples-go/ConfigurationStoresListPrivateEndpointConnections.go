package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/ConfigurationStoresListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListByConfigurationStorePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByConfigurationStorePager("myResourceGroup", "contoso", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armappconfiguration.PrivateEndpointConnectionsClientListByConfigurationStoreResponse{
		// 	PrivateEndpointConnectionListResult: armappconfiguration.PrivateEndpointConnectionListResult{
		// 		Value: []*armappconfiguration.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("myConnection"),
		// 				Type: to.Ptr("Microsoft.AppConfiguration/configurationStores/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso/privateEndpointConnections/myConnection"),
		// 				Properties: &armappconfiguration.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armappconfiguration.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/peexample01"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armappconfiguration.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						ActionsRequired: to.Ptr(armappconfiguration.ActionsRequiredNone),
		// 						Status: to.Ptr(armappconfiguration.ConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
