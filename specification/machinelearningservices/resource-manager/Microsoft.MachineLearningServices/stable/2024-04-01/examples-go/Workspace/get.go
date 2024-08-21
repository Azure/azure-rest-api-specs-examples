package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/get.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "workspace-1234", "testworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armmachinelearning.Workspace{
	// 	Name: to.Ptr("testworkspace"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace"),
	// 	Identity: &armmachinelearning.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
	// 			"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": &armmachinelearning.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armmachinelearning.WorkspaceProperties{
	// 		Description: to.Ptr("test description"),
	// 		AllowPublicAccessWhenBehindVnet: to.Ptr(false),
	// 		ApplicationInsights: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/microsoft.insights/components/testinsights"),
	// 		ContainerRegistry: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ContainerRegistry/registries/testRegistry"),
	// 		DiscoveryURL: to.Ptr("http://example.com"),
	// 		Encryption: &armmachinelearning.EncryptionProperty{
	// 			Identity: &armmachinelearning.IdentityForCmk{
	// 				UserAssignedIdentity: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai"),
	// 			},
	// 			KeyVaultProperties: &armmachinelearning.EncryptionKeyVaultProperties{
	// 				IdentityClientID: to.Ptr(""),
	// 				KeyIdentifier: to.Ptr("https://testkv.vault.azure.net/keys/testkey/aabbccddee112233445566778899aabb"),
	// 				KeyVaultArmID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
	// 			},
	// 			Status: to.Ptr(armmachinelearning.EncryptionStatusEnabled),
	// 		},
	// 		FriendlyName: to.Ptr("HelloName"),
	// 		HbiWorkspace: to.Ptr(false),
	// 		ImageBuildCompute: to.Ptr("testcompute"),
	// 		KeyVault: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.KeyVault/vaults/testkv"),
	// 		PrivateEndpointConnections: []*armmachinelearning.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("testprivatelinkconnection"),
	// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/privateEndpointConnections/testprivatelinkconnection"),
	// 				Properties: &armmachinelearning.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armmachinelearning.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg-1234/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armmachinelearning.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Auto-Approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armmachinelearning.PrivateEndpointServiceConnectionStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armmachinelearning.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		PrivateLinkCount: to.Ptr[int32](0),
	// 		PublicNetworkAccess: to.Ptr(armmachinelearning.PublicNetworkAccessDisabled),
	// 		ServiceProvisionedResourceGroup: to.Ptr("testworkspace_0000111122223333"),
	// 		SharedPrivateLinkResources: []*armmachinelearning.SharedPrivateLinkResource{
	// 			{
	// 				Name: to.Ptr("testcosmosdbresource"),
	// 				Properties: &armmachinelearning.SharedPrivateLinkResourceProperty{
	// 					GroupID: to.Ptr("Sql"),
	// 					PrivateLinkResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.DocumentDB/databaseAccounts/testcosmosdbresource/privateLinkResources/Sql"),
	// 					RequestMessage: to.Ptr("Please approve"),
	// 					Status: to.Ptr(armmachinelearning.PrivateEndpointServiceConnectionStatusApproved),
	// 				},
	// 		}},
	// 		StorageAccount: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
	// 	},
	// }
}
