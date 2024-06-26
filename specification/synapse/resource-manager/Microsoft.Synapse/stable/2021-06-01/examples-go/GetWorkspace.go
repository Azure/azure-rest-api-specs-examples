package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspace.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "resourceGroup1", "workspace1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 		CspWorkspaceAdminProperties: &armsynapse.CspWorkspaceAdminProperties{
	// 			InitialWorkspaceAdminObjectID: to.Ptr("6c20646f-8050-49ec-b3b1-80a0e58e454d"),
	// 		},
	// 		DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
	// 			AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
	// 			Filesystem: to.Ptr("default"),
	// 		},
	// 		ExtraProperties: map[string]any{
	// 			"IsScopeEnabled": "false",
	// 			"WorkspaceType": "Normal",
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("resourceGroup2"),
	// 		ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
	// 			AllowedAADTenantIDsForLinking: []*string{
	// 				to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
	// 				LinkedAccessCheckOnTargetResource: to.Ptr(false),
	// 				PreventDataExfiltration: to.Ptr(false),
	// 			},
	// 			PrivateEndpointConnections: []*armsynapse.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("sql"),
	// 					Type: to.Ptr("Microsoft.Synapse/workspaces/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/privateEndpointConnections/ExamplePrivateEndpointConnection"),
	// 					Properties: &armsynapse.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armsynapse.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/privateEndpoints/ExamplePrivateEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armsynapse.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr("Approved"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PurviewConfiguration: &armsynapse.PurviewConfiguration{
	// 				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
	// 			},
	// 			SQLAdministratorLogin: to.Ptr("login"),
	// 			WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
	// 				Type: to.Ptr("FactoryGitHubConfiguration"),
	// 				AccountName: to.Ptr("myGithubAccount"),
	// 				CollaborationBranch: to.Ptr("master"),
	// 				HostName: to.Ptr(""),
	// 				ProjectName: to.Ptr("myProject"),
	// 				RepositoryName: to.Ptr("myRepository"),
	// 				RootFolder: to.Ptr("/"),
	// 				TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 			WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 	}
}
