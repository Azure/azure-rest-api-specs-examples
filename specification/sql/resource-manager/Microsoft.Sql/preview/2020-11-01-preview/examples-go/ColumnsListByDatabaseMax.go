package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ColumnsListByDatabaseMax.json
func ExampleDatabaseColumnsClient_NewListByDatabasePager_filterDatabaseColumns() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseColumnsClient().NewListByDatabasePager("myRG", "serverName", "myDatabase", &armsql.DatabaseColumnsClientListByDatabaseOptions{Schema: []string{
		"dbo"},
		Table: []string{
			"customer",
			"address"},
		Column: []string{
			"username"},
		OrderBy: []string{
			"schema asc",
			"table",
			"column desc"},
		Skiptoken: nil,
	})
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
		// page.DatabaseColumnListResult = armsql.DatabaseColumnListResult{
		// 	Value: []*armsql.DatabaseColumn{
		// 		{
		// 			Name: to.Ptr("username"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/schemas/tables/columns"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myRG/providers/Microsoft.Sql/servers/serverName/databases/myDatabase/schemas/dbo/tables/customer/columns/username"),
		// 			Properties: &armsql.DatabaseColumnProperties{
		// 				ColumnType: to.Ptr(armsql.ColumnDataTypeNvarchar),
		// 				IsComputed: to.Ptr(false),
		// 				MemoryOptimized: to.Ptr(false),
		// 				TemporalType: to.Ptr(armsql.TableTemporalTypeNonTemporalTable),
		// 			},
		// 	}},
		// }
	}
}
