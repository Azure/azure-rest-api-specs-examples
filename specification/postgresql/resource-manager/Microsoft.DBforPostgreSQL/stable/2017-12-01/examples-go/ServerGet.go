package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerGet.json
func ExampleServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "testrg", "pgtestsvc1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armpostgresql.Server{
	// 	Name: to.Ptr("pgtestsvc1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/servers/pgtestsvc1"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armpostgresql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("testuser"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-04T21:00:58.924Z"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("pgtestsvc1.postgres.database.azure.com"),
	// 		MasterServerID: to.Ptr(""),
	// 		MinimalTLSVersion: to.Ptr(armpostgresql.MinimalTLSVersionEnumTLS12),
	// 		PrivateEndpointConnections: []*armpostgresql.ServerPrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/servers/pgtestsvc1/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
	// 				Properties: &armpostgresql.ServerPrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armpostgresql.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armpostgresql.ServerPrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr(armpostgresql.PrivateLinkServiceConnectionStateActionsRequireNone),
	// 						Status: to.Ptr(armpostgresql.PrivateLinkServiceConnectionStateStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armpostgresql.PrivateEndpointProvisioningState("Succeeded")),
	// 				},
	// 		}},
	// 		PublicNetworkAccess: to.Ptr(armpostgresql.PublicNetworkAccessEnumEnabled),
	// 		ReplicationRole: to.Ptr(""),
	// 		SSLEnforcement: to.Ptr(armpostgresql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armpostgresql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](10),
	// 			GeoRedundantBackup: to.Ptr(armpostgresql.GeoRedundantBackupDisabled),
	// 			StorageMB: to.Ptr[int32](5120),
	// 		},
	// 		UserVisibleState: to.Ptr(armpostgresql.ServerStateReady),
	// 		Version: to.Ptr(armpostgresql.ServerVersionNine5),
	// 	},
	// 	SKU: &armpostgresql.SKU{
	// 		Name: to.Ptr("B_Gen4_1"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("Gen4"),
	// 		Tier: to.Ptr(armpostgresql.SKUTierBasic),
	// 	},
	// }
}
