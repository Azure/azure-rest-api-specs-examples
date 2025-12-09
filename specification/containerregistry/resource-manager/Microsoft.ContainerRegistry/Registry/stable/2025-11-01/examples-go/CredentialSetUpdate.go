package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/examples/CredentialSetUpdate.json
func ExampleCredentialSetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCredentialSetsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myCredentialSet", armcontainerregistry.CredentialSetUpdateParameters{
		Properties: &armcontainerregistry.CredentialSetUpdateProperties{
			AuthCredentials: []*armcontainerregistry.AuthCredential{
				{
					Name:                     to.Ptr(armcontainerregistry.CredentialNameCredential1),
					PasswordSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/password2"),
					UsernameSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/username2"),
				}},
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
	// res.CredentialSet = armcontainerregistry.CredentialSet{
	// 	Name: to.Ptr("myCredentialSet"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/credentialSets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/credentialSets/myCredentialSet"),
	// 	Identity: &armcontainerregistry.IdentityProperties{
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// 	Properties: &armcontainerregistry.CredentialSetProperties{
	// 		AuthCredentials: []*armcontainerregistry.AuthCredential{
	// 			{
	// 				Name: to.Ptr(armcontainerregistry.CredentialNameCredential1),
	// 				CredentialHealth: &armcontainerregistry.CredentialHealth{
	// 					Status: to.Ptr(armcontainerregistry.CredentialHealthStatusHealthy),
	// 				},
	// 				PasswordSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/password2"),
	// 				UsernameSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/username2"),
	// 		}},
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T18:20:08.012Z"); return t}()),
	// 		LoginServer: to.Ptr("docker.io"),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 	},
	// }
}
