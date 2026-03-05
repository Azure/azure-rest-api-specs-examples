package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/GetSubscriptionOperationWithAsyncResponse.json
func ExampleGlobalClient_GetSubscriptionOperationWithAsyncResponse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGlobalClient().GetSubscriptionOperationWithAsyncResponse(ctx, "West US", "34adfa4f-cedf-4dc0-ba29-b6d1a69ab5d5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
