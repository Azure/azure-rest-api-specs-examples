using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_Retarget.json
// this example is just showing the usage of "GlobalSchedules_Retarget" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabGlobalScheduleResource created on azure
// for more information of creating DevTestLabGlobalScheduleResource, please refer to the document of DevTestLabGlobalScheduleResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string name = "{scheduleName}";
ResourceIdentifier devTestLabGlobalScheduleResourceId = DevTestLabGlobalScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DevTestLabGlobalScheduleResource devTestLabGlobalSchedule = client.GetDevTestLabGlobalScheduleResource(devTestLabGlobalScheduleResourceId);

// invoke the operation
DevTestLabGlobalScheduleRetargetContent content = new DevTestLabGlobalScheduleRetargetContent
{
    CurrentResourceId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{targetLab}"),
    TargetResourceId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{currentLab}"),
};
await devTestLabGlobalSchedule.RetargetAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
