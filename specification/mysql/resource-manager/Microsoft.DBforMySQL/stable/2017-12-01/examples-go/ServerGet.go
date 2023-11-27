package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerGet.json
func ExampleServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "testrg", "mysqltestsvc4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("mysqltestsvc4"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("mysqltestsvc4.mysql.database.azure.com"),
	// 		MasterServerID: to.Ptr(""),
	// 		PrivateEndpointConnections: []*armmysql.ServerPrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/mysqltestsvc4/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
	// 				Properties: &armmysql.ServerPrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armmysql.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armmysql.ServerPrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr(armmysql.PrivateLinkServiceConnectionStateActionsRequireNone),
	// 						Status: to.Ptr(armmysql.PrivateLinkServiceConnectionStateStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armmysql.PrivateEndpointProvisioningState("Succeeded")),
	// 				},
	// 		}},
	// 		PublicNetworkAccess: to.Ptr(armmysql.PublicNetworkAccessEnumEnabled),
	// 		ReplicaCapacity: to.Ptr[int32](5),
	// 		ReplicationRole: to.Ptr("None"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen4_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}
