Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv1.0.0/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/create.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"workspace-1234",
		"testworkspace",
		armmachinelearningservices.Workspace{
			Identity: &armmachinelearningservices.Identity{
				Type: to.Ptr(armmachinelearningservices.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": {},
				},
			},
			Location: to.Ptr("eastus2euap"),
			Properties: &armmachinelearningservices.WorkspaceProperties{
				Description:         to.Ptr("test description"),
				ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
				ContainerRegistry:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
				Encryption: &armmachinelearningservices.EncryptionProperty{
					Identity: &armmachinelearningservices.IdentityForCmk{
						UserAssignedIdentity: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai"),
					},
					KeyVaultProperties: &armmachinelearningservices.KeyVaultProperties{
						IdentityClientID: to.Ptr(""),
						KeyIdentifier:    to.Ptr("https://testkv.vault.azure.net/keys/testkey/aabbccddee112233445566778899aabb"),
						KeyVaultArmID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
					},
					Status: to.Ptr(armmachinelearningservices.EncryptionStatusEnabled),
				},
				FriendlyName: to.Ptr("HelloName"),
				HbiWorkspace: to.Ptr(false),
				KeyVault:     to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
				SharedPrivateLinkResources: []*armmachinelearningservices.SharedPrivateLinkResource{
					{
						Name: to.Ptr("testdbresource"),
						Properties: &armmachinelearningservices.SharedPrivateLinkResourceProperty{
							GroupID:               to.Ptr("Sql"),
							PrivateLinkResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.DocumentDB/databaseAccounts/testdbresource/privateLinkResources/Sql"),
							RequestMessage:        to.Ptr("Please approve"),
							Status:                to.Ptr(armmachinelearningservices.PrivateEndpointServiceConnectionStatusApproved),
						},
					}},
				StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
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
