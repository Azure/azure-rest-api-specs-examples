using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_CreateDailyShutdownPoolSchedule.json
// this example is just showing the usage of "Schedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PoolResource created on azure
// for more information of creating PoolResource, please refer to the document of PoolResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string projectName = "DevProject";
string poolName = "DevPool";
ResourceIdentifier poolResourceId = PoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, poolName);
PoolResource pool = client.GetPoolResource(poolResourceId);

// get the collection of this ScheduleResource
ScheduleCollection collection = pool.GetSchedules();

// invoke the operation
string scheduleName = "autoShutdown";
ScheduleData data = new ScheduleData()
{
    TypePropertiesType = ScheduledType.StopDevBox,
    Frequency = ScheduledFrequency.Daily,
    Time = "17:30",
    TimeZone = "America/Los_Angeles",
    State = EnableStatus.Enabled,
};
ArmOperation<ScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scheduleName, data);
ScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
