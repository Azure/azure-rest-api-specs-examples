package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/LogFileListByServer.json
func ExampleLogFilesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.LogFileListResult = armmariadb.LogFileListResult{
		// 	Value: []*armmariadb.LogFile{
		// 		{
		// 			Name: to.Ptr("mariadb-slow-mariadbtestsvc1-2018022823.log"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers/logFiles"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc1/logFiles/mariadb-slow-mariadbtestsvc1-2018022823.log"),
		// 			Properties: &armmariadb.LogFileProperties{
		// 				Type: to.Ptr("slowlog"),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-01T06:09:20.000Z"); return t}()),
		// 				SizeInKB: to.Ptr[int64](1),
		// 				URL: to.Ptr("https://wasd2prodwus1afse42.file.core.windows.net/833c99b2f36c47349e5554b903fe0440/serverlogs/mariadb-slow-mariadbtestsvc1-2018022823.log?sv=2015-04-05&sr=f&sig=D9Ga4N5Pa%2BPe5Bmjpvs7A0TPD%2FF7IZpk9e4KWR0jgpM%3D&se=2018-03-01T07%3A12%3A13Z&sp=r"),
		// 			},
		// 	}},
		// }
	}
}
