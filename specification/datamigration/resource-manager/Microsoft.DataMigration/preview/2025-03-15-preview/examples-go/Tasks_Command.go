package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/Tasks_Command.json
func ExampleTasksClient_Command() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().Command(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "DmsSdkTask", &armdatamigration.MigrateSyncCompleteCommandProperties{
		CommandType: to.Ptr(armdatamigration.CommandTypeMigrateSyncCompleteDatabase),
		Input: &armdatamigration.MigrateSyncCompleteCommandInput{
			DatabaseName: to.Ptr("TestDatabase"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.TasksClientCommandResponse{
	// 	                            CommandPropertiesClassification: &armdatamigration.MigrateSyncCompleteCommandProperties{
	// 		CommandType: to.Ptr(armdatamigration.CommandTypeMigrateSyncCompleteDatabase),
	// 		State: to.Ptr(armdatamigration.CommandStateAccepted),
	// 		Input: &armdatamigration.MigrateSyncCompleteCommandInput{
	// 			DatabaseName: to.Ptr("TestDatabase"),
	// 		},
	// 	},
	// 	                        }
}
