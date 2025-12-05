package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/MigrationsUpdate.json
func ExampleMigrationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Update(ctx, "exampleresourcegroup", "exampleserver", "examplemigration", armpostgresqlflexibleservers.MigrationResourceForPatch{
		Properties: &armpostgresqlflexibleservers.MigrationPropertiesForPatch{
			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceServerTrue),
		},
	}, nil)
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
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubstatePerformingPreRequisiteSteps),
	// 			},
	// 			Error: to.Ptr(""),
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateInProgress),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("exampledatabase1"),
	// 			to.Ptr("exampledatabase2"),
	// 			to.Ptr("exampledatabase3"),
	// 			to.Ptr("exampledatabase4")},
	// 			MigrateRoles: to.Ptr(armpostgresqlflexibleservers.MigrateRolesAndPermissionsFalse),
	// 			MigrationID: to.Ptr("d3ceacbb-a5fd-43dc-a9db-6022b5154856"),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDatabasesOnTargetServerFalse),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceServerTrue),
	// 			SourceDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
	// 				Location: to.Ptr("eastus"),
	// 				SKU: &armpostgresqlflexibleservers.ServerSKU{
	// 					Name: to.Ptr("B_Gen5_2"),
	// 					Tier: to.Ptr(armpostgresqlflexibleservers.SKUTier("Basic")),
	// 				},
	// 				StorageMb: to.Ptr[int32](1024),
	// 				Version: to.Ptr("10"),
	// 			},
	// 			SourceDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/servers/examplesource"),
	// 			StartDataMigration: to.Ptr(armpostgresqlflexibleservers.StartDataMigrationFalse),
	// 			TargetDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
	// 				Location: to.Ptr("eastus"),
	// 				SKU: &armpostgresqlflexibleservers.ServerSKU{
	// 					Name: to.Ptr("Standard_D4ds_v5"),
	// 					Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
	// 				},
	// 				StorageMb: to.Ptr[int32](2048),
	// 				Version: to.Ptr("17"),
	// 			},
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/flexibleServers/exampletarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverFalse),
	// 		},
	// 	}
}
