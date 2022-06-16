package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/watchlists/CreateWatchlistItem.json
func ExampleWatchlistItemsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewWatchlistItemsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<watchlist-alias>",
		"<watchlist-item-id>",
		armsecurityinsight.WatchlistItem{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsight.WatchlistItemProperties{
				ItemsKeyValue: map[string]interface{}{
					"Business tier":  "10.0.2.0/24",
					"Data tier":      "10.0.2.0/24",
					"Gateway subnet": "10.0.255.224/27",
					"Private DMZ in": "10.0.0.0/27",
					"Public DMZ out": "10.0.0.96/27",
					"Web Tier":       "10.0.1.0/24",
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WatchlistItemsClientCreateOrUpdateResult)
}
