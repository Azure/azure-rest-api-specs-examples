package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/GetEncryptionScope.json
func ExampleEncryptionScopesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Get(ctx, "resourceGroupName", "accountName", "encryptionScopeName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.EncryptionScopesClientGetResponse{
	// 	EncryptionScope: armcognitiveservices.EncryptionScope{
	// 		Name: to.Ptr("encryptionScopeName"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/encryptionScopes"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/encryptionScopes/encryptionScopeName"),
	// 		Properties: &armcognitiveservices.EncryptionScopeProperties{
	// 			KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
	// 			KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
	// 				IdentityClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				KeyName: to.Ptr("DevKeyWestUS2"),
	// 				KeyVaultURI: to.Ptr("https://devkvwestus2.vault.azure.net/"),
	// 				KeyVersion: to.Ptr("9f85549d7bf14ff4bf178c10d3bdca95"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.EncryptionScopeProvisioningStateSucceeded),
	// 			State: to.Ptr(armcognitiveservices.EncryptionScopeStateEnabled),
	// 		},
	// 		SystemData: &armcognitiveservices.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T06:35:08.0662558Z"); return t}()),
	// 			CreatedBy: to.Ptr("xxx@microsoft.com"),
	// 			CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T06:35:08.0662558Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xxx@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
