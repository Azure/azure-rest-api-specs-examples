package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseCreatePointInTimeRestore.json
func ExampleManagedDatabasesClient_BeginCreateOrUpdate_createsANewManagedDatabaseUsingPointInTimeRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedDatabasesClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "managedInstance", "managedDatabase", armsql.ManagedDatabase{
		Location: to.Ptr("southeastasia"),
		Properties: &armsql.ManagedDatabaseProperties{
			CreateMode:         to.Ptr(armsql.ManagedDatabaseCreateModePointInTimeRestore),
			RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-07-14T05:35:31.503Z"); return t }()),
			SourceDatabaseID:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/databases/testdb"),
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
	// 	Name: to.Ptr("testdb1"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/databases/testdb1"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("TagValue1"),
	// 	},
	// 	Properties: &armsql.ManagedDatabaseProperties{
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
	// 		DefaultSecondaryLocation: to.Ptr("North Europe"),
	// 		Status: to.Ptr(armsql.ManagedDatabaseStatusOnline),
	// 	},
	// }
}
