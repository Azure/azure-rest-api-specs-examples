using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/ScalingPlanPooledSchedule_List.json
// this example is just showing the usage of "ScalingPlanPooledSchedules_List" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ScalingPlanPooledScheduleResource
ScalingPlanPooledScheduleCollection collection = scalingPlan.GetScalingPlanPooledSchedules();

// invoke the operation and iterate over the result
int? pageSize = 10;
bool? isDescending = true;
int? initialSkip = 0;
await foreach (ScalingPlanPooledScheduleResource item in collection.GetAllAsync(pageSize: pageSize, isDescending: isDescending, initialSkip: initialSkip))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ScalingPlanPooledScheduleData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
