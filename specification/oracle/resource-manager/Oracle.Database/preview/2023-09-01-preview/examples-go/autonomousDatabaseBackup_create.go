package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabaseBackup_create.json
func ExampleAutonomousDatabaseBackupsClient_BeginCreateOrUpdate_autonomousDatabaseBackupsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabaseBackupsClient().BeginCreateOrUpdate(ctx, "rg000", "databasedb1", "1711644130", armoracledatabase.AutonomousDatabaseBackup{
		Properties: &armoracledatabase.AutonomousDatabaseBackupProperties{
			AutonomousDatabaseOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
			DisplayName:            to.Ptr("Nightly Backup"),
			RetentionPeriodInDays:  to.Ptr[int32](365),
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
	// res.AutonomousDatabaseBackup = armoracledatabase.AutonomousDatabaseBackup{
	// 	Type: to.Ptr("Oracle.Database/autonomousDatabases/autonomousDatabaseBackups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1/autonomousDatabaseBackups/1711644130"),
	// 	Properties: &armoracledatabase.AutonomousDatabaseBackupProperties{
	// 		AutonomousDatabaseOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 		BackupType: to.Ptr(armoracledatabase.AutonomousDatabaseBackupTypeFull),
	// 		DatabaseSizeInTbs: to.Ptr[float64](2),
	// 		DbVersion: to.Ptr("19.6.0.0"),
	// 		DisplayName: to.Ptr("Nightly Backup"),
	// 		IsAutomatic: to.Ptr(true),
	// 		IsRestorable: to.Ptr(true),
	// 		LifecycleDetails: to.Ptr("Backup completed successfully"),
	// 		LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseBackupLifecycleStateActive),
	// 		Ocid: to.Ptr("ocid1.autonomousdatabasebackup.oc1..aaaaaaaavwpj"),
	// 		ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 		RetentionPeriodInDays: to.Ptr[int32](365),
	// 		SizeInTbs: to.Ptr[float64](2),
	// 		TimeAvailableTil: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-09T20:44:09.466Z"); return t}()),
	// 		TimeEnded: to.Ptr("2024-01-09T20:44:09.466Z"),
	// 		TimeStarted: to.Ptr("2024-01-09T19:44:09.466Z"),
	// 	},
	// }
}
