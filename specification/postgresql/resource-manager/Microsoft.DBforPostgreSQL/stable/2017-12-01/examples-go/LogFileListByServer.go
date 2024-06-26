package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/LogFileListByServer.json
func ExampleLogFilesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLogFilesClient().NewListByServerPager("TestGroup", "testserver", nil)
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
		// page.LogFileListResult = armpostgresql.LogFileListResult{
		// 	Value: []*armpostgresql.LogFile{
		// 		{
		// 			Name: to.Ptr("postgresql-2017-06-22_010000.log"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/logFiles"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/servers/testserver/logFiles/postgresql-2017-06-22_010000.log"),
		// 			Properties: &armpostgresql.LogFileProperties{
		// 				Type: to.Ptr("text"),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-22T01:59:36.000Z"); return t}()),
		// 				SizeInKB: to.Ptr[int64](4),
		// 				URL: to.Ptr("https://wasd2stageneu1btlm4.file.core.windows.net/42679871f6cc4302b39ab9c2e3044df3/pg_log/postgresql-2017-06-22_010000.log?sv=2015-04-05&sr=f&sig=gqIQsa6VyGyUNpzYYPWLP5gM%2BeF1so9GYbHKu6Zs0DM%3D&se=2017-06-22T03%3A21%3A09Z&sp=r"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("postgresql-2017-06-22_020000.log"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/logFiles"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/servers/testserver/logFiles/postgresql-2017-06-22_020000.log"),
		// 			Properties: &armpostgresql.LogFileProperties{
		// 				Type: to.Ptr("text"),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-22T02:19:36.000Z"); return t}()),
		// 				SizeInKB: to.Ptr[int64](1),
		// 				URL: to.Ptr("https://wasd2stageneu1btlm4.file.core.windows.net/42679871f6cc4302b39ab9c2e3044df3/pg_log/postgresql-2017-06-22_020000.log?sv=2015-04-05&sr=f&sig=i99UWBlYfR0tKaxix8yHAOnfym4HV9Auto6BbZogyRs%3D&se=2017-06-22T03%3A21%3A09Z&sp=r"),
		// 			},
		// 	}},
		// }
	}
}
