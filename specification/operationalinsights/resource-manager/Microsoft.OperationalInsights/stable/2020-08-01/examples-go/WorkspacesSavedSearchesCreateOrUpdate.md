Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv1.0.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesSavedSearchesCreateOrUpdate.json
func ExampleSavedSearchesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewSavedSearchesClient("00000000-0000-0000-0000-00000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"TestRG",
		"TestWS",
		"00000000-0000-0000-0000-00000000000",
		armoperationalinsights.SavedSearch{
			Properties: &armoperationalinsights.SavedSearchProperties{
				Category:           to.Ptr("Saved Search Test Category"),
				DisplayName:        to.Ptr("Create or Update Saved Search Test"),
				FunctionAlias:      to.Ptr("heartbeat_func"),
				FunctionParameters: to.Ptr("a:int=1"),
				Query:              to.Ptr("Heartbeat | summarize Count() by Computer | take a"),
				Tags: []*armoperationalinsights.Tag{
					{
						Name:  to.Ptr("Group"),
						Value: to.Ptr("Computer"),
					}},
				Version: to.Ptr[int64](2),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
