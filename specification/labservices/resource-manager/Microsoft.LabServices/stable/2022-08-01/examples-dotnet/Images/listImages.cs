using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.LabServices;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Images/listImages.json
// this example is just showing the usage of "Images_ListByLabPlan" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this LabVirtualMachineImageResource
LabVirtualMachineImageCollection collection = labPlan.GetLabVirtualMachineImages();

// invoke the operation and iterate over the result
await foreach (LabVirtualMachineImageResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    LabVirtualMachineImageData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
