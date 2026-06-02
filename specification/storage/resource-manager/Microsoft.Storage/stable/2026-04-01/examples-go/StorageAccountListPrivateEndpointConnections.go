package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/StorageAccountListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("res6977", "sto2527", nil)
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
		// page = armstorage.PrivateEndpointConnectionsClientListResponse{
		// 	PrivateEndpointConnectionListResult: armstorage.PrivateEndpointConnectionListResult{
		// 		Value: []*armstorage.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("{privateEndpointConnectionName}"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Storage/storageAccounts/sto288/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 				Properties: &armstorage.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armstorage.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armstorage.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						ActionRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armstorage.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{privateEndpointConnectionName}"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Storage/storageAccounts/sto288/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 				Properties: &armstorage.PrivateEndpointConnectionProperties{
		// 					PrivateEndpoint: &armstorage.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest02"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armstorage.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						ActionRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armstorage.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
