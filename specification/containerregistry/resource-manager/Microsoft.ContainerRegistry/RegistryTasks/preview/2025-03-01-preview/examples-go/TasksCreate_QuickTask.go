package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/TasksCreate_QuickTask.json
func ExampleTasksClient_Create_tasksCreateQuickTask() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().Create(ctx, "myResourceGroup", "myRegistry", "quicktask", armcontainerregistrytasks.Task{
		Location: to.Ptr("eastus"),
		Properties: &armcontainerregistrytasks.TaskProperties{
			IsSystemTask: to.Ptr(true),
			LogTemplate:  to.Ptr("acr/tasks:{{.Run.OS}}"),
			Status:       to.Ptr(armcontainerregistrytasks.TaskStatusEnabled),
		},
		Tags: map[string]*string{
			"testkey": to.Ptr("value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistrytasks.TasksClientCreateResponse{
	// 	Task: &armcontainerregistrytasks.Task{
	// 		Name: to.Ptr("quicktask"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/tasks"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tasks/myTask"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armcontainerregistrytasks.TaskProperties{
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:54:23.536Z"); return t}()),
	// 			IsSystemTask: to.Ptr(true),
	// 			LogTemplate: to.Ptr("acr/tasks:{{.Run.OS}}"),
	// 			ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
	// 			Status: to.Ptr(armcontainerregistrytasks.TaskStatusEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"testkey": to.Ptr("value"),
	// 		},
	// 	},
	// }
}
