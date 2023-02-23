package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "resourceGroup1", "workspace1", armsynapse.WorkspacePatchInfo{
		Identity: &armsynapse.ManagedIdentity{
			Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armsynapse.WorkspacePatchProperties{
			Encryption: &armsynapse.EncryptionDetails{
				Cmk: &armsynapse.CustomerManagedKeyDetails{
					Key: &armsynapse.WorkspaceKeyDetails{
						Name:        to.Ptr("default"),
						KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
					},
				},
			},
			ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
				AllowedAADTenantIDsForLinking: []*string{
					to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
				LinkedAccessCheckOnTargetResource: to.Ptr(false),
				PreventDataExfiltration:           to.Ptr(false),
			},
			PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
			PurviewConfiguration: &armsynapse.PurviewConfiguration{
				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
			},
			SQLAdministratorLoginPassword: to.Ptr("password"),
			WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
				Type:                to.Ptr("FactoryGitHubConfiguration"),
				AccountName:         to.Ptr("adifferentacount"),
				CollaborationBranch: to.Ptr("master"),
				HostName:            to.Ptr(""),
				ProjectName:         to.Ptr("myproject"),
				RepositoryName:      to.Ptr("myrepository"),
				RootFolder:          to.Ptr("/"),
			},
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// 		Encryption: &armsynapse.EncryptionDetails{
	// 			Cmk: &armsynapse.CustomerManagedKeyDetails{
	// 				Key: &armsynapse.WorkspaceKeyDetails{
	// 					Name: to.Ptr("default"),
	// 					KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
	// 				},
	// 				Status: to.Ptr("Consistent"),
	// 			},
	// 			DoubleEncryptionEnabled: to.Ptr(true),
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("resourceGroup2"),
	// 		ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
	// 			AllowedAADTenantIDsForLinking: []*string{
	// 				to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
	// 				LinkedAccessCheckOnTargetResource: to.Ptr(false),
	// 				PreventDataExfiltration: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
	// 			PurviewConfiguration: &armsynapse.PurviewConfiguration{
	// 				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
	// 			},
	// 			SQLAdministratorLogin: to.Ptr("login"),
	// 			WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 	}
}
