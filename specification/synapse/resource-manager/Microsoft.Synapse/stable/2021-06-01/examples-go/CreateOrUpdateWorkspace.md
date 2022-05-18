Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.5.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resourceGroup1",
		"workspace1",
		armsynapse.Workspace{
			Location: to.Ptr("East US"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Identity: &armsynapse.ManagedIdentity{
				Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armsynapse.UserAssignedManagedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": {},
				},
			},
			Properties: &armsynapse.WorkspaceProperties{
				CspWorkspaceAdminProperties: &armsynapse.CspWorkspaceAdminProperties{
					InitialWorkspaceAdminObjectID: to.Ptr("6c20646f-8050-49ec-b3b1-80a0e58e454d"),
				},
				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
					AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
					Filesystem: to.Ptr("default"),
				},
				Encryption: &armsynapse.EncryptionDetails{
					Cmk: &armsynapse.CustomerManagedKeyDetails{
						KekIdentity: &armsynapse.KekIdentityProperties{
							UseSystemAssignedIdentity: false,
							UserAssignedIdentity:      to.Ptr("/subscriptions/b64d7b94-73e7-4d36-94b2-7764ea3fd74a/resourcegroups/SynapseCI/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
						},
						Key: &armsynapse.WorkspaceKeyDetails{
							Name:        to.Ptr("default"),
							KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
						},
					},
				},
				ManagedResourceGroupName: to.Ptr("workspaceManagedResourceGroupUnique"),
				ManagedVirtualNetwork:    to.Ptr("default"),
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
				SQLAdministratorLogin:         to.Ptr("login"),
				SQLAdministratorLoginPassword: to.Ptr("password"),
				WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
					Type:                to.Ptr("FactoryGitHubConfiguration"),
					AccountName:         to.Ptr("mygithubaccount"),
					CollaborationBranch: to.Ptr("master"),
					HostName:            to.Ptr(""),
					ProjectName:         to.Ptr("myproject"),
					RepositoryName:      to.Ptr("myrepository"),
					RootFolder:          to.Ptr("/"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
