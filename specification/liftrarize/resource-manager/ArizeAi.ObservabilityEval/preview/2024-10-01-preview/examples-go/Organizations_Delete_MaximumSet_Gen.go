package armarizeaiobservabilityeval_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/arizeaiobservabilityeval/armarizeaiobservabilityeval"
)

// Generated from example definition: 2024-10-01-preview/Organizations_Delete_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armarizeaiobservabilityeval.NewClientFactory("4DEBE8B4-8BA4-42F8-AE50-FBEF318751D1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginDelete(ctx, "rgopenapi", "test-organization-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
