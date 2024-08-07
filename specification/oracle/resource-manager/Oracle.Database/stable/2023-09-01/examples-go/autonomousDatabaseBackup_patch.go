package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1c63635d66ae38cff18045ab416a6572d3e15f6e/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_patch.json
func ExampleAutonomousDatabaseBackupsClient_BeginUpdate_autonomousDatabaseBackupsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabaseBackupsClient().BeginUpdate(ctx, "rg000", "databasedb1", "1711644130", armoracledatabase.AutonomousDatabaseBackupUpdate{}, nil)
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
	// 		LifecycleDetails: to.Ptr("Backup updated successfully"),
	// 		LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseBackupLifecycleStateActive),
	// 		Ocid: to.Ptr("ocid1.autonomousdatabasebackup.oc1..aaaaaaaavwpj"),
	// 		ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 		RetentionPeriodInDays: to.Ptr[int32](400),
	// 		SizeInTbs: to.Ptr[float64](2),
	// 	},
	// }
}
