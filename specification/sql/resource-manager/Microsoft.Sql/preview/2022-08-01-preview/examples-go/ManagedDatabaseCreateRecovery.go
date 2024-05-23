package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedDatabaseCreateRecovery.json
func ExampleManagedDatabasesClient_BeginCreateOrUpdate_createsANewManagedDatabaseFromRestoringAGeoReplicatedBackup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedDatabasesClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "server1", "testdb_recovered", armsql.ManagedDatabase{
		Location: to.Ptr("southeastasia"),
		Properties: &armsql.ManagedDatabaseProperties{
			CreateMode:            to.Ptr(armsql.ManagedDatabaseCreateModeRecovery),
			RecoverableDatabaseID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/Default-SQL-WestEurope/providers/Microsoft.Sql/managedInstances/testsvr/recoverableDatabases/testdb"),
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
	// res.ManagedDatabase = armsql.ManagedDatabase{
	// 	Name: to.Ptr("testdb_recovered"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/server1/recoverableDatabases/testdb_recovered"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("TagValue1"),
	// 	},
	// 	Properties: &armsql.ManagedDatabaseProperties{
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-07T04:41:33.937Z"); return t}()),
	// 		IsLedgerOn: to.Ptr(false),
	// 		Status: to.Ptr(armsql.ManagedDatabaseStatusOnline),
	// 	},
	// }
}
