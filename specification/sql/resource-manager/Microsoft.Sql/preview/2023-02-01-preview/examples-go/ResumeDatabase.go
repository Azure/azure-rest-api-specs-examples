package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/ResumeDatabase.json
func ExampleDatabasesClient_BeginResume() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginResume(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdwdb", nil)
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
	// res.Database = armsql.Database{
	// 	Name: to.Ptr("testdwdb"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdwdb"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Kind: to.Ptr("v12.0,user,datawarehouse"),
	// 	Properties: &armsql.DatabaseProperties{
	// 		CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-13T02:04:26.187Z"); return t}()),
	// 		CurrentServiceObjectiveName: to.Ptr("D1"),
	// 		DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
	// 		DefaultSecondaryLocation: to.Ptr("North Europe"),
	// 		MaxSizeBytes: to.Ptr[int64](10995116277760),
	// 		Status: to.Ptr(armsql.DatabaseStatusOnline),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("DataWarehouse"),
	// 	},
	// }
}
