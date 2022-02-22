Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.Workspace{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			Identity: &armsynapse.ManagedIdentity{
				Type: armsynapse.ResourceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armsynapse.UserAssignedManagedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": {},
				},
			},
			Properties: &armsynapse.WorkspaceProperties{
				CspWorkspaceAdminProperties: &armsynapse.CspWorkspaceAdminProperties{
					InitialWorkspaceAdminObjectID: to.StringPtr("<initial-workspace-admin-object-id>"),
				},
				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
					AccountURL: to.StringPtr("<account-url>"),
					Filesystem: to.StringPtr("<filesystem>"),
				},
				Encryption: &armsynapse.EncryptionDetails{
					Cmk: &armsynapse.CustomerManagedKeyDetails{
						KekIdentity: &armsynapse.KekIdentityProperties{
							UseSystemAssignedIdentity: false,
							UserAssignedIdentity:      to.StringPtr("<user-assigned-identity>"),
						},
						Key: &armsynapse.WorkspaceKeyDetails{
							Name:        to.StringPtr("<name>"),
							KeyVaultURL: to.StringPtr("<key-vault-url>"),
						},
					},
				},
				ManagedResourceGroupName: to.StringPtr("<managed-resource-group-name>"),
				ManagedVirtualNetwork:    to.StringPtr("<managed-virtual-network>"),
				ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
					AllowedAADTenantIDsForLinking: []*string{
						to.StringPtr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
					LinkedAccessCheckOnTargetResource: to.BoolPtr(false),
					PreventDataExfiltration:           to.BoolPtr(false),
				},
				PublicNetworkAccess: armsynapse.WorkspacePublicNetworkAccess("Enabled").ToPtr(),
				PurviewConfiguration: &armsynapse.PurviewConfiguration{
					PurviewResourceID: to.StringPtr("<purview-resource-id>"),
				},
				SQLAdministratorLogin:         to.StringPtr("<sqladministrator-login>"),
				SQLAdministratorLoginPassword: to.StringPtr("<sqladministrator-login-password>"),
				WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
					Type:                to.StringPtr("<type>"),
					AccountName:         to.StringPtr("<account-name>"),
					CollaborationBranch: to.StringPtr("<collaboration-branch>"),
					HostName:            to.StringPtr("<host-name>"),
					ProjectName:         to.StringPtr("<project-name>"),
					RepositoryName:      to.StringPtr("<repository-name>"),
					RootFolder:          to.StringPtr("<root-folder>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientCreateOrUpdateResult)
}
```
