package armlocks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_DeleteAtScope.json
func ExampleManagementLocksClient_DeleteByScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlocks.NewManagementLocksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.DeleteByScope(ctx,
		"subscriptions/subscriptionId",
		"testlock",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
