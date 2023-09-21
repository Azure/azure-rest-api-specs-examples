package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksCreate_QuickTask.json
func ExampleTasksClient_BeginCreate_tasksCreateQuickTask() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTasksClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "quicktask", armcontainerregistry.Task{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"testkey": to.Ptr("value"),
		},
		Properties: &armcontainerregistry.TaskProperties{
			IsSystemTask: to.Ptr(true),
			LogTemplate:  to.Ptr("acr/tasks:{{.Run.OS}}"),
			Status:       to.Ptr(armcontainerregistry.TaskStatusEnabled),
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
	// res.Task = armcontainerregistry.Task{
	// 	Name: to.Ptr("quicktask"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/tasks"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tasks/myTask"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"testkey": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.TaskProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:54:23.536Z"); return t}()),
	// 		IsSystemTask: to.Ptr(true),
	// 		LogTemplate: to.Ptr("acr/tasks:{{.Run.OS}}"),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		Status: to.Ptr(armcontainerregistry.TaskStatusEnabled),
	// 	},
	// }
}
