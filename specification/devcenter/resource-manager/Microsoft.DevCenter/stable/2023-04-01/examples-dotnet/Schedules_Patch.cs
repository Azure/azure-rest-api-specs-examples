using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Schedules_Patch.json
// this example is just showing the usage of "Schedules_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterScheduleResource created on azure
// for more information of creating DevCenterScheduleResource, please refer to the document of DevCenterScheduleResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string projectName = "TestProject";
string poolName = "DevPool";
string scheduleName = "autoShutdown";
ResourceIdentifier devCenterScheduleResourceId = DevCenterScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, poolName, scheduleName);
DevCenterScheduleResource devCenterSchedule = client.GetDevCenterScheduleResource(devCenterScheduleResourceId);

// invoke the operation
DevCenterSchedulePatch patch = new DevCenterSchedulePatch
{
    Time = "18:00",
};
ArmOperation<DevCenterScheduleResource> lro = await devCenterSchedule.UpdateAsync(WaitUntil.Completed, patch);
DevCenterScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevCenterScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
