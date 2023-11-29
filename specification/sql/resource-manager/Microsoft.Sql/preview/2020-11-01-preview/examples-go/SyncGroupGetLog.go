package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncGroupGetLog.json
func ExampleSyncGroupsClient_NewListLogsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncGroupsClient().NewListLogsPager("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "2017-01-01T00:00:00", "2017-12-31T00:00:00", armsql.SyncGroupsTypeAll, &armsql.SyncGroupsClientListLogsOptions{ContinuationToken: nil})
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
		// page.SyncGroupLogListResult = armsql.SyncGroupLogListResult{
		// 	Value: []*armsql.SyncGroupLogProperties{
		// 		{
		// 			Type: to.Ptr(armsql.SyncGroupLogTypeSuccess),
		// 			OperationStatus: to.Ptr("SchemaRefreshSuccess"),
		// 			Source: to.Ptr("syncgroupcrud-8475.database.windows.net/hub"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-30T07:16:08.250Z"); return t}()),
		// 			TracingID: to.Ptr("c0480c8e-6269-424e-9404-b00efce0ebae"),
		// 			Details: to.Ptr("Schema information obtained successfully."),
		// 		},
		// 		{
		// 			Type: to.Ptr(armsql.SyncGroupLogTypeError),
		// 			OperationStatus: to.Ptr("SchemaRefreshFailure"),
		// 			Source: to.Ptr("syncgroupcrud-8475.database.windows.net/member"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-30T07:03:37.573Z"); return t}()),
		// 			TracingID: to.Ptr("cd3aa28c-0c31-471f-8a77-f1b21c908cbd"),
		// 			Details: to.Ptr("Getting schema information for the database failed with the exception \"Failed to connect to server .\nInner exception: SqlException Error Code: -2146232060 - SqlError Number:53, Message: A network-related or instance-specific error occurred while establishing a connection to SQL Server. The server was not found or was not accessible. Verify that the instance name is correct and that SQL Server is configured to allow remote connections. (provider: Named Pipes Provider, error: 40 - Could not open a connection to SQL Server) \nInner exception: The network path was not found\n For more information, provide tracing ID ‘cd3aa28c-0c31-471f-8a77-f1b21c908cbd’ to customer support.\""),
		// 	}},
		// }
	}
}
