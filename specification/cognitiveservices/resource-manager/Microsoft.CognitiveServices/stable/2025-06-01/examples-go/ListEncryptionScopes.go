package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/ListEncryptionScopes.json
func ExampleEncryptionScopesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEncryptionScopesClient().NewListPager("resourceGroupName", "accountName", nil)
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
		// page.EncryptionScopeListResult = armcognitiveservices.EncryptionScopeListResult{
		// 	Value: []*armcognitiveservices.EncryptionScope{
		// 		{
		// 			Name: to.Ptr("encryptionScopeName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/encryptionScopes"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/encryptionScopes/encryptionScopeName"),
		// 			Etag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armcognitiveservices.EncryptionScopeProperties{
		// 				KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
		// 				KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
		// 					IdentityClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					KeyName: to.Ptr("DevKeyWestUS2"),
		// 					KeyVaultURI: to.Ptr("https://devkvwestus2.vault.azure.net/"),
		// 					KeyVersion: to.Ptr("9f85549d7bf14ff4bf178c10d3bdca95"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcognitiveservices.EncryptionScopeProvisioningStateSucceeded),
		// 				State: to.Ptr(armcognitiveservices.EncryptionScopeStateEnabled),
		// 			},
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T06:35:08.066Z"); return t}()),
		// 				CreatedBy: to.Ptr("xxx@microsoft.com"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T06:35:08.066Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("xxx@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
