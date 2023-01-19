using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsMetrics.json
// this example is just showing the usage of "LogAnalytics_GetWafLogAnalyticsMetrics" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResource created on azure
// for more information of creating ProfileResource, please refer to the document of ProfileResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
ResourceIdentifier profileResourceId = ProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
ProfileResource profile = client.GetProfileResource(profileResourceId);

// invoke the operation
IEnumerable<WafMetric> metrics = new WafMetric[]
{
WafMetric.ClientRequestCount
};
DateTimeOffset dateTimeBegin = DateTimeOffset.Parse("2020-11-04T06:49:27.554Z");
DateTimeOffset dateTimeEnd = DateTimeOffset.Parse("2020-11-04T09:49:27.554Z");
WafGranularity granularity = WafGranularity.PT5M;
IEnumerable<WafAction> actions = new WafAction[]
{
WafAction.Block,WafAction.Log
};
WafMetricsResponse result = await profile.GetWafLogAnalyticsMetricsAsync(metrics, dateTimeBegin, dateTimeEnd, granularity, actions: actions);

Console.WriteLine($"Succeeded: {result}");
