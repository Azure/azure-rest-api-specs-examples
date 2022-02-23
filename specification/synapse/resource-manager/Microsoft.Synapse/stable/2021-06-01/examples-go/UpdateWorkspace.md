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

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.WorkspacePatchInfo{
			Identity: &armsynapse.ManagedIdentity{
				Type: armsynapse.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armsynapse.WorkspacePatchProperties{
				Encryption: &armsynapse.EncryptionDetails{
					Cmk: &armsynapse.CustomerManagedKeyDetails{
						Key: &armsynapse.WorkspaceKeyDetails{
							Name:        to.StringPtr("<name>"),
							KeyVaultURL: to.StringPtr("<key-vault-url>"),
						},
					},
				},
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
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
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
	log.Printf("Response result: %#v\n", res.WorkspacesClientUpdateResult)
}
```
