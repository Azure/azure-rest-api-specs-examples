package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerCreateReplica.json
func ExampleServersClient_BeginCreate_createAReplicaServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "testgr", "replica-server", armmysqlflexibleservers.Server{
		Location: to.Ptr("SoutheastAsia"),
		Properties: &armmysqlflexibleservers.ServerProperties{
			CreateMode:             to.Ptr(armmysqlflexibleservers.CreateModeReplica),
			SourceServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/source-server"),
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
	// res.Server = armmysqlflexibleservers.Server{
	// 	Name: to.Ptr("replica-server"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/replica-server"),
	// 	Location: to.Ptr("Southeast Asia"),
	// 	Tags: map[string]*string{
	// 		"ElasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		AvailabilityZone: to.Ptr("3"),
	// 		Backup: &armmysqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-24T08:19:18.572Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("replica-server.database.mysql.azure.com"),
	// 		HighAvailability: &armmysqlflexibleservers.HighAvailability{
	// 			Mode: to.Ptr(armmysqlflexibleservers.HighAvailabilityModeDisabled),
	// 			State: to.Ptr(armmysqlflexibleservers.HighAvailabilityStateNotEnabled),
	// 		},
	// 		MaintenanceWindow: &armmysqlflexibleservers.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		Network: &armmysqlflexibleservers.Network{
	// 			PublicNetworkAccess: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
	// 		},
	// 		ReplicaCapacity: to.Ptr[int32](0),
	// 		ReplicationRole: to.Ptr(armmysqlflexibleservers.ReplicationRoleReplica),
	// 		SourceServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/source-server"),
	// 		State: to.Ptr(armmysqlflexibleservers.ServerStateReady),
	// 		Storage: &armmysqlflexibleservers.Storage{
	// 			AutoGrow: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
	// 			Iops: to.Ptr[int32](360),
	// 			StorageSizeGB: to.Ptr[int32](20),
	// 			StorageSKU: to.Ptr("Premium_LRS"),
	// 		},
	// 		Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysqlflexibleservers.SKU{
	// 		Name: to.Ptr("Standard_D2ds_v4"),
	// 		Tier: to.Ptr(armmysqlflexibleservers.SKUTierGeneralPurpose),
	// 	},
	// }
}
