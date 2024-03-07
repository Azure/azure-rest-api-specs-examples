package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/UpdateTaskSubscriptionLocation_example.json
func ExampleTasksClient_UpdateSubscriptionLevelTaskState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTasksClient().UpdateSubscriptionLevelTaskState(ctx, "westeurope", "62609ee7-d0a5-8616-9fe4-1df5cca7758d", armsecurity.TaskUpdateActionTypeDismiss, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
