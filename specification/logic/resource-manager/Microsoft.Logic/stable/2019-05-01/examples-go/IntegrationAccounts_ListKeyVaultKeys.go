package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_ListKeyVaultKeys.json
func ExampleIntegrationAccountsClient_NewListKeyVaultKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationAccountsClient().NewListKeyVaultKeysPager("testResourceGroup", "testIntegrationAccount", armlogic.ListKeyVaultKeysDefinition{
		KeyVault: &armlogic.KeyVaultReference{
			ID: to.Ptr("subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
		},
		SkipToken: to.Ptr("testSkipToken"),
	}, nil)
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
		// page.KeyVaultKeyCollection = armlogic.KeyVaultKeyCollection{
		// 	SkipToken: to.Ptr("testSkipToken"),
		// 	Value: []*armlogic.KeyVaultKey{
		// 		{
		// 			Attributes: &armlogic.KeyVaultKeyAttributes{
		// 				Created: to.Ptr[int64](1498072075),
		// 				Enabled: to.Ptr(true),
		// 				Updated: to.Ptr[int64](1498072075),
		// 			},
		// 			Kid: to.Ptr("https://testKeyVault.vault.azure.net/keys/testkey"),
		// 	}},
		// }
	}
}
