package armagricultureplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agricultureplatform/armagricultureplatform"
)

// Generated from example definition: 2024-06-01-preview/AgriService_Delete_MaximumSet_Gen.json
func ExampleAgriServiceClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagricultureplatform.NewClientFactory("83D293F5-DEFD-4D48-B120-1DC713BE338A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgriServiceClient().BeginDelete(ctx, "rgopenapi", "abc123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
