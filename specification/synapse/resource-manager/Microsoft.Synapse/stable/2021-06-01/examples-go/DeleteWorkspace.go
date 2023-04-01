package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspace.json
func ExampleWorkspacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginDelete(ctx, "resourceGroup1", "workspace1", nil)
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
	// res.Workspace = armsynapse.Workspace{
	// 	Name: to.Ptr("workspace1"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armsynapse.WorkspaceProperties{
	// 		ConnectivityEndpoints: map[string]*string{
	// 			"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
	// 			"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
	// 		},
	// 		DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
	// 			AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
	// 			Filesystem: to.Ptr("default"),
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("resourceGroup2"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SQLAdministratorLogin: to.Ptr("login"),
	// 	},
	// }
}
