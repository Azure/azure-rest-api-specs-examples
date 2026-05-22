package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/Tasks_Update.json
func ExampleTasksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().Update(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "DmsSdkTask", armdatamigration.ProjectTask{
		Properties: &armdatamigration.ConnectToTargetSQLDbTaskProperties{
			Input: &armdatamigration.ConnectToTargetSQLDbTaskInput{
				TargetConnectionInfo: &armdatamigration.SQLConnectionInfo{
					Type:                   to.Ptr("SqlConnectionInfo"),
					Authentication:         to.Ptr(armdatamigration.AuthenticationTypeSQLAuthentication),
					DataSource:             to.Ptr("ssma-test-server.database.windows.net"),
					EncryptConnection:      to.Ptr(true),
					Password:               to.Ptr("testpassword"),
					TrustServerCertificate: to.Ptr(true),
					UserName:               to.Ptr("testuser"),
				},
			},
			TaskType: to.Ptr(armdatamigration.TaskTypeConnectToTargetSQLDb),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.TasksClientUpdateResponse{
	// 	ProjectTask: armdatamigration.ProjectTask{
	// 		Name: to.Ptr("DmsSdkTask"),
	// 		Type: to.Ptr("Microsoft.DataMigration/services/projects/tasks"),
	// 		Etag: to.Ptr("0vPYxzfnDaDH9yhOJAnqTyTRpa09Kb7pm+LEukDBbw8="),
	// 		ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/tasks/DmsSdkTask"),
	// 		Properties: &armdatamigration.ConnectToTargetSQLDbTaskProperties{
	// 			Input: &armdatamigration.ConnectToTargetSQLDbTaskInput{
	// 				TargetConnectionInfo: &armdatamigration.SQLConnectionInfo{
	// 					Type: to.Ptr("SqlConnectionInfo"),
	// 					Authentication: to.Ptr(armdatamigration.AuthenticationTypeSQLAuthentication),
	// 					DataSource: to.Ptr("ssma-test-server.database.windows.net"),
	// 					EncryptConnection: to.Ptr(true),
	// 					TrustServerCertificate: to.Ptr(true),
	// 					UserName: to.Ptr("testuser"),
	// 				},
	// 			},
	// 			State: to.Ptr(armdatamigration.TaskStateQueued),
	// 			TaskType: to.Ptr(armdatamigration.TaskTypeConnectToTargetSQLDb),
	// 		},
	// 	},
	// }
}
