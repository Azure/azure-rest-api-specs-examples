package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabaseBackup_listByParent.json
func ExampleAutonomousDatabaseBackupsClient_NewListByAutonomousDatabasePager_autonomousDatabaseBackupsListByAutonomousDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutonomousDatabaseBackupsClient().NewListByAutonomousDatabasePager("rg000", "databasedb1", nil)
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
		// page.AutonomousDatabaseBackupListResult = armoracledatabase.AutonomousDatabaseBackupListResult{
		// 	Value: []*armoracledatabase.AutonomousDatabaseBackup{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/autonomousDatabases/autonomousDatabaseBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1/autonomousDatabaseBackups/1711644130"),
		// 			Properties: &armoracledatabase.AutonomousDatabaseBackupProperties{
		// 				AutonomousDatabaseOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 				BackupType: to.Ptr(armoracledatabase.AutonomousDatabaseBackupTypeFull),
		// 				DatabaseSizeInTbs: to.Ptr[float64](2),
		// 				DbVersion: to.Ptr("19.6.0.0"),
		// 				DisplayName: to.Ptr("Nightly Backup"),
		// 				IsAutomatic: to.Ptr(true),
		// 				IsRestorable: to.Ptr(true),
		// 				LifecycleDetails: to.Ptr("Backup completed successfully"),
		// 				LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseBackupLifecycleStateActive),
		// 				Ocid: to.Ptr("ocid1.autonomousdatabasebackup.oc1..aaaaaaaavwpj"),
		// 				ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 				RetentionPeriodInDays: to.Ptr[int32](365),
		// 				SizeInTbs: to.Ptr[float64](2),
		// 				TimeAvailableTil: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-09T20:44:09.466Z"); return t}()),
		// 				TimeEnded: to.Ptr("2024-01-09T20:44:09.466Z"),
		// 			},
		// 	}},
		// }
	}
}
