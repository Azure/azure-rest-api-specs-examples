using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheGetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

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
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
