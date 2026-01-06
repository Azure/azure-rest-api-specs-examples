package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53d56e4ec74156c450d1e51745a971d3f2031dd7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/WorkspacesSavedSearchesGet.json
func ExampleSavedSearchesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavedSearchesClient().Get(ctx, "TestRG", "TestWS", "00000000-0000-0000-0000-00000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavedSearch = armoperationalinsights.SavedSearch{
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000005/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/AtlantisDemo/savedSearches/test-new-saved-search-id-2015"),
	// 	Etag: to.Ptr("W/\"datetime'2017-10-02T23%3A15%3A41.0709875Z'\""),
	// 	Properties: &armoperationalinsights.SavedSearchProperties{
	// 		Category: to.Ptr(" Saved Search Test Category"),
	// 		DisplayName: to.Ptr("Create or Update Saved Search Test"),
	// 		FunctionAlias: to.Ptr("heartbeat_func"),
	// 		FunctionParameters: to.Ptr("a:int=1"),
	// 		Query: to.Ptr("* | measure Count() by Computer | take a"),
	// 		Version: to.Ptr[int64](1),
	// 	},
	// }
}
