package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerCreateWithBYOK.json
func ExampleServersClient_BeginCreate_createAServerWithByok() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "testrg", "mysqltestserver", armmysqlflexibleservers.Server{
		Location: to.Ptr("southeastasia"),
		Tags: map[string]*string{
			"num": to.Ptr("1"),
		},
		Identity: &armmysqlflexibleservers.Identity{
			Type: to.Ptr("UserAssigned"),
			UserAssignedIdentities: map[string]any{
				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity": map[string]any{},
			},
		},
		Properties: &armmysqlflexibleservers.ServerProperties{
			AdministratorLogin:         to.Ptr("cloudsa"),
			AdministratorLoginPassword: to.Ptr("your_password"),
			AvailabilityZone:           to.Ptr("1"),
			Backup: &armmysqlflexibleservers.Backup{
				BackupRetentionDays: to.Ptr[int32](7),
				GeoRedundantBackup:  to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
			},
			CreateMode: to.Ptr(armmysqlflexibleservers.CreateModeDefault),
			DataEncryption: &armmysqlflexibleservers.DataEncryption{
				Type:                            to.Ptr(armmysqlflexibleservers.DataEncryptionTypeAzureKeyVault),
				GeoBackupKeyURI:                 to.Ptr("https://test-geo.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a"),
				GeoBackupUserAssignedIdentityID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-geo-identity"),
				PrimaryKeyURI:                   to.Ptr("https://test.vault.azure.net/keys/key/c8a92236622244c0a4fdb892666f671a"),
				PrimaryUserAssignedIdentityID:   to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity"),
			},
			HighAvailability: &armmysqlflexibleservers.HighAvailability{
				Mode:                    to.Ptr(armmysqlflexibleservers.HighAvailabilityModeZoneRedundant),
				StandbyAvailabilityZone: to.Ptr("3"),
			},
			Storage: &armmysqlflexibleservers.Storage{
				AutoGrow:      to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
				Iops:          to.Ptr[int32](600),
				StorageSizeGB: to.Ptr[int32](100),
			},
			Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
		},
		SKU: &armmysqlflexibleservers.SKU{
			Name: to.Ptr("Standard_D2ds_v4"),
			Tier: to.Ptr(armmysqlflexibleservers.SKUTierGeneralPurpose),
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
	// 	Name: to.Ptr("mysqltestserver"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver"),
	// 	Location: to.Ptr("Southeast Asia"),
	// 	Tags: map[string]*string{
	// 		"num": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysqlflexibleservers.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		AvailabilityZone: to.Ptr("1"),
	// 		Backup: &armmysqlflexibleservers.Backup{
	// 			BackupRetentionDays: to.Ptr[int32](7),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-17T06:11:38.415Z"); return t}()),
	// 			GeoRedundantBackup: to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
	// 		},
	// 		FullyQualifiedDomainName: to.Ptr("mysqltestserver.database.mysql.azure.com"),
	// 		HighAvailability: &armmysqlflexibleservers.HighAvailability{
	// 			Mode: to.Ptr(armmysqlflexibleservers.HighAvailabilityModeZoneRedundant),
	// 			StandbyAvailabilityZone: to.Ptr("3"),
	// 			State: to.Ptr(armmysqlflexibleservers.HighAvailabilityStateHealthy),
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
	// 			Iops: to.Ptr[int32](600),
	// 			StorageSizeGB: to.Ptr[int32](100),
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
