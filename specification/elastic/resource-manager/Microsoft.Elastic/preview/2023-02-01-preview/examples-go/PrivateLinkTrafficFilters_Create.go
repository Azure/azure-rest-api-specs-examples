package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dbd896bc9a795bcb3ec7db0a340b517fd3059620/specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/PrivateLinkTrafficFilters_Create.json
func ExampleCreateAndAssociatePLFilterClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCreateAndAssociatePLFilterClient().BeginCreate(ctx, "myResourceGroup", "myMonitor", &armelastic.CreateAndAssociatePLFilterClientBeginCreateOptions{Name: nil,
		PrivateEndpointGUID: to.Ptr("fdb54d3b-e85e-4d08-8958-0d2f7g523df9"),
		PrivateEndpointName: to.Ptr("myPrivateEndpoint"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
