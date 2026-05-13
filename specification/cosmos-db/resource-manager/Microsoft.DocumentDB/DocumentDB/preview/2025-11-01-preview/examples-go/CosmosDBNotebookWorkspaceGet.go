package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBNotebookWorkspaceGet.json
func ExampleNotebookWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNotebookWorkspacesClient().Get(ctx, "rg1", "ddb1", armcosmos.NotebookWorkspaceNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.NotebookWorkspacesClientGetResponse{
	// 	NotebookWorkspace: &armcosmos.NotebookWorkspace{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/notebookWorkspaces"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/notebookWorkspaces/default"),
	// 		Properties: &armcosmos.NotebookWorkspaceProperties{
	// 			NotebookServerEndpoint: to.Ptr("endpoint"),
	// 			Status: to.Ptr("Online"),
	// 		},
	// 	},
	// }
}
