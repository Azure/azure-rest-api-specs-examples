using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Delete.json
// this example is just showing the usage of "ScalingPlanPersonalSchedules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScalingPlanPersonalScheduleResource created on azure
// for more information of creating ScalingPlanPersonalScheduleResource, please refer to the document of ScalingPlanPersonalScheduleResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string scalingPlanName = "scalingPlan1";
string scalingPlanScheduleName = "scalingPlanScheduleWeekdays1";
ResourceIdentifier scalingPlanPersonalScheduleResourceId = ScalingPlanPersonalScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scalingPlanName, scalingPlanScheduleName);
ScalingPlanPersonalScheduleResource scalingPlanPersonalSchedule = client.GetScalingPlanPersonalScheduleResource(scalingPlanPersonalScheduleResourceId);

// invoke the operation
await scalingPlanPersonalSchedule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
