using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_Patch.json
// this example is just showing the usage of "Schedules_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScheduleResource created on azure
// for more information of creating ScheduleResource, please refer to the document of ScheduleResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string projectName = "TestProject";
string poolName = "DevPool";
string scheduleName = "autoShutdown";
ResourceIdentifier scheduleResourceId = ScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, poolName, scheduleName);
ScheduleResource schedule = client.GetScheduleResource(scheduleResourceId);

// invoke the operation
SchedulePatch patch = new SchedulePatch()
{
    Time = "18:00",
};
await schedule.UpdateAsync(WaitUntil.Completed, patch);

Console.WriteLine($"Succeeded");
