package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DeleteDirectLineSpeechChannel.json
func ExampleChannelsClient_Delete_deleteDirectLineSpeechChannel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewChannelsClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "OneResourceGroupName", "samplebotname", "DirectLineSpeechChannel", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
