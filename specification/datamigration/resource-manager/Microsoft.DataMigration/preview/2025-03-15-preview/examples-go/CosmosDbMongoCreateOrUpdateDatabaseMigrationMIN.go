package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/CosmosDbMongoCreateOrUpdateDatabaseMigrationMIN.json
func ExampleDatabaseMigrationsMongoToCosmosDbvCoreMongoClient_BeginCreate_createMongoToCosmosDbMongoVCoreDatabaseMigrationResourceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseMigrationsMongoToCosmosDbvCoreMongoClient().BeginCreate(ctx, "testrg", "targetCosmosDbClusterName", "migrationRequest", armdatamigration.DatabaseMigrationCosmosDbMongo{
		Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
			Kind:             to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
			MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/MigrationServices/testMigrationService"),
			Scope:            to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
			CollectionList: []*armdatamigration.MongoMigrationCollection{
				{
					SourceCollection: to.Ptr("sourceCol1"),
					SourceDatabase:   to.Ptr("sourceDb1"),
					TargetCollection: to.Ptr("targetCol1"),
					TargetDatabase:   to.Ptr("targetDb1"),
				},
				{
					SourceCollection: to.Ptr("sourceCol2"),
					SourceDatabase:   to.Ptr("sourceDb2"),
				}},
			SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
				Host:     to.Ptr("abc.mongodb.com"),
				Password: to.Ptr("placeholder"),
				Port:     to.Ptr[int32](88),
				UseSSL:   to.Ptr(true),
				UserName: to.Ptr("abc"),
			},
			TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
				ConnectionString: to.Ptr("placeholder"),
			},
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
	// res.DatabaseMigrationCosmosDbMongo = armdatamigration.DatabaseMigrationCosmosDbMongo{
	// 	Name: to.Ptr("migrationRequest"),
	// 	Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName/providers/Microsoft.DataMigration/databaseMigrations/migrationRequest"),
	// 	Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
	// 		Kind: to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
	// 		MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/testMigrationService"),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 		Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
	// 		CollectionList: []*armdatamigration.MongoMigrationCollection{
	// 			{
	// 				SourceCollection: to.Ptr("sourceCol1"),
	// 				SourceDatabase: to.Ptr("sourceDb1"),
	// 				TargetCollection: to.Ptr("targetCol1"),
	// 				TargetDatabase: to.Ptr("targetDb1"),
	// 			},
	// 			{
	// 				SourceCollection: to.Ptr("sourceCol2"),
	// 				SourceDatabase: to.Ptr("sourceDb2"),
	// 				TargetCollection: to.Ptr("sourceCol2"),
	// 				TargetDatabase: to.Ptr("sourceDb2"),
	// 		}},
	// 		SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
	// 			Host: to.Ptr("abc.mongodb.com"),
	// 			Port: to.Ptr[int32](88),
	// 		},
	// 		TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
	// 			Host: to.Ptr("xyz.mongo.cosmos.azure.com"),
	// 			Port: to.Ptr[int32](10255),
	// 		},
	// 	},
	// }
}
