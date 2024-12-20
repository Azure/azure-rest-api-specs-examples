using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Redis.Models;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCachePutPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisResource created on azure
// for more information of creating RedisResource, please refer to the document of RedisResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rgtest01";
string cacheName = "cachetest01";
ResourceIdentifier redisResourceId = RedisResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName);
RedisResource redis = client.GetRedisResource(redisResourceId);

// get the collection of this RedisPrivateEndpointConnectionResource
RedisPrivateEndpointConnectionCollection collection = redis.GetRedisPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "pectest01";
RedisPrivateEndpointConnectionData data = new RedisPrivateEndpointConnectionData
{
    RedisPrivateLinkServiceConnectionState = new RedisPrivateLinkServiceConnectionState
    {
        Status = RedisPrivateEndpointServiceConnectionStatus.Approved,
        Description = "Auto-Approved",
    },
};
ArmOperation<RedisPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
RedisPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
