package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemDelete.json
func ExampleAnalyticsItemsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAnalyticsItemsClient().Delete(ctx, "my-resource-group", "my-component", armapplicationinsights.ItemScopePathAnalyticsItems, &armapplicationinsights.AnalyticsItemsClientDeleteOptions{ID: to.Ptr("3466c160-4a10-4df8-afdf-0007f3f6dee5"),
		Name: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
