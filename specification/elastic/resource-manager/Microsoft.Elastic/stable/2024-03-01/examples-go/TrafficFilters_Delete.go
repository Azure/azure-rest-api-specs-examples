package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/TrafficFilters_Delete.json
func ExampleTrafficFiltersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTrafficFiltersClient().Delete(ctx, "myResourceGroup", "myMonitor", &armelastic.TrafficFiltersClientDeleteOptions{RulesetID: to.Ptr("31d91b5afb6f4c2eaaf104c97b1991dd")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
