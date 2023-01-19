using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheLinkedServer_Get.json
// this example is just showing the usage of "LinkedServer_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisResource created on azure
// for more information of creating RedisResource, please refer to the document of RedisResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string name = "cache1";
ResourceIdentifier redisResourceId = RedisResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
RedisResource redis = client.GetRedisResource(redisResourceId);

// get the collection of this RedisLinkedServerWithPropertyResource
RedisLinkedServerWithPropertyCollection collection = redis.GetRedisLinkedServerWithProperties();

// invoke the operation
string linkedServerName = "cache2";
bool result = await collection.ExistsAsync(linkedServerName);

Console.WriteLine($"Succeeded: {result}");
