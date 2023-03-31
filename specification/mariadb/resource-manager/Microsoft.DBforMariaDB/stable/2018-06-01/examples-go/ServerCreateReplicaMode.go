package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerCreateReplicaMode.json
func ExampleServersClient_BeginCreate_createAReplicaServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "TargetResourceGroup", "targetserver", armmariadb.ServerForCreate{
		Location: to.Ptr("westus"),
		Properties: &armmariadb.ServerPropertiesForReplica{
			CreateMode:     to.Ptr(armmariadb.CreateModeReplica),
			SourceServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMariaDB/servers/masterserver"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armmariadb.Server{
	// 	Name: to.Ptr("targetserver"),
	// 	Type: to.Ptr("Microsoft.DBforMariaDB/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TargetResourceGroup/providers/Microsoft.DBforMariaDB/servers/targetserver"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmariadb.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577+00:00"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("targetserver.mariadb.database.azure.com"),
	// 		MasterServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/MasterResourceGroup/providers/Microsoft.DBforMariaDB/servers/masterserver"),
	// 		ReplicaCapacity: to.Ptr[int32](0),
	// 		ReplicationRole: to.Ptr("Replica"),
	// 		SSLEnforcement: to.Ptr(armmariadb.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmariadb.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](14),
	// 			GeoRedundantBackup: to.Ptr(armmariadb.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmariadb.ServerStateReady),
	// 		Version: to.Ptr(armmariadb.ServerVersionTen3),
	// 	},
	// 	SKU: &armmariadb.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmariadb.SKUTierGeneralPurpose),
	// 	},
	// }
}
