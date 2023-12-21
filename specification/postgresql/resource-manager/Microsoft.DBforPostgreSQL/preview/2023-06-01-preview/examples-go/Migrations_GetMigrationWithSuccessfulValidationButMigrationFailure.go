package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_GetMigrationWithSuccessfulValidationButMigrationFailure.json
func ExampleMigrationsClient_Get_migrationsGetMigrationWithSuccessfulValidationButMigrationFailure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Get(ctx, "ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", "testmigrationwithsuccessfulvalidationbutmigrationfailure", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationResource = armpostgresqlflexibleservers.MigrationResource{
	// 	Name: to.Ptr("testmigrationwithsuccessfulvalidationbutmigrationfailure"),
	// 	Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget/migrations/testmigrationwithsuccessfulvalidationbutmigrationfailure"),
	// 	Location: to.Ptr("eastus2"),
	// 	Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
	// 		CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 			CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubStateDetails{
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubStateCompleted),
	// 				DbDetails: map[string]*armpostgresqlflexibleservers.DbMigrationStatus{
	// 					"testdb6": &armpostgresqlflexibleservers.DbMigrationStatus{
	// 						AppliedChanges: to.Ptr[int32](0),
	// 						CdcDeleteCounter: to.Ptr[int32](0),
	// 						CdcInsertCounter: to.Ptr[int32](0),
	// 						CdcUpdateCounter: to.Ptr[int32](0),
	// 						DatabaseName: to.Ptr("testdb6"),
	// 						EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:09:53.252Z"); return t}()),
	// 						FullLoadCompletedTables: to.Ptr[int32](0),
	// 						FullLoadErroredTables: to.Ptr[int32](0),
	// 						FullLoadLoadingTables: to.Ptr[int32](0),
	// 						FullLoadQueuedTables: to.Ptr[int32](0),
	// 						IncomingChanges: to.Ptr[int32](0),
	// 						Latency: to.Ptr[int32](0),
	// 						Message: to.Ptr("Collation/Encoding not Supported Error:  User defined Collations are present in the source database. Please drop them before retrying the migration."),
	// 						MigrationState: to.Ptr(armpostgresqlflexibleservers.MigrationDbStateFailed),
	// 						StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:09:52.922Z"); return t}()),
	// 					},
	// 				},
	// 				ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 					DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 						{
	// 							DatabaseName: to.Ptr("address_standardizer"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:04:24.565Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:04:24.565Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							}},
	// 					}},
	// 					ServerLevelValidationDetails: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 						{
	// 							Type: to.Ptr("AuthenticationAndConnectivityValidation"),
	// 							State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 					}},
	// 					Status: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 					ValidationEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:04:24.565Z"); return t}()),
	// 					ValidationStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:04:21.623Z"); return t}()),
	// 				},
	// 			},
	// 			Error: to.Ptr("testdb6: Collation/Encoding not Supported Error:  User defined Collations are present in the source database. Please drop them before retrying the migration. "),
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateFailed),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("testdb6")},
	// 			MigrationID: to.Ptr("da52db29-cfeb-4670-a1ad-683edb14c621"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidateAndMigrate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:04:24.565Z"); return t}()),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-21T13:04:24.565Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDbsInTargetEnumTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceDbEnumTrue),
	// 			SourceDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
	// 				Location: to.Ptr("eastus2"),
	// 				SKU: &armpostgresqlflexibleservers.ServerSKU{
	// 				},
	// 				StorageMb: to.Ptr[int32](102400),
	// 			},
	// 			SourceDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/source-rg/providers/Microsoft.DBforPostgreSQL/servers/testsource"),
	// 			TargetDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
	// 				Location: to.Ptr("East US 2"),
	// 				SKU: &armpostgresqlflexibleservers.ServerSKU{
	// 					Name: to.Ptr("Standard_D2ds_v4"),
	// 					Tier: to.Ptr(armpostgresqlflexibleservers.SKUTier("Standard_D2ds_v4")),
	// 				},
	// 				StorageMb: to.Ptr[int32](131072),
	// 			},
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testtarget"),
	// 		},
	// 	}
}
