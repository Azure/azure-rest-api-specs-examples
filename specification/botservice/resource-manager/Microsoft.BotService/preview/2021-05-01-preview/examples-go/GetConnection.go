package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetConnection.json
func ExampleBotConnectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewBotConnectionClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"OneResourceGroupName",
		"samplebotname",
		"sampleConnection",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
