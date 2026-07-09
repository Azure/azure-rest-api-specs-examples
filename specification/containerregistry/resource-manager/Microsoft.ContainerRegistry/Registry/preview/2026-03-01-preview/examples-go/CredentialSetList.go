package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-03-01-preview/CredentialSetList.json
func ExampleCredentialSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCredentialSetsClient().NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcontainerregistry.CredentialSetsClientListResponse{
		// 	CredentialSetListResult: armcontainerregistry.CredentialSetListResult{
		// 		Value: []*armcontainerregistry.CredentialSet{
		// 			{
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/credentialSets"),
		// 				Identity: &armcontainerregistry.IdentityProperties{
		// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/credentialSets/myCredentialSet"),
		// 				Name: to.Ptr("myCredentialSet"),
		// 				Properties: &armcontainerregistry.CredentialSetProperties{
		// 					LoginServer: to.Ptr("docker.io"),
		// 					AuthCredentials: []*armcontainerregistry.AuthCredential{
		// 						{
		// 							Name: to.Ptr(armcontainerregistry.CredentialNameCredential1),
		// 							UsernameSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/username"),
		// 							PasswordSecretIdentifier: to.Ptr("https://myvault.vault.azure.net/secrets/password"),
		// 							CredentialHealth: &armcontainerregistry.CredentialHealth{
		// 								Status: to.Ptr(armcontainerregistry.CredentialHealthStatusHealthy),
		// 							},
		// 						},
		// 					},
		// 					CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T18:20:08.012276+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
