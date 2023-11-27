package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerList.json
func ExampleServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ServerListResult = armmysql.ServerListResult{
		// 	Value: []*armmysql.Server{
		// 		{
		// 			Name: to.Ptr("mysqltestsvc1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:56:02.627Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestsvc1.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 				},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("B_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierBasic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltstsvc2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltstsvc2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:56:54.300Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltstsvc2.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltstsvc2/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmysql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](7),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupDisabled),
		// 					StorageMB: to.Ptr[int32](5120),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_2"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mysqltestsvc3"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg1/providers/Microsoft.DBforMySQL/servers/mysqltestsvc3"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armmysql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("testuser"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-28T23:59:44.847Z"); return t}()),
		// 				FullyQualifiedDomainName: to.Ptr("mysqltestsvc3.mysql.database.azure.com"),
		// 				PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc3/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armmysql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
		// 				SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
		// 				StorageProfile: &armmysql.StorageProfile{
		// 					BackupRetentionDays: to.Ptr[int32](35),
		// 					GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
		// 					StorageMB: to.Ptr[int32](102400),
		// 				},
		// 				UserVisibleState: to.Ptr(armmysql.ServerStateReady),
		// 				Version: to.Ptr(armmysql.ServerVersionFive7),
		// 			},
		// 			SKU: &armmysql.SKU{
		// 				Name: to.Ptr("GP_Gen4_4"),
		// 				Capacity: to.Ptr[int32](4),
		// 				Family: to.Ptr("Gen4"),
		// 				Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
