using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ExternalSubscriptionForecast.json
// this example is just showing the usage of "Forecast_ExternalCloudProviderUsage" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ExternalCloudProviderType externalCloudProviderType = ExternalCloudProviderType.ExternalSubscriptions;
string externalCloudProviderId = "100";
ForecastDefinition forecastDefinition = new ForecastDefinition(ForecastType.Usage, ForecastTimeframe.Custom, new ForecastDataset(new Dictionary<string, ForecastAggregation>()
{
    ["totalCost"] = new ForecastAggregation(FunctionName.Cost, FunctionType.Sum),
})
{
    Granularity = GranularityType.Daily,
    Filter = new ForecastFilter()
    {
        And =
        {
        new ForecastFilter()
        {
        Or =
        {
        new ForecastFilter()
        {
        Dimensions = new ForecastComparisonExpression("ResourceLocation",ForecastOperatorType.In,new string[]
        {
        "East US","West Europe"
        }),
        },new ForecastFilter()
        {
        Tags = new ForecastComparisonExpression("Environment",ForecastOperatorType.In,new string[]
        {
        "UAT","Prod"
        }),
        }
        },
        },new ForecastFilter()
        {
        Dimensions = new ForecastComparisonExpression("ResourceGroup",ForecastOperatorType.In,new string[]
        {
        "API"
        }),
        }
        },
    },
})
{
    TimePeriod = new ForecastTimePeriod(DateTimeOffset.Parse("2022-08-01T00:00:00+00:00"), DateTimeOffset.Parse("2022-08-31T23:59:59+00:00")),
};
ForecastResult result = await tenantResource.ExternalCloudProviderUsageForecastAsync(externalCloudProviderType, externalCloudProviderId, forecastDefinition);

Console.WriteLine($"Succeeded: {result}");
