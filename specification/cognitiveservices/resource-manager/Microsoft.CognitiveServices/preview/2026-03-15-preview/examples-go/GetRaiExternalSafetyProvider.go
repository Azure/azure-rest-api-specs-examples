package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/GetRaiExternalSafetyProvider.json
func ExampleRaiExternalSafetyProviderClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiExternalSafetyProviderClient().Get(ctx, "safetyProviderName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiExternalSafetyProviderClientGetResponse{
	// 	RaiExternalSafetyProviderSchema: armcognitiveservices.RaiExternalSafetyProviderSchema{
	// 		Name: to.Ptr("safetyProviderName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CognitiveServices/raiExternalSafetyProviders/safetyProviderName"),
	// 		Properties: &armcognitiveservices.RaiExternalSafetyProviderSchemaProperties{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-01T00:00:00Z"); return t}()),
	// 			KeyVaultURI: to.Ptr("https://example.vault.azure.net"),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-02T00:00:00Z"); return t}()),
	// 			ManagedIdentity: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Mode: to.Ptr("sync"),
	// 			ProviderID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ProviderName: to.Ptr("safetyProviderName"),
	// 			SecretName: to.Ptr("mySecretName"),
	// 			URL: to.Ptr("https://example.webhook.endpoint"),
	// 		},
	// 	},
	// }
}
