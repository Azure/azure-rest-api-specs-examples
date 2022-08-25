using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MinimumSet_Gen.json
// this example is just showing the usage of "VirtualMachineImages_ListByEdgeZone" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "5ece5940-d962-4dad-a98f-ca9ac0f021a5";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);
            
// invoke the operation and iterate over the result
AzureLocation location = new AzureLocation("WestUS");
string edgeZone = "microsoftlosangeles1";
await foreach (VirtualMachineImageBase item in subscriptionResource.GetVirtualMachineImagesByEdgeZoneAsync(location, edgeZone))
{
    Console.WriteLine($"Succeeded: {item}");
}
            
Console.WriteLine($"Succeeded");
