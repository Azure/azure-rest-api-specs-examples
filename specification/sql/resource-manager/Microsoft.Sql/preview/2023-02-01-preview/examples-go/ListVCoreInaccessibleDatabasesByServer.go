package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/ListVCoreInaccessibleDatabasesByServer.json
func ExampleDatabasesClient_NewListInaccessibleByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListInaccessibleByServerPager("Default-SQL-SouthEastAsia", "testsvr", nil)
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
		// page.DatabaseListResult = armsql.DatabaseListResult{
		// 	Value: []*armsql.Database{
		// 		{
		// 			Name: to.Ptr("testdb"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Kind: to.Ptr("v12.0,user,vcore"),
		// 			Properties: &armsql.DatabaseProperties{
		// 				CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
		// 				CurrentServiceObjectiveName: to.Ptr("BC_Gen4_2"),
		// 				DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				LicenseType: to.Ptr(armsql.DatabaseLicenseTypeLicenseIncluded),
		// 				MaxLogSizeBytes: to.Ptr[int64](104857600),
		// 				MaxSizeBytes: to.Ptr[int64](268435456000),
		// 				ReadScale: to.Ptr(armsql.DatabaseReadScaleEnabled),
		// 				Status: to.Ptr(armsql.DatabaseStatusInaccessible),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("BC_Gen4"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Tier: to.Ptr("BusinessCritical"),
		// 			},
		// 	}},
		// }
	}
}
