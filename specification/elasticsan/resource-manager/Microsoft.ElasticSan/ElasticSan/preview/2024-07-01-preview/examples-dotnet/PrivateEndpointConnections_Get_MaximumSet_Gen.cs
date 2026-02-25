using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: 2024-07-01-preview/PrivateEndpointConnections_Get_MaximumSet_Gen.json
// this example is just showing the usage of "PrivateEndpointConnection_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ElasticSanPrivateEndpointConnectionResource
ElasticSanPrivateEndpointConnectionCollection collection = elasticSan.GetElasticSanPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "privateendpointconnectionname";
NullableResponse<ElasticSanPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
ElasticSanPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ElasticSanPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
