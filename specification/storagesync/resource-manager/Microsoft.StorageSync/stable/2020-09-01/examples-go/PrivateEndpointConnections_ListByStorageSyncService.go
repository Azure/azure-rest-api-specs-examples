package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/PrivateEndpointConnections_ListByStorageSyncService.json
func ExamplePrivateEndpointConnectionsClient_NewListByStorageSyncServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByStorageSyncServicePager("res6977", "sss2527", nil)
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
		// page.PrivateEndpointConnectionListResult = armstoragesync.PrivateEndpointConnectionListResult{
		// 	Value: []*armstoragesync.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("{privateEndpointConnectionName}"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.StorageSync/storageSyncServices/sss2527/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 			Properties: &armstoragesync.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armstoragesync.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armstoragesync.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					Status: to.Ptr(armstoragesync.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armstoragesync.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{privateEndpointConnectionName}"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.StorageSync/storageSyncServices/sss2527/privateEndpointConnections/{privateEndpointConnectionName}"),
		// 			Properties: &armstoragesync.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armstoragesync.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Network/privateEndpoints/petest02"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armstoragesync.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					Status: to.Ptr(armstoragesync.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armstoragesync.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
