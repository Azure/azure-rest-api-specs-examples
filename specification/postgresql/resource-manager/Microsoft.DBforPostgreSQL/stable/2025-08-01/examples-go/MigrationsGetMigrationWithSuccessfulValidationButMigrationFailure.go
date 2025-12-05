package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/MigrationsGetMigrationWithSuccessfulValidationButMigrationFailure.json
func ExampleMigrationsClient_Get_getInformationAboutAMigrationWithSuccessfulValidationButFailedMigration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Get(ctx, "exampleresourcegroup", "exampleserver", "examplemigration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Migration = armpostgresqlflexibleservers.Migration{
	// 	Name: to.Ptr("examplemigration"),
	// 	Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/flexibleServers/exampletarget/migrations/examplemigration"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpostgresqlflexibleservers.MigrationProperties{
	// 		CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 			CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubstateDetails{
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubstateCompleted),
	// 				DbDetails: map[string]*armpostgresqlflexibleservers.DatabaseMigrationState{
	// 					"exampledatabase": &armpostgresqlflexibleservers.DatabaseMigrationState{
	// 						AppliedChanges: to.Ptr[int32](0),
	// 						CdcDeleteCounter: to.Ptr[int32](0),
	// 						CdcInsertCounter: to.Ptr[int32](0),
	// 						CdcUpdateCounter: to.Ptr[int32](0),
	// 						DatabaseName: to.Ptr("exampledatabase"),
	// 						EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
	// 						FullLoadCompletedTables: to.Ptr[int32](0),
	// 						FullLoadErroredTables: to.Ptr[int32](0),
	// 						FullLoadLoadingTables: to.Ptr[int32](0),
	// 						FullLoadQueuedTables: to.Ptr[int32](0),
	// 						IncomingChanges: to.Ptr[int32](0),
	// 						Latency: to.Ptr[int32](0),
	// 						Message: to.Ptr("Collation/Encoding not Supported Error:  User defined Collations are present in the source database. Please drop them before retrying the migration."),
	// 						MigrationState: to.Ptr(armpostgresqlflexibleservers.MigrationDatabaseStateFailed),
	// 						StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123Z"); return t}()),
	// 					},
	// 				},
	// 				ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 					DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 						{
	// 							DatabaseName: to.Ptr("address_standardizer"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123Z"); return t}()),
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
	// 					ValidationEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
	// 					ValidationStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123Z"); return t}()),
	// 				},
	// 			},
	// 			Error: to.Ptr("exampledatabase: Collation/Encoding not Supported Error:  User defined Collations are present in the source database. Please drop them before retrying the migration."),
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateFailed),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("exampledatabase")},
	// 			MigrateRoles: to.Ptr(armpostgresqlflexibleservers.MigrateRolesAndPermissionsFalse),
	// 			MigrationID: to.Ptr("da52db29-cfeb-4670-a1ad-683edb14c621"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidateAndMigrate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDatabasesOnTargetServerTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceServerTrue),
	// 			SourceDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
	// 				Location: to.Ptr("eastus"),
	// 				SKU: &armpostgresqlflexibleservers.ServerSKU{
	// 				},
	// 				StorageMb: to.Ptr[int32](102400),
	// 			},
	// 			SourceDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/servers/examplesource"),
	// 			TargetDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
	// 				Location: to.Ptr("eastus"),
	// 				SKU: &armpostgresqlflexibleservers.ServerSKU{
	// 					Name: to.Ptr("Standard_D2ds_v4"),
	// 					Tier: to.Ptr(armpostgresqlflexibleservers.SKUTier("Standard_D2ds_v4")),
	// 				},
	// 				StorageMb: to.Ptr[int32](131072),
	// 			},
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampletarget"),
	// 		},
	// 	}
}
