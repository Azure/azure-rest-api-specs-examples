using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteUpdate.json
// this example is just showing the usage of "Favorites_Update" operation, for the dependent resources, they will have to be created separately.

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
string favoriteId = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
ApplicationInsightsComponentFavorite favoriteProperties = new ApplicationInsightsComponentFavorite()
{
    Name = "Derek Changed This",
    Config = "{\"MEDataModelRawJSON\":\"{\\\"version\\\": \\\"1.4.1\\\",\\\"isCustomDataModel\\\": true,\\\"items\\\": [{\\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\\"chartType\\\": \\\"Area\\\",\\\"chartHeight\\\": 4,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Sum\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"fail\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\\"chartType\\\": \\\"Area\\\",\\\"chartHeight\\\": 2,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/duration\\\",\\\"metricAggregation\\\": \\\"Avg\\\",\\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Avg\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"greenHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\\"chartType\\\": \\\"Bar\\\",\\\"chartHeight\\\": 4,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/duration\\\",\\\"metricAggregation\\\": \\\"Avg\\\",\\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Avg\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"magentaHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\\"grouping\\\": {\\\"kind\\\": \\\"ByDimension\\\",\\\"dimension\\\": \\\"context.application.version\\\"},\\\"chartType\\\": \\\"Grid\\\",\\\"chartHeight\\\": 1,\\\"metrics\\\": [{\\\"id\\\": \\\"basicException.count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"},{\\\"id\\\": \\\"requestFailed.count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"}],\\\"priorPeriod\\\": true,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"percentage\\\": false,\\\"palette\\\": \\\"blueHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"}],\\\"currentFilter\\\": {\\\"eventTypes\\\": [1,2],\\\"typeFacets\\\": {},\\\"isPermissive\\\": false},\\\"timeContext\\\": {\\\"durationMs\\\": 75600000,\\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\\"isInitialTime\\\": false,\\\"grain\\\": 1,\\\"useDashboardTimeRange\\\": false},\\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\\"timeSource\\\": 0}\"}",
    Version = "ME",
    FavoriteType = FavoriteType.Shared,
    SourceType = null,
    Tags =
    {
    "TagSample01","TagSample02","TagSample03"
    },
    Category = null,
    IsGeneratedFromTemplate = false,
};
ApplicationInsightsComponentFavorite result = await applicationInsightsComponent.UpdateFavoriteAsync(favoriteId, favoriteProperties);

Console.WriteLine($"Succeeded: {result}");
