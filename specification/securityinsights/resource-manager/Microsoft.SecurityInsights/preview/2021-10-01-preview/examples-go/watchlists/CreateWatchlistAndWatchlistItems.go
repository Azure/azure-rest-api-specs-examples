package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/watchlists/CreateWatchlistAndWatchlistItems.json
func ExampleWatchlistsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewWatchlistsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<watchlist-alias>",
		armsecurityinsights.Watchlist{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsights.WatchlistProperties{
				Description:         to.StringPtr("<description>"),
				ContentType:         to.StringPtr("<content-type>"),
				DisplayName:         to.StringPtr("<display-name>"),
				ItemsSearchKey:      to.StringPtr("<items-search-key>"),
				NumberOfLinesToSkip: to.Int32Ptr(1),
				Provider:            to.StringPtr("<provider>"),
				RawContent:          to.StringPtr("<raw-content>"),
				Source:              armsecurityinsights.Source("Local file").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WatchlistsClientCreateOrUpdateResult)
}
