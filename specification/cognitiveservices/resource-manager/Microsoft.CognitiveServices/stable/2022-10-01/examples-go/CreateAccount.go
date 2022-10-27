package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/CreateAccount.json
func ExampleAccountsClient_BeginCreate_createAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewAccountsClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myResourceGroup", "testCreate1", armcognitiveservices.Account{
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
	// TODO: use response item
	_ = res
}
