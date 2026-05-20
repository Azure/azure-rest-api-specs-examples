package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2022-09-01/StorageSyncServices_Get.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
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
	// res = armstoragesync.ServicesClientGetResponse{
	// 	XMSCorrelationRequestID: to.Ptr("d166ca76-dad2-49df-b409-d2acfd42d730"),
	// 	XMSRequestID: to.Ptr("74e55a4d-1c6f-46de-9a8d-278e53a47403"),
	// 	Service: armstoragesync.Service{
	// 		Name: to.Ptr("SampleStorageSyncService_1"),
	// 		Type: to.Ptr("Microsoft.StorageSync/storageSyncServices"),
	// 		ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1"),
	// 		Identity: &armstoragesync.ManagedServiceIdentity{
	// 			Type: to.Ptr(armstoragesync.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			UserAssignedIdentities: map[string]*armstoragesync.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": &armstoragesync.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("WestUS"),
	// 		Properties: &armstoragesync.ServiceProperties{
	// 			IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
	// 			PrivateEndpointConnections: []*armstoragesync.PrivateEndpointConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_2/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/privateEndpointConnections/SampleStorageSyncService_1.cd99f12ba6f3483f9292229e4f822258"),
	// 					Properties: &armstoragesync.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armstoragesync.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_2/providers/Microsoft.Network/privateEndpoints/testpe01"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armstoragesync.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Please approve my request, thanks."),
	// 							Status: to.Ptr(armstoragesync.PrivateEndpointServiceConnectionStatusPending),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StorageSyncServiceStatus: to.Ptr[int32](0),
	// 			StorageSyncServiceUID: to.Ptr("\"2de01144-72da-4d7f-9d0c-e858855114a8\""),
	// 		},
	// 		SystemData: &armstoragesync.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 			CreatedBy: to.Ptr("sample-user"),
	// 			CreatedByType: to.Ptr(armstoragesync.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("sample-user"),
	// 			LastModifiedByType: to.Ptr(armstoragesync.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
