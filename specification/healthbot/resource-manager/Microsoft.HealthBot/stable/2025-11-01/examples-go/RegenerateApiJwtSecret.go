package armhealthbot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot/v2"
)

// Generated from example definition: 2025-11-01/RegenerateApiJwtSecret.json
func ExampleBotsClient_RegenerateAPIJwtSecret() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthbot.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBotsClient().RegenerateAPIJwtSecret(ctx, "healthbotClient", "samplebotname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthbot.BotsClientRegenerateAPIJwtSecretResponse{
	// 	Key: &armhealthbot.Key{
	// 		KeyName: to.Ptr("API_JWT_SECRET"),
	// 		Value: to.Ptr("XXXXX"),
	// 	},
	// }
}
