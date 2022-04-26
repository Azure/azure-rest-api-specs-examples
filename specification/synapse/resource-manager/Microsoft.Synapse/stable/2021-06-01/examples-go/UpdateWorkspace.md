Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.4.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsynapse.WorkspacePatchInfo{
			Identity: &armsynapse.ManagedIdentity{
				Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armsynapse.WorkspacePatchProperties{
				Encryption: &armsynapse.EncryptionDetails{
					Cmk: &armsynapse.CustomerManagedKeyDetails{
						Key: &armsynapse.WorkspaceKeyDetails{
							Name:        to.Ptr("<name>"),
							KeyVaultURL: to.Ptr("<key-vault-url>"),
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
					PurviewResourceID: to.Ptr("<purview-resource-id>"),
				},
				SQLAdministratorLoginPassword: to.Ptr("<sqladministrator-login-password>"),
				WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
					Type:                to.Ptr("<type>"),
					AccountName:         to.Ptr("<account-name>"),
					CollaborationBranch: to.Ptr("<collaboration-branch>"),
					HostName:            to.Ptr("<host-name>"),
					ProjectName:         to.Ptr("<project-name>"),
					RepositoryName:      to.Ptr("<repository-name>"),
					RootFolder:          to.Ptr("<root-folder>"),
				},
			},
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
		},
		&armsynapse.WorkspacesClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
