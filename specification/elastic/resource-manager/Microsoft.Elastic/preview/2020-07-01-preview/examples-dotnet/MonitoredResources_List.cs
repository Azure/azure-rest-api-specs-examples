using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Elastic;
using Azure.ResourceManager.Elastic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/preview/2020-07-01-preview/examples/MonitoredResources_List.json
// this example is just showing the usage of "MonitoredResources_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticMonitorResource created on azure
// for more information of creating ElasticMonitorResource, please refer to the document of ElasticMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
ResourceIdentifier elasticMonitorResourceId = ElasticMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
ElasticMonitorResource elasticMonitorResource = client.GetElasticMonitorResource(elasticMonitorResourceId);

// invoke the operation and iterate over the result
await foreach (MonitoredResourceContent item in elasticMonitorResource.GetMonitoredResourcesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
