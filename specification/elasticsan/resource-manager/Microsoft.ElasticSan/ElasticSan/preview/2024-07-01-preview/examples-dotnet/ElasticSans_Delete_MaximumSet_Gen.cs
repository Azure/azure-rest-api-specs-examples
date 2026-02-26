using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: 2024-07-01-preview/ElasticSans_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "ElasticSan_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanResource created on azure
// for more information of creating ElasticSanResource, please refer to the document of ElasticSanResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
string elasticSanName = "elasticsanname";
ResourceIdentifier elasticSanResourceId = ElasticSanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName);
ElasticSanResource elasticSan = client.GetElasticSanResource(elasticSanResourceId);

// invoke the operation
await elasticSan.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
