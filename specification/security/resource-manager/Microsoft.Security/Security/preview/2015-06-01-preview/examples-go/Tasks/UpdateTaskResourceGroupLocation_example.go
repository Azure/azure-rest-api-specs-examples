package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2015-06-01-preview/Tasks/UpdateTaskResourceGroupLocation_example.json
func ExampleTasksClient_UpdateResourceGroupLevelTaskState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTasksClient().UpdateResourceGroupLevelTaskState(ctx, "myRg", "westeurope", "d55b4dc0-779c-c66c-33e5-d7bce24c4222", armsecurity.TaskUpdateActionTypeDismiss, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
