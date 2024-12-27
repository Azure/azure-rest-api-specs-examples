using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic.Models;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/OpenAI_CreateOrUpdate.json
// this example is just showing the usage of "OpenAI_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticOpenAIIntegrationResource created on azure
// for more information of creating ElasticOpenAIIntegrationResource, please refer to the document of ElasticOpenAIIntegrationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string integrationName = "default";
ResourceIdentifier elasticOpenAIIntegrationResourceId = ElasticOpenAIIntegrationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, integrationName);
ElasticOpenAIIntegrationResource elasticOpenAIIntegration = client.GetElasticOpenAIIntegrationResource(elasticOpenAIIntegrationResourceId);

// invoke the operation
ElasticOpenAIIntegrationData data = new ElasticOpenAIIntegrationData();
ArmOperation<ElasticOpenAIIntegrationResource> lro = await elasticOpenAIIntegration.UpdateAsync(WaitUntil.Completed, data);
ElasticOpenAIIntegrationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticOpenAIIntegrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
