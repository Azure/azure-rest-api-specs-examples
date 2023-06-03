using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RedisEnterprise;
using Azure.ResourceManager.RedisEnterprise.Models;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseGetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterpriseClusterResource created on azure
// for more information of creating RedisEnterpriseClusterResource, please refer to the document of RedisEnterpriseClusterResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string clusterName = "cache1";
ResourceIdentifier redisEnterpriseClusterResourceId = RedisEnterpriseClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
RedisEnterpriseClusterResource redisEnterpriseCluster = client.GetRedisEnterpriseClusterResource(redisEnterpriseClusterResourceId);

// get the collection of this RedisEnterprisePrivateEndpointConnectionResource
RedisEnterprisePrivateEndpointConnectionCollection collection = redisEnterpriseCluster.GetRedisEnterprisePrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "pectest01";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
