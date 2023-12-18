using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCachePatchSchedulesCreateOrUpdate.json
// this example is just showing the usage of "PatchSchedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this RedisPatchScheduleResource
RedisPatchScheduleCollection collection = redis.GetRedisPatchSchedules();

// invoke the operation
RedisPatchScheduleDefaultName defaultName = RedisPatchScheduleDefaultName.Default;
RedisPatchScheduleData data = new RedisPatchScheduleData(new RedisPatchScheduleSetting[]
{
new RedisPatchScheduleSetting(RedisDayOfWeek.Monday,12)
{
MaintenanceWindow = XmlConvert.ToTimeSpan("PT5H"),
},new RedisPatchScheduleSetting(RedisDayOfWeek.Tuesday,12)
});
ArmOperation<RedisPatchScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, defaultName, data);
RedisPatchScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisPatchScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
