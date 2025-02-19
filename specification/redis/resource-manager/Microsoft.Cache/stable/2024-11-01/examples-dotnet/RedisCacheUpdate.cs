using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Redis.Models;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheUpdate.json
// this example is just showing the usage of "Redis_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
RedisPatch patch = new RedisPatch
{
    EnableNonSslPort = true,
    ReplicasPerPrimary = 2,
};
ArmOperation<RedisResource> lro = await redis.UpdateAsync(WaitUntil.Completed, patch);
RedisResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
