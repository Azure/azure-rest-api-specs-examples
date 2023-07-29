package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncMemberGetSchema.json
func ExampleSyncMembersClient_NewListMemberSchemasPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncMembersClient().NewListMemberSchemasPager("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "syncgroupcrud-4879", nil)
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
		// page.SyncFullSchemaPropertiesListResult = armsql.SyncFullSchemaPropertiesListResult{
		// 	Value: []*armsql.SyncFullSchemaProperties{
		// 		{
		// 			LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-30T07:16:08.21Z"); return t}()),
		// 			Tables: []*armsql.SyncFullSchemaTable{
		// 				{
		// 					Name: to.Ptr("dbo.myTable"),
		// 					Columns: []*armsql.SyncFullSchemaTableColumn{
		// 						{
		// 							Name: to.Ptr("intField"),
		// 							DataSize: to.Ptr("4"),
		// 							DataType: to.Ptr("int"),
		// 							HasError: to.Ptr(false),
		// 							IsPrimaryKey: to.Ptr(false),
		// 							QuotedName: to.Ptr("[intField]"),
		// 						},
		// 						{
		// 							Name: to.Ptr("charField"),
		// 							DataSize: to.Ptr("100"),
		// 							DataType: to.Ptr("nvarchar"),
		// 							HasError: to.Ptr(false),
		// 							IsPrimaryKey: to.Ptr(false),
		// 							QuotedName: to.Ptr("[charField]"),
		// 					}},
		// 					ErrorID: to.Ptr("Schema_TableHasNoPrimaryKey"),
		// 					HasError: to.Ptr(true),
		// 					QuotedName: to.Ptr("[dbo].[myTable]"),
		// 			}},
		// 	}},
		// }
	}
}
