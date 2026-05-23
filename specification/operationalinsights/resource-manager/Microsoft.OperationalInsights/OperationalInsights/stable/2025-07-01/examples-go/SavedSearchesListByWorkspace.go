package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/SavedSearchesListByWorkspace.json
func ExampleSavedSearchesClient_ListByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavedSearchesClient().ListByWorkspace(ctx, "TestRG", "TestWS", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoperationalinsights.SavedSearchesClientListByWorkspaceResponse{
	// 	SavedSearchesListResult: armoperationalinsights.SavedSearchesListResult{
	// 		Value: []*armoperationalinsights.SavedSearch{
	// 			{
	// 				Etag: to.Ptr("W/\"datetime'2017-10-02T23%3A15%3A41.0709875Z'\""),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/workspaces/AtlantisDemo/savedSearches/test-new-saved-search-id-2015"),
	// 				Properties: &armoperationalinsights.SavedSearchProperties{
	// 					Category: to.Ptr(" Saved Search Test Category"),
	// 					DisplayName: to.Ptr("Create or Update Saved Search Test"),
	// 					Query: to.Ptr("* | measure Count() by Computer"),
	// 					Tags: []*armoperationalinsights.Tag{
	// 						{
	// 							Name: to.Ptr("Group"),
	// 							Value: to.Ptr("Computer"),
	// 						},
	// 					},
	// 					Version: to.Ptr[int64](1),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
