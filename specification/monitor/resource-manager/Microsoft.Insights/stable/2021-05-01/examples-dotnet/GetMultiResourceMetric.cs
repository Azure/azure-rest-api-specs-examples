using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMultiResourceMetric.json
// this example is just showing the usage of "Metrics_ListAtSubscriptionScopePost" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "92d2a2d8-b514-432d-8cc9-a5f9272630d5";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
SubscriptionResourceGetMonitorMetricsWithPostOptions options = new SubscriptionResourceGetMonitorMetricsWithPostOptions(region: "westus2") { Timespan = "2021-06-08T19:00:00Z/2021-06-12T01:00:00Z", Interval = XmlConvert.ToTimeSpan("PT6H"), Metricnames = "Data Disk Max Burst IOPS", Aggregation = "count", Top = 10, Orderby = "count desc", Filter = "LUN eq '0' and Microsoft.ResourceId eq '*'", Metricnamespace = "microsoft.compute/virtualmachines", AutoAdjustTimegrain = true, ValidateDimensions = false };
await foreach (SubscriptionMonitorMetric item in subscriptionResource.GetMonitorMetricsWithPostAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
