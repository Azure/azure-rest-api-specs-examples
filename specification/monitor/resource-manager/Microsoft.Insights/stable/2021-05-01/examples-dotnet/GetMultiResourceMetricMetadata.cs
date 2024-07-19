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

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-05-01/examples/GetMultiResourceMetricMetadata.json
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
SubscriptionResourceGetMonitorMetricsWithPostOptions options = new SubscriptionResourceGetMonitorMetricsWithPostOptions(region: "westus2") { Timespan = "2021-06-10T02:23:16.129Z/2021-06-12T02:23:16.129Z", Metricnames = "Data Disk Max Burst IOPS", Filter = "LUN eq '0'", Metricnamespace = "microsoft.compute/virtualmachines" };
await foreach (SubscriptionMonitorMetric item in subscriptionResource.GetMonitorMetricsWithPostAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
