package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/CredentialSetGet.json
func ExampleCredentialSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewCredentialSetsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myResourceGroup", "myRegistry", "myCredentialSet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 				PasswordSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/password"),
	// 				UsernameSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/username"),
	// 		}},
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T18:20:08.012276+00:00"); return t}()),
	// 		LoginServer: to.Ptr("docker.io"),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 	},
	// }
}
