using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheListPrivateLinkResources.json
// this example is just showing the usage of "PrivateLinkResources_ListByRedisCache" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisResource created on azure
// for more information of creating RedisResource, please refer to the document of RedisResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgtest01";
string cacheName = "cacheTest01";
ResourceIdentifier redisResourceId = RedisResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName);
RedisResource redis = client.GetRedisResource(redisResourceId);

// invoke the operation and iterate over the result
await foreach (RedisPrivateLinkResource item in redis.GetPrivateLinkResourcesByRedisCacheAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
