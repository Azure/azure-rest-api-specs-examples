package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/AutonomousDatabaseBackups_Update_MaximumSet_Gen.json
func ExampleAutonomousDatabaseBackupsClient_BeginUpdate_patchAutonomousDatabaseBackupGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabaseBackupsClient().BeginUpdate(ctx, "rgopenapi", "databasedb1", "1711644130", armoracledatabase.AutonomousDatabaseBackupUpdate{
		Properties: &armoracledatabase.AutonomousDatabaseBackupUpdateProperties{
			RetentionPeriodInDays: to.Ptr[int32](90),
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
	// res = armoracledatabase.AutonomousDatabaseBackupsClientUpdateResponse{
	// 	AutonomousDatabaseBackup: &armoracledatabase.AutonomousDatabaseBackup{
	// 		Properties: &armoracledatabase.AutonomousDatabaseBackupProperties{
	// 			AutonomousDatabaseOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			DisplayName: to.Ptr("Nightly Backup"),
	// 			RetentionPeriodInDays: to.Ptr[int32](365),
	// 			DatabaseSizeInTbs: to.Ptr[float64](2),
	// 			DbVersion: to.Ptr("19.6.0.0"),
	// 			Ocid: to.Ptr("ocid1.autonomousdatabasebackup.oc1..aaaaaaaavwpj"),
	// 			IsAutomatic: to.Ptr(true),
	// 			IsRestorable: to.Ptr(true),
	// 			LifecycleDetails: to.Ptr("Backup completed successfully"),
	// 			LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseBackupLifecycleStateActive),
	// 			SizeInTbs: to.Ptr[float64](2),
	// 			TimeAvailableTil: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-09T20:44:09.466Z"); return t}()),
	// 			TimeStarted: to.Ptr("lcogiebqmxudyzo"),
	// 			TimeEnded: to.Ptr("2024-01-09T20:44:09.466Z"),
	// 			BackupType: to.Ptr(armoracledatabase.AutonomousDatabaseBackupTypeFull),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1/autonomousDatabaseBackups/1711644130"),
	// 		Name: to.Ptr("urphoddzfgrjwtuyegxelktvcx"),
	// 		Type: to.Ptr("Oracle.Database/autonomousDatabases/autonomousDatabaseBackups"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("sqehacivpuim"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 		},
	// 	},
	// }
}
