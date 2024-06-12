package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlistByAlias.json
func ExampleWatchlistsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistsClient().Get(ctx, "myRg", "myWorkspace", "highValueAsset", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Watchlist = armsecurityinsights.Watchlist{
	// 	Name: to.Ptr("highValueAsset"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/watchlists/highValueAsset"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.WatchlistProperties{
	// 		Description: to.Ptr("Watchlist from CSV content"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:54.774Z"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		DefaultDuration: to.Ptr("P1279DT12H30M5S"),
	// 		DisplayName: to.Ptr("High Value Assets Watchlist"),
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsSearchKey: to.Ptr("header1"),
	// 		Labels: []*string{
	// 			to.Ptr("Tag1"),
	// 			to.Ptr("Tag2")},
	// 			Provider: to.Ptr("Microsoft"),
	// 			Source: to.Ptr("watchlist.csv"),
	// 			SourceType: to.Ptr(armsecurityinsights.SourceTypeLocalFile),
	// 			TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 			Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-28T00:26:57.000Z"); return t}()),
	// 			UpdatedBy: &armsecurityinsights.UserInfo{
	// 				Name: to.Ptr("john doe"),
	// 				Email: to.Ptr("john@contoso.com"),
	// 				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 			},
	// 			WatchlistAlias: to.Ptr("highValueAsset"),
	// 			WatchlistID: to.Ptr("76d5a51f-ba1f-4038-9d22-59fda38dc017"),
	// 			WatchlistType: to.Ptr("watchlist"),
	// 		},
	// 	}
}
