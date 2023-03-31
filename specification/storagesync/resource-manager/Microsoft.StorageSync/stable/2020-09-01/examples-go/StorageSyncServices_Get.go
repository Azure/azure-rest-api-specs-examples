package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServices_Get.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armstoragesync.Service{
	// 	Name: to.Ptr("SampleStorageSyncService_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices"),
	// 	ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armstoragesync.ServiceProperties{
	// 		IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
	// 		PrivateEndpointConnections: []*armstoragesync.PrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/{resource}/{resourceName}/privateEndpointConnections/SampleStorageSyncService_1.cd99f12ba6f3483f9292229e4f822258"),
	// 				Properties: &armstoragesync.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armstoragesync.PrivateEndpoint{
	// 						ID: to.Ptr("subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_2/providers/Microsoft.Network/privateEndpoints/testpe01"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armstoragesync.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Please approve my request, thanks."),
	// 						Status: to.Ptr(armstoragesync.PrivateEndpointServiceConnectionStatusPending),
	// 					},
	// 				},
	// 		}},
	// 		StorageSyncServiceStatus: to.Ptr[int32](0),
	// 		StorageSyncServiceUID: to.Ptr("\"2de01144-72da-4d7f-9d0c-e858855114a8\""),
	// 	},
	// }
}
