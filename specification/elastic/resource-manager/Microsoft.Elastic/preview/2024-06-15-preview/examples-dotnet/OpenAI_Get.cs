using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/preview/2024-06-15-preview/examples/OpenAI_Get.json
// this example is just showing the usage of "OpenAI_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this OpenAIIntegrationRPModelResource
OpenAIIntegrationRPModelCollection collection = elasticMonitorResource.GetOpenAIIntegrationRPModels();

// invoke the operation
string integrationName = "default";
NullableResponse<OpenAIIntegrationRPModelResource> response = await collection.GetIfExistsAsync(integrationName);
OpenAIIntegrationRPModelResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    OpenAIIntegrationRPModelData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
