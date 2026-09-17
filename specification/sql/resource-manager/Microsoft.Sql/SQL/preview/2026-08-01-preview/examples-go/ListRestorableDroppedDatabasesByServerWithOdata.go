package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/ListRestorableDroppedDatabasesByServerWithOdata.json
func ExampleRestorableDroppedDatabasesClient_NewListByServerPager_getsAListOfRestorableDroppedDatabasesWithODataFiltering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableDroppedDatabasesClient().NewListByServerPager("Default-SQL-SouthEastAsia", "testsvr", &armsql.RestorableDroppedDatabasesClientListByServerOptions{
		Top: to.Ptr[int64](25)})
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
		// page = armsql.RestorableDroppedDatabasesClientListByServerResponse{
		// 	RestorableDroppedDatabaseListResult: armsql.RestorableDroppedDatabaseListResult{
		// 		Value: []*armsql.RestorableDroppedDatabase{
		// 			{
		// 				SKU: &armsql.SKU{
		// 					Name: to.Ptr("GP_Gen4_1"),
		// 					Tier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				Location: to.Ptr("southeastasia"),
		// 				Properties: &armsql.RestorableDroppedDatabaseProperties{
		// 					DatabaseName: to.Ptr("testDb1"),
		// 					MaxSizeBytes: to.Ptr[int64](321),
		// 					CreationDate: to.Ptr(time.Date(2015, time.February, 3, 4, 5, 6, 0, time.UTC)),
		// 					DeletionDate: to.Ptr(time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/restorableDroppedDatabases/testDb1,131592384000000000"),
		// 				Name: to.Ptr("testDb1,131592384000000000"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/restorableDroppedDatabases"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://myreferrer.example/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/restorableDroppedDatabases?api-version=2026-08-01-preview&$top=25&$skipToken=eyJuYW1lIjoidGVzdERiMSwxMzE1OTIzODQwMDAwMDAwMDAifQ%3D%3D"),
		// 	},
		// }
	}
}
