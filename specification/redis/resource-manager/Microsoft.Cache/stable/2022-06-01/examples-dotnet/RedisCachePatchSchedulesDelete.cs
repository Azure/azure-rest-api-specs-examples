using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;
using Azure.ResourceManager.Redis.Models;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCachePatchSchedulesDelete.json
// this example is just showing the usage of "PatchSchedules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisPatchScheduleResource created on azure
// for more information of creating RedisPatchScheduleResource, please refer to the document of RedisPatchScheduleResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string name = "cache1";
RedisPatchScheduleDefaultName defaultName = RedisPatchScheduleDefaultName.Default;
ResourceIdentifier redisPatchScheduleResourceId = RedisPatchScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, defaultName);
RedisPatchScheduleResource redisPatchSchedule = client.GetRedisPatchScheduleResource(redisPatchScheduleResourceId);

// invoke the operation
await redisPatchSchedule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
