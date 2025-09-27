package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/ServiceTasks_Cancel.json
func ExampleServiceTasksClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceTasksClient().Cancel(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkTask", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectTask = armdatamigration.ProjectTask{
	// 	Name: to.Ptr("DmsSdkTask"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/serviceTasks"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/serviceTasks/DmsSdkTask"),
	// 	Etag: to.Ptr("0vPYxzfnDaDH9yhOJAnqTyTRpa09Kb7pm+LEukDBbw8="),
	// 	Properties: &armdatamigration.ConnectToSourceMySQLTaskProperties{
	// 		State: to.Ptr(armdatamigration.TaskStateQueued),
	// 		TaskType: to.Ptr(armdatamigration.TaskTypeConnectToSourceMySQL),
	// 		Input: &armdatamigration.ConnectToSourceMySQLTaskInput{
	// 			SourceConnectionInfo: &armdatamigration.MySQLConnectionInfo{
	// 				Type: to.Ptr("MySqlConnectionInfo"),
	// 				Port: to.Ptr[int32](3306),
	// 				ServerName: to.Ptr("localhost"),
	// 			},
	// 		},
	// 	},
	// }
}
