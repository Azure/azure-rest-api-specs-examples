package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2024-10-01-preview/examples/ServerCreateWithPointInTimeRestore.json
func ExampleServersClient_BeginCreate_createAServerAsAPointInTimeRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "TargetResourceGroup", "targetserver", armmysqlflexibleservers.Server{
		Location: to.Ptr("SoutheastAsia"),
		Tags: map[string]*string{
			"num": to.Ptr("1"),
		},
		Properties: &armmysqlflexibleservers.ServerProperties{
			CreateMode:             to.Ptr(armmysqlflexibleservers.CreateModePointInTimeRestore),
			RestorePointInTime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-24T00:00:37.467Z"); return t }()),
			SourceServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver"),
		},
		SKU: &armmysqlflexibleservers.MySQLServerSKU{
			Name: to.Ptr("Standard_D14_v2"),
			Tier: to.Ptr(armmysqlflexibleservers.ServerSKUTierGeneralPurpose),
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
	// 	Name: to.Ptr("targetserver"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TargetResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/targetserver"),
	// 	Location: to.Ptr("Southeast Asia"),
	// 	Tags: map[string]*string{
	// 		"num": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("adminuser"),
	// 		AvailabilityZone: to.Ptr("1"),
	// 		Backup: &armmysqlflexibleservers.Backup{
	// 			BackupIntervalHours: to.Ptr[int32](24),
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-24T00:15:24.000Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
	// 		},
	// 		DatabasePort: to.Ptr[int32](3306),
	// 		FullVersion: to.Ptr("5.7.44"),
	// 		FullyQualifiedDomainName: to.Ptr("targetserver.database.mysql.azure.com"),
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
	// 		ReplicaCapacity: to.Ptr[int32](10),
	// 		ReplicationRole: to.Ptr(armmysqlflexibleservers.ReplicationRoleNone),
	// 		State: to.Ptr(armmysqlflexibleservers.ServerStateReady),
	// 		Storage: &armmysqlflexibleservers.Storage{
	// 			AutoGrow: to.Ptr(armmysqlflexibleservers.EnableStatusEnumEnabled),
	// 			Iops: to.Ptr[int32](360),
	// 			StorageRedundancy: to.Ptr(armmysqlflexibleservers.StorageRedundancyEnumLocalRedundancy),
	// 			StorageSizeGB: to.Ptr[int32](20),
	// 			StorageSKU: to.Ptr("Premium_LRS"),
	// 		},
	// 		Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysqlflexibleservers.MySQLServerSKU{
	// 		Name: to.Ptr("Standard_D2ds_v4"),
	// 		Tier: to.Ptr(armmysqlflexibleservers.ServerSKUTierGeneralPurpose),
	// 	},
	// }
}
