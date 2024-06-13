package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlistItemById.json
func ExampleWatchlistItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistItemsClient().Get(ctx, "myRg", "myWorkspace", "highValueAsset", "3f8901fe-63d9-4875-9ad5-9fb3b8105797", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WatchlistItem = armsecurityinsights.WatchlistItem{
	// 	Name: to.Ptr("fd37d325-7090-47fe-851a-5b5a00c3f576"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists/WatchlistItems"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Watchlists/highValueAsset/WatchlistItems/fd37d325-7090-47fe-851a-5b5a00c3f576"),
	// 	Etag: to.Ptr("\"f2089bfa-0000-0d00-0000-601c58b42021\""),
	// 	Properties: &armsecurityinsights.WatchlistItemProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-04T20:27:32.378Z"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		EntityMapping: map[string]any{
	// 		},
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsKeyValue: map[string]any{
	// 			"Header-1": "v1_1",
	// 			"Header-2": "v1_2",
	// 			"Header-3": "v1_3",
	// 			"Header-4": "v1_4",
	// 			"Header-5": "v1_5",
	// 		},
	// 		TenantID: to.Ptr("3f8901fe-63d9-4875-9ad5-9fb3b8105797"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-04T20:27:32.378Z"); return t}()),
	// 		UpdatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		WatchlistItemID: to.Ptr("fd37d325-7090-47fe-851a-5b5a00c3f576"),
	// 		WatchlistItemType: to.Ptr("watchlist-item"),
	// 	},
	// }
}
