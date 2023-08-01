using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ElasticSan;
using Azure.ResourceManager.ElasticSan.Models;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/PrivateEndpointConnections_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanPrivateEndpointConnectionResource created on azure
// for more information of creating ElasticSanPrivateEndpointConnectionResource, please refer to the document of ElasticSanPrivateEndpointConnectionResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
string elasticSanName = "elasticsanname";
string privateEndpointConnectionName = "privateendpointconnectionname";
ResourceIdentifier elasticSanPrivateEndpointConnectionResourceId = ElasticSanPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName, privateEndpointConnectionName);
ElasticSanPrivateEndpointConnectionResource elasticSanPrivateEndpointConnection = client.GetElasticSanPrivateEndpointConnectionResource(elasticSanPrivateEndpointConnectionResourceId);

// invoke the operation
await elasticSanPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
