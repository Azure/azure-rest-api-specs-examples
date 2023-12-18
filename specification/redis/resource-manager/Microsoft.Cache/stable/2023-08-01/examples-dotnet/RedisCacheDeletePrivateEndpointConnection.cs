using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheDeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisPrivateEndpointConnectionResource created on azure
// for more information of creating RedisPrivateEndpointConnectionResource, please refer to the document of RedisPrivateEndpointConnectionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rgtest01";
string cacheName = "cachetest01";
string privateEndpointConnectionName = "pectest01";
ResourceIdentifier redisPrivateEndpointConnectionResourceId = RedisPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName, privateEndpointConnectionName);
RedisPrivateEndpointConnectionResource redisPrivateEndpointConnection = client.GetRedisPrivateEndpointConnectionResource(redisPrivateEndpointConnectionResourceId);

// invoke the operation
await redisPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
