using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheLinkedServer_Delete.json
// this example is just showing the usage of "LinkedServer_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisLinkedServerWithPropertyResource created on azure
// for more information of creating RedisLinkedServerWithPropertyResource, please refer to the document of RedisLinkedServerWithPropertyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string name = "cache1";
string linkedServerName = "cache2";
ResourceIdentifier redisLinkedServerWithPropertyResourceId = RedisLinkedServerWithPropertyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, linkedServerName);
RedisLinkedServerWithPropertyResource redisLinkedServerWithProperty = client.GetRedisLinkedServerWithPropertyResource(redisLinkedServerWithPropertyResourceId);

// invoke the operation
await redisLinkedServerWithProperty.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
