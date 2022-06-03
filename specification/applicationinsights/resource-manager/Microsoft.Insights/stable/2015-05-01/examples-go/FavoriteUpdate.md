Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv1.0.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteUpdate.json
func ExampleFavoritesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"my-resource-group",
		"my-ai-component",
		"deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
		armapplicationinsights.ComponentFavorite{
			Config:                  to.Ptr("{\"MEDataModelRawJSON\":\"{\\\"version\\\": \\\"1.4.1\\\",\\\"isCustomDataModel\\\": true,\\\"items\\\": [{\\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\\"chartType\\\": \\\"Area\\\",\\\"chartHeight\\\": 4,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Sum\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"fail\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\\"chartType\\\": \\\"Area\\\",\\\"chartHeight\\\": 2,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/duration\\\",\\\"metricAggregation\\\": \\\"Avg\\\",\\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Avg\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"greenHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\\"chartType\\\": \\\"Bar\\\",\\\"chartHeight\\\": 4,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/duration\\\",\\\"metricAggregation\\\": \\\"Avg\\\",\\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Avg\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"magentaHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\\"grouping\\\": {\\\"kind\\\": \\\"ByDimension\\\",\\\"dimension\\\": \\\"context.application.version\\\"},\\\"chartType\\\": \\\"Grid\\\",\\\"chartHeight\\\": 1,\\\"metrics\\\": [{\\\"id\\\": \\\"basicException.count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"},{\\\"id\\\": \\\"requestFailed.count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"}],\\\"priorPeriod\\\": true,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"percentage\\\": false,\\\"palette\\\": \\\"blueHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"}],\\\"currentFilter\\\": {\\\"eventTypes\\\": [1,2],\\\"typeFacets\\\": {},\\\"isPermissive\\\": false},\\\"timeContext\\\": {\\\"durationMs\\\": 75600000,\\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\\"isInitialTime\\\": false,\\\"grain\\\": 1,\\\"useDashboardTimeRange\\\": false},\\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\\"timeSource\\\": 0}\"}"),
			FavoriteID:              to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
			FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
			IsGeneratedFromTemplate: to.Ptr(false),
			Name:                    to.Ptr("Derek Changed This"),
			Tags: []*string{
				to.Ptr("TagSample01"),
				to.Ptr("TagSample02"),
				to.Ptr("TagSample03")},
			TimeModified: to.Ptr("2018-02-02T18:39:11.6569686Z"),
			Version:      to.Ptr("ME"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
