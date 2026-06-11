package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/listKeys.json
func ExampleWorkspacesClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().ListKeys(ctx, "testrg123", "workspaces123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.WorkspacesClientListKeysResponse{
	// 	ListWorkspaceKeysResult: armmachinelearning.ListWorkspaceKeysResult{
	// 		ContainerRegistryCredentials: &armmachinelearning.RegistryListCredentialsResult{
	// 			Passwords: []*armmachinelearning.Password{
	// 				{
	// 					Name: to.Ptr("password"),
	// 					Value: to.Ptr("<value>"),
	// 				},
	// 				{
	// 					Name: to.Ptr("password2"),
	// 					Value: to.Ptr("0KARRQoQHSUq1yViPWg7YFernOS=Ic/t"),
	// 				},
	// 			},
	// 			Username: to.Ptr("testdemoworkjmjmeykp"),
	// 		},
	// 		NotebookAccessKeys: &armmachinelearning.ListNotebookKeysResult{
	// 		},
	// 		UserStorageArmID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/ragargeastus2euap/providers/Microsoft.Storage/storageAccounts/testdemoworkazashomr"),
	// 	},
	// }
}
