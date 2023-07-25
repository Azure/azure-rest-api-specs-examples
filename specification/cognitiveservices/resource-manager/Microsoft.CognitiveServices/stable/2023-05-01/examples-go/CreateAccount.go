package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/CreateAccount.json
func ExampleAccountsClient_BeginCreate_createAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "myResourceGroup", "testCreate1", armcognitiveservices.Account{
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr("Emotion"),
		Location: to.Ptr("West US"),
		Properties: &armcognitiveservices.AccountProperties{
			Encryption: &armcognitiveservices.Encryption{
				KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
				KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
					KeyName:     to.Ptr("KeyName"),
					KeyVaultURI: to.Ptr("https://pltfrmscrts-use-pc-dev.vault.azure.net/"),
					KeyVersion:  to.Ptr("891CF236-D241-4738-9462-D506AF493DFA"),
				},
			},
			UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
				{
					ResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
				}},
		},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
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
	// res.Account = armcognitiveservices.Account{
	// 	Name: to.Ptr("testCreate1"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/testCreate1"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T08%3A00%3A05.445595Z'\""),
	// 	Identity: &armcognitiveservices.Identity{
	// 		Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("b5cf119e-a5c2-42c7-802f-592e0efb169f"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Kind: to.Ptr("Emotion"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.AccountProperties{
	// 		Encryption: &armcognitiveservices.Encryption{
	// 			KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
	// 			KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
	// 				KeyName: to.Ptr("FakeKeyName"),
	// 				KeyVaultURI: to.Ptr("https://pltfrmscrts-use-pc-dev.vault.azure.net/"),
	// 				KeyVersion: to.Ptr("891CF236-D241-4738-9462-D506AF493DFA"),
	// 			},
	// 		},
	// 		Endpoint: to.Ptr("https://westus.api.cognitive.microsoft.com/emotion/v1.0"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 		UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
	// 		}},
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// }
}
