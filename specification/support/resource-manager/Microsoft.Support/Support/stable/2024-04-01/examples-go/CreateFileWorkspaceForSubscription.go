package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: 2024-04-01/CreateFileWorkspaceForSubscription.json
func ExampleFileWorkspacesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("132d901f-189d-4381-9214-fe68e27e05a1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFileWorkspacesClient().Create(ctx, "testworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
