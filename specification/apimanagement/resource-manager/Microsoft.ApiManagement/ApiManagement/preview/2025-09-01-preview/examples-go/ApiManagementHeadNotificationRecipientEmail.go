package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementHeadNotificationRecipientEmail.json
func ExampleNotificationRecipientEmailClient_CheckEntityExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewNotificationRecipientEmailClient().CheckEntityExists(ctx, "rg1", "apimService1", armapimanagement.NotificationNameRequestPublisherNotificationMessage, "contoso@live.com", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
