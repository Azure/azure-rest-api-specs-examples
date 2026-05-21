package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/ServiceTasks_Get.json
func ExampleServiceTasksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceTasksClient().Get(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkTask", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.ServiceTasksClientGetResponse{
	// 	ProjectTask: armdatamigration.ProjectTask{
	// 		Name: to.Ptr("DmsSdkTask"),
	// 		Type: to.Ptr("Microsoft.DataMigration/services/serviceTasks"),
	// 		Etag: to.Ptr("0vPYxzfnDaDH9yhOJAnqTyTRpa09Kb7pm+LEukDBbw8="),
	// 		ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/serviceTasks/DmsSdkTask"),
	// 		Properties: &armdatamigration.ConnectToSourceMySQLTaskProperties{
	// 			Input: &armdatamigration.ConnectToSourceMySQLTaskInput{
	// 				SourceConnectionInfo: &armdatamigration.MySQLConnectionInfo{
	// 					Type: to.Ptr("MySqlConnectionInfo"),
	// 					Port: to.Ptr[int32](3306),
	// 					ServerName: to.Ptr("localhost"),
	// 				},
	// 			},
	// 			State: to.Ptr(armdatamigration.TaskStateQueued),
	// 			TaskType: to.Ptr(armdatamigration.TaskTypeConnectToSourceMySQL),
	// 		},
	// 	},
	// }
}
