using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Elastic;
using Azure.ResourceManager.Elastic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/preview/2020-07-01-preview/examples/Monitors_Update.json
// this example is just showing the usage of "Monitors_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
ElasticMonitorResourcePatch patch = new ElasticMonitorResourcePatch();
ElasticMonitorResource result = await elasticMonitorResource.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticMonitorResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
