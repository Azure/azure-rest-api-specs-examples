package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/Migrations_ListByTargetServer.json
func ExampleMigrationsClient_NewListByTargetServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMigrationsClient().NewListByTargetServerPager("ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", &armpostgresqlflexibleservers.MigrationsClientListByTargetServerOptions{MigrationListFilter: to.Ptr(armpostgresqlflexibleservers.MigrationListFilterAll)})
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
		// page.MigrationResourceListResult = armpostgresqlflexibleservers.MigrationResourceListResult{
		// 	Value: []*armpostgresqlflexibleservers.MigrationResource{
		// 		{
		// 			Name: to.Ptr("testmigration"),
		// 			Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget/migrations/testmigration"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
		// 				CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
		// 					CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubStateDetails{
		// 						CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubStatePerformingPreRequisiteSteps),
		// 					},
		// 					Error: to.Ptr(""),
		// 					State: to.Ptr(armpostgresqlflexibleservers.MigrationStateInProgress),
		// 				},
		// 				DbsToMigrate: []*string{
		// 					to.Ptr("db1"),
		// 					to.Ptr("db2"),
		// 					to.Ptr("db3"),
		// 					to.Ptr("db4")},
		// 					MigrateRoles: to.Ptr(armpostgresqlflexibleservers.MigrateRolesEnumFalse),
		// 					MigrationID: to.Ptr("d3ceacbb-a5fd-43dc-a9db-6022b5154856"),
		// 					MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-06T16:05:58.895Z"); return t}()),
		// 					OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDbsInTargetEnumFalse),
		// 					SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceDbEnumFalse),
		// 					SourceDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
		// 						Location: to.Ptr("westus"),
		// 						SKU: &armpostgresqlflexibleservers.ServerSKU{
		// 							Name: to.Ptr("B_Gen5_2"),
		// 							Tier: to.Ptr(armpostgresqlflexibleservers.SKUTier("Basic")),
		// 						},
		// 						StorageMb: to.Ptr[int32](1024),
		// 						Version: to.Ptr("10"),
		// 					},
		// 					SourceDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource"),
		// 					StartDataMigration: to.Ptr(armpostgresqlflexibleservers.StartDataMigrationEnumFalse),
		// 					TargetDbServerMetadata: &armpostgresqlflexibleservers.DbServerMetadata{
		// 						Location: to.Ptr("westus"),
		// 						SKU: &armpostgresqlflexibleservers.ServerSKU{
		// 							Name: to.Ptr("Standard_D4s_v3"),
		// 							Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
		// 						},
		// 						StorageMb: to.Ptr[int32](2048),
		// 						Version: to.Ptr("11"),
		// 					},
		// 					TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget"),
		// 					TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverEnumFalse),
		// 				},
		// 		}},
		// 	}
	}
}
