using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise.Models;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2025-04-01/examples/RedisEnterpriseGetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterprisePrivateEndpointConnectionResource created on azure
// for more information of creating RedisEnterprisePrivateEndpointConnectionResource, please refer to the document of RedisEnterprisePrivateEndpointConnectionResource
string subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
string resourceGroupName = "rg1";
string clusterName = "cache1";
string privateEndpointConnectionName = "pectest01";
ResourceIdentifier redisEnterprisePrivateEndpointConnectionResourceId = RedisEnterprisePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, privateEndpointConnectionName);
RedisEnterprisePrivateEndpointConnectionResource redisEnterprisePrivateEndpointConnection = client.GetRedisEnterprisePrivateEndpointConnectionResource(redisEnterprisePrivateEndpointConnectionResourceId);

// invoke the operation
RedisEnterprisePrivateEndpointConnectionResource result = await redisEnterprisePrivateEndpointConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisEnterprisePrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
