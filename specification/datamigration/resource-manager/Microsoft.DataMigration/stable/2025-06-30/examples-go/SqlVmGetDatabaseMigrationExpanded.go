package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/SqlVmGetDatabaseMigrationExpanded.json
func ExampleDatabaseMigrationsSQLVMClient_Get_getSqlVmDatabaseMigrationWithTheExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseMigrationsSQLVMClient().Get(ctx, "testrg", "testvm", "db1", &armdatamigration.DatabaseMigrationsSQLVMClientGetOptions{MigrationOperationID: nil,
		Expand: to.Ptr("MigrationStatusDetails"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseMigrationSQLVM = armdatamigration.DatabaseMigrationSQLVM{
	// 	Name: to.Ptr("db1"),
	// 	Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm/providers/Microsoft.DataMigration/databaseMigrations/db1"),
	// 	Properties: &armdatamigration.DatabaseMigrationPropertiesSQLVM{
	// 		EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		Kind: to.Ptr(armdatamigration.ResourceTypeSQLVM),
	// 		MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
	// 		MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
	// 		MigrationStatus: to.Ptr("InProgress"),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 		Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm"),
	// 		StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		SourceDatabaseName: to.Ptr("sourcename"),
	// 		SourceServerName: to.Ptr("sourceserver"),
	// 		BackupConfiguration: &armdatamigration.BackupConfiguration{
	// 			SourceLocation: &armdatamigration.SourceLocation{
	// 				FileStorageType: to.Ptr("FileShare"),
	// 			},
	// 		},
	// 		MigrationStatusDetails: &armdatamigration.MigrationStatusDetails{
	// 			ActiveBackupSets: []*armdatamigration.SQLBackupSetInfo{
	// 			},
	// 			IsFullBackupRestored: to.Ptr(false),
	// 			MigrationState: to.Ptr("WaitForFullBackupUploadOperation"),
	// 		},
	// 	},
	// }
}
