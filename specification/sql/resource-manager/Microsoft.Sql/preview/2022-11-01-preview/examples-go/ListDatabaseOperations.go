package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/ListDatabaseOperations.json
func ExampleDatabaseOperationsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseOperationsClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "testdb", nil)
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
		// page.DatabaseOperationListResult = armsql.DatabaseOperationListResult{
		// 	Value: []*armsql.DatabaseOperation{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/operations/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armsql.DatabaseOperationProperties{
		// 				Description: to.Ptr("'UpdateLogicalDatabase' on database 'testdb', Source ServiceLevelObjective 'SQLDB_GP_Gen5_2', target ServiceLevelObjective 'SQLDB_HS_Gen5_2', target database MaxSize '32'GB"),
		// 				DatabaseName: to.Ptr("testdb"),
		// 				EstimatedCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-17T14:30:30.710Z"); return t}()),
		// 				IsCancellable: to.Ptr(true),
		// 				Operation: to.Ptr("UpdateLogicalDatabase"),
		// 				OperationFriendlyName: to.Ptr("ALTER DATABASE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("testsvr"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T06:25:23.670Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-5555-5555-5555-555555555555"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/operations/55555555-5555-5555-5555-555555555555"),
		// 			Properties: &armsql.DatabaseOperationProperties{
		// 				Description: to.Ptr("'UpdateLogicalDatabase' on database 'testdb', Source ServiceLevelObjective 'SQLDB_GP_Gen5_2', target ServiceLevelObjective 'SQLDB_HS_Gen5_2', target database MaxSize '32'GB"),
		// 				DatabaseName: to.Ptr("testdb"),
		// 				EstimatedCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-17T14:30:30.710Z"); return t}()),
		// 				IsCancellable: to.Ptr(true),
		// 				Operation: to.Ptr("UpdateLogicalDatabase"),
		// 				OperationFriendlyName: to.Ptr("ALTER DATABASE"),
		// 				OperationPhaseDetails: &armsql.PhaseDetails{
		// 					Phase: to.Ptr(armsql.PhaseWaitingForCutover),
		// 					PhaseInformation: map[string]*string{
		// 						"currentStep": to.Ptr("3"),
		// 						"dataDelayInMb": to.Ptr("31"),
		// 						"performCutoverBy": to.Ptr("2023-02-17T11:57:06.71Z"),
		// 						"totalSteps": to.Ptr("4"),
		// 					},
		// 				},
		// 				PercentComplete: to.Ptr[int32](19),
		// 				ServerName: to.Ptr("testsvr"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T06:25:23.670Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateInProgress),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-5555-5555-5555-555555555555"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/operations/55555555-5555-5555-5555-555555555555"),
		// 			Properties: &armsql.DatabaseOperationProperties{
		// 				Description: to.Ptr("'UpdateLogicalDatabase' on database 'testdb'"),
		// 				DatabaseName: to.Ptr("testdb"),
		// 				ErrorCode: to.Ptr[int32](40640),
		// 				ErrorDescription: to.Ptr("The server encountered an unexpected exception."),
		// 				ErrorSeverity: to.Ptr[int32](20),
		// 				IsUserError: to.Ptr(true),
		// 				Operation: to.Ptr("UpdateLogicalDatabase"),
		// 				OperationFriendlyName: to.Ptr("ALTER DATABASE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("testsvr"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T11:48:29.160Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateFailed),
		// 			},
		// 	}},
		// }
	}
}
