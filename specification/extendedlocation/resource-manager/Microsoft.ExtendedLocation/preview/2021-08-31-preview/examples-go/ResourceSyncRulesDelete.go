package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb9c8e2ca33e9723c2b2fc849f627329067feb54/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/ResourceSyncRulesDelete.json
func ExampleResourceSyncRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewResourceSyncRulesClient().Delete(ctx, "testresourcegroup", "customLocation01", "resourceSyncRule01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
