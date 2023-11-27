using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteAdd.json
// this example is just showing the usage of "Favorites_Add" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationInsightsComponentResource created on azure
// for more information of creating ApplicationInsightsComponentResource, please refer to the document of ApplicationInsightsComponentResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
string resourceName = "my-ai-component";
ResourceIdentifier applicationInsightsComponentResourceId = ApplicationInsightsComponentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ApplicationInsightsComponentResource applicationInsightsComponent = client.GetApplicationInsightsComponentResource(applicationInsightsComponentResourceId);

// invoke the operation
string favoriteId = "deadb33f-8bee-4d3b-a059-9be8dac93960";
ApplicationInsightsComponentFavorite favoriteProperties = new ApplicationInsightsComponentFavorite()
{
    Name = "Blah Blah Blah",
    Config = "{\"MEDataModelRawJSON\":\"{\\n  \\\"version\\\": \\\"1.4.1\\\",\\n  \\\"isCustomDataModel\\\": true,\\n  \\\"items\\\": [\\n    {\\n      \\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Sum\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"fail\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 2,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"greenHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\n      \\\"chartType\\\": \\\"Bar\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"magentaHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\n      \\\"grouping\\\": {\\n        \\\"kind\\\": \\\"ByDimension\\\",\\n        \\\"dimension\\\": \\\"context.application.version\\\"\\n      },\\n      \\\"chartType\\\": \\\"Grid\\\",\\n      \\\"chartHeight\\\": 1,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"basicException.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"\\n        },\\n        {\\n          \\\"id\\\": \\\"requestFailed.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": true,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"blueHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    }\\n  ],\\n  \\\"currentFilter\\\": {\\n    \\\"eventTypes\\\": [\\n      1,\\n      2\\n    ],\\n    \\\"typeFacets\\\": {},\\n    \\\"isPermissive\\\": false\\n  },\\n  \\\"timeContext\\\": {\\n    \\\"durationMs\\\": 75600000,\\n    \\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\n    \\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\n    \\\"isInitialTime\\\": false,\\n    \\\"grain\\\": 1,\\n    \\\"useDashboardTimeRange\\\": false\\n  },\\n  \\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\n  \\\"timeSource\\\": 0\\n}\"}",
    Version = "ME",
    FavoriteType = FavoriteType.Shared,
    SourceType = null,
    Tags =
    {
    "TagSample01","TagSample02"
    },
    Category = null,
    IsGeneratedFromTemplate = false,
};
ApplicationInsightsComponentFavorite result = await applicationInsightsComponent.AddFavoriteAsync(favoriteId, favoriteProperties);

Console.WriteLine($"Succeeded: {result}");
