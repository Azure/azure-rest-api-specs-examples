using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LabServices;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/saveImageVirtualMachine.json
// this example is just showing the usage of "LabPlans_SaveImage" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LabPlanResource created on azure
// for more information of creating LabPlanResource, please refer to the document of LabPlanResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string labPlanName = "testlabplan";
ResourceIdentifier labPlanResourceId = LabPlanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labPlanName);
LabPlanResource labPlan = client.GetLabPlanResource(labPlanResourceId);

// invoke the operation
LabVirtualMachineImageContent content = new LabVirtualMachineImageContent()
{
    Name = "Test Image",
    LabVirtualMachineId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labs/testlab/virtualMachines/template"),
};
await labPlan.SaveImageAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
