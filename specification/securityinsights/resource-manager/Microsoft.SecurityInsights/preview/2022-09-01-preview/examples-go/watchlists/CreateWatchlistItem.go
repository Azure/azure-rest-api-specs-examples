package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/CreateWatchlistItem.json
func ExampleWatchlistItemsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistItemsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "highValueAsset", "82ba292c-dc97-4dfc-969d-d4dd9e666842", armsecurityinsights.WatchlistItem{
		Etag: to.Ptr("0300bf09-0000-0000-0000-5c37296e0000"),
		Properties: &armsecurityinsights.WatchlistItemProperties{
			ItemsKeyValue: map[string]any{
				"Business tier":  "10.0.2.0/24",
				"Data tier":      "10.0.2.0/24",
				"Gateway subnet": "10.0.255.224/27",
				"Private DMZ in": "10.0.0.0/27",
				"Public DMZ out": "10.0.0.96/27",
				"Web Tier":       "10.0.1.0/24",
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WatchlistItem = armsecurityinsights.WatchlistItem{
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists/WatchlistItems"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Watchlists/highValueAsset/WatchlistItems/82ba292c-dc97-4dfc-969d-d4dd9e666842"),
	// 	Etag: to.Ptr("0300bf09-0000-0000-0000-5c37296e0000"),
	// 	Properties: &armsecurityinsights.WatchlistItemProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-15T04:58:56.0748363+00:00"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsKeyValue: map[string]any{
	// 			"Business tier": "10.0.2.0/24",
	// 			"Data tier": "10.0.2.0/24",
	// 			"Gateway subnet": "10.0.255.224/27",
	// 			"Private DMZ in": "10.0.0.0/27",
	// 			"Public DMZ out": "10.0.0.96/27",
	// 			"Web Tier": "10.0.1.0/24",
	// 		},
	// 		TenantID: to.Ptr("4008512e-1d30-48b2-9ee2-d3612ed9d3ea"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-16T16:05:20+00:00"); return t}()),
	// 		UpdatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		WatchlistItemID: to.Ptr("82ba292c-dc97-4dfc-969d-d4dd9e666842"),
	// 		WatchlistItemType: to.Ptr("watchlist-item"),
	// 	},
	// }
}
