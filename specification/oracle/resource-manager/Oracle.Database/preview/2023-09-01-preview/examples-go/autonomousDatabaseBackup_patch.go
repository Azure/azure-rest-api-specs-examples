package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabaseBackup_patch.json
func ExampleAutonomousDatabaseBackupsClient_BeginUpdate() {
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
	// 	Type: to.Ptr("Oracle.Database/autonomousDatabaseBackups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabaseBackups/1711644130"),
	// 	Properties: &armoracledatabase.AutonomousDatabaseBackupProperties{
	// 		Type: to.Ptr(armoracledatabase.AutonomousDatabaseBackupTypeFull),
	// 		AutonomousDatabaseID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 		DatabaseSizeInTBs: to.Ptr[int32](2),
	// 		DbVersion: to.Ptr("19.6.0.0"),
	// 		DisplayName: to.Ptr("Nightly Backup"),
	// 		IsAutomatic: to.Ptr(true),
	// 		IsRestorable: to.Ptr(true),
	// 		LifecycleDetails: to.Ptr("Backup updated successfully"),
	// 		LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseBackupLifecycleStateActive),
	// 		Ocid: to.Ptr("ocid1.autonomousdatabasebackup.oc1..aaaaaaaavwpj"),
	// 		ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 		RetentionPeriodInDays: to.Ptr[int32](400),
	// 		SizeInTBs: to.Ptr[int32](2),
	// 	},
	// }
}
