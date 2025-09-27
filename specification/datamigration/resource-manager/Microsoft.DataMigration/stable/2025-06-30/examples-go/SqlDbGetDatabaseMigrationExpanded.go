package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/SqlDbGetDatabaseMigrationExpanded.json
func ExampleDatabaseMigrationsSQLDbClient_Get_getSqlDbDatabaseMigrationWithTheExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseMigrationsSQLDbClient().Get(ctx, "testrg", "sqldbinstance", "db1", &armdatamigration.DatabaseMigrationsSQLDbClientGetOptions{MigrationOperationID: nil,
		Expand: to.Ptr("MigrationStatusDetails"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseMigrationSQLDb = armdatamigration.DatabaseMigrationSQLDb{
	// 	Name: to.Ptr("db1"),
	// 	Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance/providers/Microsoft.DataMigration/databaseMigrations/db1"),
	// 	Properties: &armdatamigration.DatabaseMigrationPropertiesSQLDb{
	// 		EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		Kind: to.Ptr(armdatamigration.ResourceTypeSQLDb),
	// 		MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
	// 		MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
	// 		MigrationStatus: to.Ptr("Succeeded"),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 		Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance"),
	// 		StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		SourceDatabaseName: to.Ptr("sourcename"),
	// 		SourceServerName: to.Ptr("sourceserver"),
	// 		MigrationStatusDetails: &armdatamigration.SQLDbMigrationStatusDetails{
	// 			ListOfCopyProgressDetails: []*armdatamigration.CopyProgressDetails{
	// 				{
	// 					CopyDuration: to.Ptr[int32](90),
	// 					CopyStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-19T00:00:00.000Z"); return t}()),
	// 					CopyThroughput: to.Ptr[float64](100),
	// 					DataRead: to.Ptr[int64](100),
	// 					DataWritten: to.Ptr[int64](100),
	// 					ParallelCopyType: to.Ptr("None"),
	// 					RowsCopied: to.Ptr[int64](10),
	// 					RowsRead: to.Ptr[int64](10),
	// 					Status: to.Ptr("Succeeded"),
	// 					TableName: to.Ptr("Table.Name"),
	// 					UsedParallelCopies: to.Ptr[int32](1),
	// 				},
	// 				{
	// 					CopyDuration: to.Ptr[int32](0),
	// 					CopyStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-19T04:08:00.000Z"); return t}()),
	// 					CopyThroughput: to.Ptr[float64](0),
	// 					DataRead: to.Ptr[int64](100),
	// 					DataWritten: to.Ptr[int64](0),
	// 					ParallelCopyType: to.Ptr("DynamicRange"),
	// 					RowsCopied: to.Ptr[int64](0),
	// 					RowsRead: to.Ptr[int64](10),
	// 					Status: to.Ptr("InProgress"),
	// 					TableName: to.Ptr("Table2.Name"),
	// 					UsedParallelCopies: to.Ptr[int32](4),
	// 			}},
	// 			MigrationState: to.Ptr("MonitorMigration"),
	// 			SQLDataCopyErrors: []*string{
	// 				to.Ptr("")},
	// 			},
	// 			OfflineConfiguration: &armdatamigration.SQLDbOfflineConfiguration{
	// 				Offline: to.Ptr(true),
	// 			},
	// 		},
	// 	}
}
