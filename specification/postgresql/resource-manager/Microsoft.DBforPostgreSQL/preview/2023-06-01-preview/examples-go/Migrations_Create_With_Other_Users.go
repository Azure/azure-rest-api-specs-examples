package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_Create_With_Other_Users.json
func ExampleMigrationsClient_Create_migrationsCreateByPassingUserNames() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Create(ctx, "ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", "testmigration", armpostgresqlflexibleservers.MigrationResource{
		Location: to.Ptr("westus"),
		Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
			DbsToMigrate: []*string{
				to.Ptr("db1"),
				to.Ptr("db2"),
				to.Ptr("db3"),
				to.Ptr("db4")},
			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
			SecretParameters: &armpostgresqlflexibleservers.MigrationSecretParameters{
				AdminCredentials: &armpostgresqlflexibleservers.AdminCredentials{
					SourceServerPassword: to.Ptr("xxxxxxxx"),
					TargetServerPassword: to.Ptr("xxxxxxxx"),
				},
				SourceServerUsername: to.Ptr("newadmin@testsource"),
				TargetServerUsername: to.Ptr("targetadmin"),
			},
			SourceDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationResource = armpostgresqlflexibleservers.MigrationResource{
	// 	Name: to.Ptr("testmigration"),
	// 	Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget/migrations/testmigration"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key1624": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
	// 		CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 			CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubStateDetails{
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubStatePerformingPreRequisiteSteps),
	// 			},
	// 			Error: to.Ptr(""),
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateInProgress),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("db1"),
	// 			to.Ptr("db2"),
	// 			to.Ptr("db3"),
	// 			to.Ptr("db4")},
	// 			MigrationID: to.Ptr("d3ceacbb-a5fd-43dc-a9db-6022b5154856"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-06T16:05:58.895Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDbsInTargetEnumFalse),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceDbEnumFalse),
	// 			SourceDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource"),
	// 			StartDataMigration: to.Ptr(armpostgresqlflexibleservers.StartDataMigrationEnumFalse),
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverEnumFalse),
	// 		},
	// 	}
}
