package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/CosmosDbMongoListByScopeDatabaseMigration2.json
func ExampleDatabaseMigrationsMongoToCosmosDbvCoreMongoClient_NewGetForScopePager_getMongoToCosmosDbMongoVCoreDatabaseMigrationWithoutTheExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseMigrationsMongoToCosmosDbvCoreMongoClient().NewGetForScopePager("testrg", "targetCosmosDbClusterName", nil)
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
		// page = armdatamigration.DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse{
		// 	DatabaseMigrationCosmosDbMongoListResult: armdatamigration.DatabaseMigrationCosmosDbMongoListResult{
		// 		Value: []*armdatamigration.DatabaseMigrationCosmosDbMongo{
		// 			{
		// 				Name: to.Ptr("migrationRequest1"),
		// 				Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName/providers/Microsoft.DataMigration/databaseMigrations/migrationRequest1"),
		// 				Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
		// 					Kind: to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
		// 					MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 					MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/testMigrationService"),
		// 					MigrationStatus: to.Ptr("InProgress"),
		// 					ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
		// 					Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
		// 					SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 						Host: to.Ptr("abc.mongodb.com"),
		// 						Port: to.Ptr[int32](88),
		// 					},
		// 					StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
		// 					TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 						Host: to.Ptr("xyz.mongocluster.cosmos.azure.com"),
		// 						Port: to.Ptr[int32](10255),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("migrationRequest2"),
		// 				Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName/providers/Microsoft.DataMigration/databaseMigrations/migrationRequest2"),
		// 				Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
		// 					EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:01:00Z"); return t}()),
		// 					Kind: to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
		// 					MigrationFailureError: &armdatamigration.ErrorInfo{
		// 						Code: to.Ptr("400"),
		// 						Message: to.Ptr("Source or Target database connectivity could not be validated."),
		// 					},
		// 					MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 					MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/testMigrationService"),
		// 					MigrationStatus: to.Ptr("Failed"),
		// 					ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
		// 					Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
		// 					SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 						Host: to.Ptr("abc.mongodb.com"),
		// 						Port: to.Ptr[int32](88),
		// 					},
		// 					StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
		// 					TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 						Host: to.Ptr("xyz.mongocluster.cosmos.azure.com"),
		// 						Port: to.Ptr[int32](10255),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
