using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Get.json
// this example is just showing the usage of "ScalingPlanPersonalSchedules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScalingPlanPersonalScheduleResource created on azure
// for more information of creating ScalingPlanPersonalScheduleResource, please refer to the document of ScalingPlanPersonalScheduleResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string scalingPlanName = "PersonalScalingPlan1";
string scalingPlanScheduleName = "PersonalScalingPlanSchedule";
ResourceIdentifier scalingPlanPersonalScheduleResourceId = ScalingPlanPersonalScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scalingPlanName, scalingPlanScheduleName);
ScalingPlanPersonalScheduleResource scalingPlanPersonalSchedule = client.GetScalingPlanPersonalScheduleResource(scalingPlanPersonalScheduleResourceId);

// invoke the operation
ScalingPlanPersonalScheduleResource result = await scalingPlanPersonalSchedule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScalingPlanPersonalScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
