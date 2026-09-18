package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/DataWarehouseRestorePointsListByDatabase.json
func ExampleRestorePointsClient_NewListByDatabasePager_listDatawarehouseDatabaseRestorePoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointsClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testserver", "testDatabase", nil)
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
		// page = armsql.RestorePointsClientListByDatabaseResponse{
		// 	RestorePointListResult: armsql.RestorePointListResult{
		// 		Value: []*armsql.RestorePoint{
		// 			{
		// 				Name: to.Ptr("131546477590000000"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131546477590000000"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.RestorePointProperties{
		// 					RestorePointCreationDate: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					RestorePointLabel: to.Ptr("mylabel1"),
		// 					RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("131553636140000000"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131553636140000000"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.RestorePointProperties{
		// 					RestorePointCreationDate: to.Ptr(time.Date(2017, time.November, 17, 3, 40, 14, 0, time.UTC)),
		// 					RestorePointLabel: to.Ptr("mylabel2"),
		// 					RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("131553619750000000"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131553619750000000"),
		// 				Location: to.Ptr("japaneast"),
		// 				Properties: &armsql.RestorePointProperties{
		// 					RestorePointCreationDate: to.Ptr(time.Date(2017, time.November, 17, 3, 12, 55, 0, time.UTC)),
		// 					RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
