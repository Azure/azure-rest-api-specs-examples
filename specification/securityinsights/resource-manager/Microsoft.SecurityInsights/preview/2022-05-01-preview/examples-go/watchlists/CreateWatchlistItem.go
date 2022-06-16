package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/watchlists/CreateWatchlistItem.json
func ExampleWatchlistItemsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewWatchlistItemsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRg",
		"myWorkspace",
		"highValueAsset",
		"82ba292c-dc97-4dfc-969d-d4dd9e666842",
		armsecurityinsights.WatchlistItem{
			Etag: to.Ptr("0300bf09-0000-0000-0000-5c37296e0000"),
			Properties: &armsecurityinsights.WatchlistItemProperties{
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
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
