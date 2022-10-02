package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/CreateWatchlist.json
func ExampleWatchlistsClient_CreateOrUpdate_createsOrUpdatesAWatchlist() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewWatchlistsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "highValueAsset", armsecurityinsights.Watchlist{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.WatchlistProperties{
			Description:    to.Ptr("Watchlist from CSV content"),
			DisplayName:    to.Ptr("High Value Assets Watchlist"),
			ItemsSearchKey: to.Ptr("header1"),
			Provider:       to.Ptr("Microsoft"),
			Source:         to.Ptr("watchlist.csv"),
			SourceType:     to.Ptr(armsecurityinsights.SourceTypeLocalFile),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
