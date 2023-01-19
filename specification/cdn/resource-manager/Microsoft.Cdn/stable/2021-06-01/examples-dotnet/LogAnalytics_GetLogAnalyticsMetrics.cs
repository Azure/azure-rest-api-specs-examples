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

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsMetrics.json
// this example is just showing the usage of "LogAnalytics_GetLogAnalyticsMetrics" operation, for the dependent resources, they will have to be created separately.

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
IEnumerable<LogMetric> metrics = new LogMetric[]
{
LogMetric.ClientRequestCount
};
DateTimeOffset dateTimeBegin = DateTimeOffset.Parse("2020-11-04T04:30:00.000Z");
DateTimeOffset dateTimeEnd = DateTimeOffset.Parse("2020-11-04T05:00:00.000Z");
LogMetricsGranularity granularity = LogMetricsGranularity.PT5M;
IEnumerable<string> customDomains = new string[]
{
"customdomain1.azurecdn.net","customdomain2.azurecdn.net"
};
IEnumerable<string> protocols = new string[]
{
"https"
};
IEnumerable<LogMetricsGroupBy> groupBy = new LogMetricsGroupBy[]
{
LogMetricsGroupBy.Protocol
};
MetricsResponse result = await profile.GetLogAnalyticsMetricsAsync(metrics, dateTimeBegin, dateTimeEnd, granularity, customDomains, protocols, groupBy: groupBy);

Console.WriteLine($"Succeeded: {result}");
