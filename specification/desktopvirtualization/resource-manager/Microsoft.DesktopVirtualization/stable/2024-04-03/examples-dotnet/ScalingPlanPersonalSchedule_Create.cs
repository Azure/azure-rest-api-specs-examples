using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/ScalingPlanPersonalSchedule_Create.json
// this example is just showing the usage of "ScalingPlanPersonalSchedules_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScalingPlanResource created on azure
// for more information of creating ScalingPlanResource, please refer to the document of ScalingPlanResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string scalingPlanName = "scalingPlan1";
ResourceIdentifier scalingPlanResourceId = ScalingPlanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scalingPlanName);
ScalingPlanResource scalingPlan = client.GetScalingPlanResource(scalingPlanResourceId);

// get the collection of this ScalingPlanPersonalScheduleResource
ScalingPlanPersonalScheduleCollection collection = scalingPlan.GetScalingPlanPersonalSchedules();

// invoke the operation
string scalingPlanScheduleName = "scalingPlanScheduleWeekdays1";
ScalingPlanPersonalScheduleData data = new ScalingPlanPersonalScheduleData
{
    DaysOfWeek = { DesktopVirtualizationDayOfWeek.Monday, DesktopVirtualizationDayOfWeek.Tuesday, DesktopVirtualizationDayOfWeek.Wednesday, DesktopVirtualizationDayOfWeek.Thursday, DesktopVirtualizationDayOfWeek.Friday },
    RampUpStartTime = new ScalingActionTime(6, 0),
    RampUpAutoStartHosts = StartupBehavior.All,
    RampUpStartVmOnConnect = SetStartVmOnConnect.Enable,
    RampUpActionOnDisconnect = SessionHandlingOperation.None,
    RampUpMinutesToWaitOnDisconnect = 10,
    RampUpActionOnLogoff = SessionHandlingOperation.None,
    RampUpMinutesToWaitOnLogoff = 10,
    PeakStartTime = new ScalingActionTime(8, 0),
    PeakStartVmOnConnect = SetStartVmOnConnect.Enable,
    PeakActionOnDisconnect = SessionHandlingOperation.None,
    PeakMinutesToWaitOnDisconnect = 10,
    PeakActionOnLogoff = SessionHandlingOperation.Deallocate,
    PeakMinutesToWaitOnLogoff = 10,
    RampDownStartTime = new ScalingActionTime(18, 0),
    RampDownStartVmOnConnect = SetStartVmOnConnect.Enable,
    RampDownActionOnDisconnect = SessionHandlingOperation.None,
    RampDownMinutesToWaitOnDisconnect = 10,
    RampDownActionOnLogoff = SessionHandlingOperation.Deallocate,
    RampDownMinutesToWaitOnLogoff = 10,
    OffPeakStartTime = new ScalingActionTime(20, 0),
    OffPeakStartVmOnConnect = SetStartVmOnConnect.Enable,
    OffPeakActionOnDisconnect = SessionHandlingOperation.None,
    OffPeakMinutesToWaitOnDisconnect = 10,
    OffPeakActionOnLogoff = SessionHandlingOperation.Deallocate,
    OffPeakMinutesToWaitOnLogoff = 10,
};
ArmOperation<ScalingPlanPersonalScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scalingPlanScheduleName, data);
ScalingPlanPersonalScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScalingPlanPersonalScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
