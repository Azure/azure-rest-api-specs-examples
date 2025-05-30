using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Redis.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheCreateDefaultVersion.json
// this example is just showing the usage of "Redis_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this RedisResource
RedisCollection collection = resourceGroupResource.GetAllRedis();

// invoke the operation
string name = "cache1";
RedisCreateOrUpdateContent content = new RedisCreateOrUpdateContent(new AzureLocation("East US"), new RedisSku(RedisSkuName.Premium, RedisSkuFamily.Premium, 1))
{
    Zones = { "1" },
    RedisConfiguration = new RedisCommonConfiguration
    {
        MaxMemoryPolicy = "allkeys-lru",
    },
    EnableNonSslPort = true,
    ReplicasPerPrimary = 2,
    ShardCount = 2,
    MinimumTlsVersion = RedisTlsVersion.Tls1_2,
    SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/network1/subnets/subnet1"),
    StaticIP = IPAddress.Parse("192.168.0.5"),
};
ArmOperation<RedisResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, content);
RedisResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
