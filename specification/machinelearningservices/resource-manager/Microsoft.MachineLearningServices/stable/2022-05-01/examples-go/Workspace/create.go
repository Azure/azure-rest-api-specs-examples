package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/Workspace/create.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"workspace-1234",
		"testworkspace",
		armmachinelearning.Workspace{
			Identity: &armmachinelearning.ManagedServiceIdentity{
				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": {},
				},
			},
			Location: to.Ptr("eastus2euap"),
			Properties: &armmachinelearning.WorkspaceProperties{
				Description:         to.Ptr("test description"),
				ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
				ContainerRegistry:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
				Encryption: &armmachinelearning.EncryptionProperty{
					Identity: &armmachinelearning.IdentityForCmk{
						UserAssignedIdentity: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai"),
					},
					KeyVaultProperties: &armmachinelearning.EncryptionKeyVaultProperties{
						IdentityClientID: to.Ptr(""),
						KeyIdentifier:    to.Ptr("https://testkv.vault.azure.net/keys/testkey/aabbccddee112233445566778899aabb"),
						KeyVaultArmID:    to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
					},
					Status: to.Ptr(armmachinelearning.EncryptionStatusEnabled),
				},
				FriendlyName: to.Ptr("HelloName"),
				HbiWorkspace: to.Ptr(false),
				KeyVault:     to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
				SharedPrivateLinkResources: []*armmachinelearning.SharedPrivateLinkResource{
					{
						Name: to.Ptr("testdbresource"),
						Properties: &armmachinelearning.SharedPrivateLinkResourceProperty{
							GroupID:               to.Ptr("Sql"),
							PrivateLinkResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.DocumentDB/databaseAccounts/testdbresource/privateLinkResources/Sql"),
							RequestMessage:        to.Ptr("Please approve"),
							Status:                to.Ptr(armmachinelearning.PrivateEndpointServiceConnectionStatusApproved),
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
