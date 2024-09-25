using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/preview/2024-06-15-preview/examples/ConnectedPartnerResources_List.json
// this example is just showing the usage of "ConnectedPartnerResources_List" operation, for the dependent resources, they will have to be created separately.

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
await foreach (ConnectedPartnerResourcesListFormat item in elasticMonitorResource.GetConnectedPartnerResourcesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
