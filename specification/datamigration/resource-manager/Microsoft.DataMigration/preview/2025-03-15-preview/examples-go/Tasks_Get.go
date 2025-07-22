package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/Tasks_Get.json
func ExampleTasksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().Get(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "DmsSdkTask", &armdatamigration.TasksClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectTask = armdatamigration.ProjectTask{
	// 	Name: to.Ptr("DmsSdkTask"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects/tasks"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/tasks/DmsSdkTask"),
	// 	Etag: to.Ptr("0vPYxzfnDaDH9yhOJAnqTyTRpa09Kb7pm+LEukDBbw8="),
	// 	Properties: &armdatamigration.ConnectToTargetSQLDbTaskProperties{
	// 		State: to.Ptr(armdatamigration.TaskStateQueued),
	// 		TaskType: to.Ptr(armdatamigration.TaskTypeConnectToTargetSQLDb),
	// 		Input: &armdatamigration.ConnectToTargetSQLDbTaskInput{
	// 			TargetConnectionInfo: &armdatamigration.SQLConnectionInfo{
	// 				Type: to.Ptr("SqlConnectionInfo"),
	// 				UserName: to.Ptr("testuser"),
	// 				Authentication: to.Ptr(armdatamigration.AuthenticationTypeSQLAuthentication),
	// 				DataSource: to.Ptr("ssma-test-server.database.windows.net"),
	// 				EncryptConnection: to.Ptr(true),
	// 				TrustServerCertificate: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
