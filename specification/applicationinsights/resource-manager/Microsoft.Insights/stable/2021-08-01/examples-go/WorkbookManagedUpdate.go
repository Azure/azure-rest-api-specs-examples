package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookManagedUpdate.json
func ExampleWorkbooksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkbooksClient("6b643656-33eb-422f-aee8-3ac145d124af", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"my-resource-group",
		"deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
		&armapplicationinsights.WorkbooksClientUpdateOptions{SourceID: to.Ptr("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group"),
			WorkbookUpdateParameters: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
