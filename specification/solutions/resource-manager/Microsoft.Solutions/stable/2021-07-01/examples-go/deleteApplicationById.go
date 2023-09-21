package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/deleteApplicationById.json
func ExampleApplicationsClient_BeginDeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginDeleteByID(ctx, "subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
