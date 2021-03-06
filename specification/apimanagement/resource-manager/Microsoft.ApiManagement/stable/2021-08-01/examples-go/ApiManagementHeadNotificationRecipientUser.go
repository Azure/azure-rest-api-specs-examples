package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadNotificationRecipientUser.json
func ExampleNotificationRecipientUserClient_CheckEntityExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewNotificationRecipientUserClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CheckEntityExists(ctx,
		"rg1",
		"apimService1",
		armapimanagement.NotificationNameRequestPublisherNotificationMessage,
		"576823d0a40f7e74ec07d642",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
