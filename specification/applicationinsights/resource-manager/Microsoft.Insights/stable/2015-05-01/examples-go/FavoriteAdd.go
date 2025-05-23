package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteAdd.json
func ExampleFavoritesClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFavoritesClient().Add(ctx, "my-resource-group", "my-ai-component", "deadb33f-8bee-4d3b-a059-9be8dac93960", armapplicationinsights.ComponentFavorite{
		Config:                  to.Ptr("{\"MEDataModelRawJSON\":\"{\\n  \\\"version\\\": \\\"1.4.1\\\",\\n  \\\"isCustomDataModel\\\": true,\\n  \\\"items\\\": [\\n    {\\n      \\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Sum\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"fail\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 2,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"greenHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\n      \\\"chartType\\\": \\\"Bar\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"magentaHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\n      \\\"grouping\\\": {\\n        \\\"kind\\\": \\\"ByDimension\\\",\\n        \\\"dimension\\\": \\\"context.application.version\\\"\\n      },\\n      \\\"chartType\\\": \\\"Grid\\\",\\n      \\\"chartHeight\\\": 1,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"basicException.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"\\n        },\\n        {\\n          \\\"id\\\": \\\"requestFailed.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": true,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"blueHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    }\\n  ],\\n  \\\"currentFilter\\\": {\\n    \\\"eventTypes\\\": [\\n      1,\\n      2\\n    ],\\n    \\\"typeFacets\\\": {},\\n    \\\"isPermissive\\\": false\\n  },\\n  \\\"timeContext\\\": {\\n    \\\"durationMs\\\": 75600000,\\n    \\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\n    \\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\n    \\\"isInitialTime\\\": false,\\n    \\\"grain\\\": 1,\\n    \\\"useDashboardTimeRange\\\": false\\n  },\\n  \\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\n  \\\"timeSource\\\": 0\\n}\"}"),
		FavoriteID:              to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
		FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
		IsGeneratedFromTemplate: to.Ptr(false),
		Name:                    to.Ptr("Blah Blah Blah"),
		Tags: []*string{
			to.Ptr("TagSample01"),
			to.Ptr("TagSample02")},
		Version: to.Ptr("ME"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentFavorite = armapplicationinsights.ComponentFavorite{
	// 	Category: to.Ptr(""),
	// 	Config: to.Ptr("{\"MEDataModelRawJSON\":{\n  \"version\": \"1.4.1\",\n  \"isCustomDataModel\": true,\n  \"items\": [\n    {\n      \"id\": \"90a7134d-9a38-4c25-88d3-a495209873eb\",\n      \"chartType\": \"Area\",\n \"chartHeight\": 4,\n      \"metrics\": [\n        {\n          \"id\": \"preview/requests/count\",\n          \"metricAggregation\": \"Sum\",\n          \"color\": \"msportalfx-bgcolor-d0\"\n        }\n],\n      \"priorPeriod\": false,\n      \"clickAction\": {\n        \"defaultBlade\": \"SearchBlade\"\n      },\n      \"horizontalBars\": true,\n      \"showOther\": true,\n      \"aggregation\": \"Sum\",\n \"percentage\": false,\n      \"palette\": \"fail\",\n      \"yAxisOption\": 0,\n      \"title\": \"\"\n    },\n    {\n      \"id\": \"0c289098-88e8-4010-b212-546815cddf70\",\n      \"chartType\": \"Area\",\n      \"chartHeight\": 2,\n      \"metrics\": [\n        {\n          \"id\": \"preview/requests/duration\",\n          \"metricAggregation\": \"Avg\",\n          \"color\": \"msportalfx-bgcolor-j1\"\n        }\n      ],\n      \"priorPeriod\": false,\n      \"clickAction\": {\n        \"defaultBlade\": \"SearchBlade\"\n      },\n      \"horizontalBars\": true,\n \"showOther\": true,\n      \"aggregation\": \"Avg\",\n      \"percentage\": false,\n      \"palette\": \"greenHues\",\n      \"yAxisOption\": 0,\n      \"title\": \"\"\n    },\n    {\n      \"id\": \"cbdaab6f-a808-4f71-aca5-b3976cbb7345\",\n      \"chartType\": \"Bar\",\n      \"chartHeight\": 4,\n      \"metrics\": [\n        {\n          \"id\": \"preview/requests/duration\",\n \"metricAggregation\": \"Avg\",\n          \"color\": \"msportalfx-bgcolor-d0\"\n        }\n      ],\n      \"priorPeriod\": false,\n      \"clickAction\": {\n        \"defaultBlade\": \"SearchBlade\"\n },\n      \"horizontalBars\": true,\n      \"showOther\": true,\n      \"aggregation\": \"Avg\",\n      \"percentage\": false,\n      \"palette\": \"magentaHues\",\n      \"yAxisOption\": 0,\n      \"title\": \"\"\n    },\n    {\n      \"id\": \"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\",\n      \"grouping\": {\n        \"kind\": \"ByDimension\",\n        \"dimension\": \"context.application.version\"\n      },\n \"chartType\": \"Grid\",\n      \"chartHeight\": 1,\n      \"metrics\": [\n        {\n          \"id\": \"basicException.count\",\n          \"metricAggregation\": \"Sum\",\n          \"color\": \"msportalfx-bgcolor-g0\"\n        },\n        {\n          \"id\": \"requestFailed.count\",\n          \"metricAggregation\": \"Sum\",\n          \"color\": \"msportalfx-bgcolor-f0s2\"\n        }\n      ],\n \"priorPeriod\": true,\n      \"clickAction\": {\n        \"defaultBlade\": \"SearchBlade\"\n      },\n      \"horizontalBars\": true,\n      \"showOther\": true,\n      \"percentage\": false,\n \"palette\": \"blueHues\",\n      \"yAxisOption\": 0,\n      \"title\": \"\"\n    }\n  ],\n  \"currentFilter\": {\n    \"eventTypes\": [\n      1,\n      2\n    ],\n    \"typeFacets\": {},\n \"isPermissive\": false\n  },\n  \"timeContext\": {\n    \"durationMs\": 75600000,\n    \"endTime\": \"2018-01-31T20:30:00.000Z\",\n    \"createdTime\": \"2018-01-31T23:54:26.280Z\",\n    \"isInitialTime\": false,\n    \"grain\": 1,\n    \"useDashboardTimeRange\": false\n  },\n  \"jsonUri\": \"Favorite_BlankChart\",\n  \"timeSource\": 0\n}\"}"),
	// 	FavoriteID: to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
	// 	FavoriteType: to.Ptr(armapplicationinsights.FavoriteTypeShared),
	// 	IsGeneratedFromTemplate: to.Ptr(false),
	// 	Name: to.Ptr("Blah Blah Blah"),
	// 	SourceType: to.Ptr(""),
	// 	Tags: []*string{
	// 		to.Ptr("TagSample01"),
	// 		to.Ptr("TagSample02")},
	// 		TimeModified: to.Ptr("2018-02-02T23:18:32.1850959Z"),
	// 		Version: to.Ptr("ME"),
	// 	}
}
