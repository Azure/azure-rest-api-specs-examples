using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/ScalingPlan_Delete.json
// this example is just showing the usage of "ScalingPlans_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
await scalingPlan.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
