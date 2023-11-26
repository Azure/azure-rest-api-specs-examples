package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerList.json
func ExampleServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListPager(nil)
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
		// page.ServerListResult = armmariadb.ServerListResult{
		// 	Value: []*armmariadb.Server{
		// 		{
		// 			Name: to.Ptr("mariadbtestsvc1"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmariadb.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:56:02.627Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mariadbtestsvc1.mariadb.database.azure.com"),
		// 				MinimalTLSVersion: to.Ptr(armmariadb.MinimalTLSVersionEnumTLS12),
		// 				PrivateEndpointConnections: []*armmariadb.ServerPrivateEndpointConnection{
		// 				},
		// 				PublicNetworkAccess: to.Ptr(armmariadb.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmariadb.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
		// 				Version: to.Ptr(armmariadb.ServerVersionTen3),
		// 			},
		// 			SKU: &armmariadb.SKU{
		// 				Name: to.Ptr("B_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmariadb.SKUTierBasic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mariadbtstsvc2"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtstsvc2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmariadb.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mariadbtstsvc2.mariadb.database.azure.com"),
		// 				MinimalTLSVersion: to.Ptr(armmariadb.MinimalTLSVersionEnumTLS11),
		// 				PrivateEndpointConnections: []*armmariadb.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtstsvc2/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmariadb.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmariadb.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmariadb.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmariadb.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmariadb.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmariadb.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmariadb.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmariadb.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
		// 				Version: to.Ptr(armmariadb.ServerVersionTen3),
		// 			},
		// 			SKU: &armmariadb.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mariadbtestsvc3"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg1/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc3"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmariadb.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:59:44.847Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mariadbtestsvc3.mariadb.database.azure.com"),
		// 				MinimalTLSVersion: to.Ptr(armmariadb.MinimalTLSVersionEnumTLS10),
		// 				PrivateEndpointConnections: []*armmariadb.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc3/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmariadb.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmariadb.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmariadb.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmariadb.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmariadb.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmariadb.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmariadb.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmariadb.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](102400),
		// 				},
		// 				UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
		// 				Version: to.Ptr(armmariadb.ServerVersionTen3),
		// 			},
		// 			SKU: &armmariadb.SKU{
		// 				Name: to.Ptr("GP_Gen4_4"),
		// 				Capacity: to.Ptr[int32](4),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
