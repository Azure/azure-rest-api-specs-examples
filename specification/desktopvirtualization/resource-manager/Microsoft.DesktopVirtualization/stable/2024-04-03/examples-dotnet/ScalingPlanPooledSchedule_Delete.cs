using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/ScalingPlanPooledSchedule_Delete.json
// this example is just showing the usage of "ScalingPlanPooledSchedules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScalingPlanPooledScheduleResource created on azure
// for more information of creating ScalingPlanPooledScheduleResource, please refer to the document of ScalingPlanPooledScheduleResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string scalingPlanName = "scalingPlan1";
string scalingPlanScheduleName = "scalingPlanScheduleWeekdays1";
ResourceIdentifier scalingPlanPooledScheduleResourceId = ScalingPlanPooledScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scalingPlanName, scalingPlanScheduleName);
ScalingPlanPooledScheduleResource scalingPlanPooledSchedule = client.GetScalingPlanPooledScheduleResource(scalingPlanPooledScheduleResourceId);

// invoke the operation
await scalingPlanPooledSchedule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
