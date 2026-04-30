package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2025-08-01-preview/RedisEnterpriseGet.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "rg1", "cache1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredisenterprise.ClientGetResponse{
	// 	Cluster: &armredisenterprise.Cluster{
	// 		Name: to.Ptr("cache1"),
	// 		Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 		ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 		Kind: to.Ptr(armredisenterprise.KindV1),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armredisenterprise.ClusterCreateProperties{
	// 			HostName: to.Ptr("cache1.westus.something.azure.com"),
	// 			MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 			PrivateEndpointConnections: []*armredisenterprise.PrivateEndpointConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/cachePec"),
	// 					Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/cachePe"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Please approve my connection"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armredisenterprise.PublicNetworkAccessDisabled),
	// 			MaintenanceConfiguration: &armredisenterprise.MaintenanceConfiguration{
	// 				MaintenanceWindows: []*armredisenterprise.MaintenanceWindow{
	// 					{
	// 						Type: to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
	// 						Duration: to.Ptr("PT6H"),
	// 						StartHourUTC: to.Ptr[int32](3),
	// 						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
	// 							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekMonday),
	// 						},
	// 					},
	// 					{
	// 						Type: to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
	// 						Duration: to.Ptr("PT6H"),
	// 						StartHourUTC: to.Ptr[int32](3),
	// 						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
	// 							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekTuesday),
	// 						},
	// 					},
	// 					{
	// 						Type: to.Ptr(armredisenterprise.MaintenanceWindowTypeWeekly),
	// 						Duration: to.Ptr("PT6H"),
	// 						StartHourUTC: to.Ptr[int32](3),
	// 						Schedule: &armredisenterprise.MaintenanceWindowSchedule{
	// 							DayOfWeek: to.Ptr(armredisenterprise.MaintenanceDayOfWeekWednesday),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			RedisVersion: to.Ptr("6"),
	// 			ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
	// 		},
	// 		SKU: &armredisenterprise.SKU{
	// 			Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 			Capacity: to.Ptr[int32](3),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 			to.Ptr("3"),
	// 		},
	// 	},
	// }
}
