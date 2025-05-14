package armlambdatesthyperexecute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/lambdatesthyperexecute/armlambdatesthyperexecute"
)

// Generated from example definition: 2024-02-01-preview/Organizations_Delete_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlambdatesthyperexecute.NewClientFactory("171E7A75-341B-4472-BC4C-7603C5AB9F32", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginDelete(ctx, "rgopenapi", "testorg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
