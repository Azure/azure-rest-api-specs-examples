using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineImagesEdgeZone_List" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);
            
// invoke the operation and iterate over the result
AzureLocation location = new AzureLocation("aaaaaaaaaaaaaaa");
string edgeZone = "aaaaaaaaaaaaaaaaaaaaaaaaa";
string publisherName = "aaaa";
string offer = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
string skus = "aaaaaaaaaaaaaaaaaaaaaaa";
string expand = "aaaaaaaaaaaaaaaaaaaaaaaa";
int? top = 12;
string orderby = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
await foreach (VirtualMachineImageBase item in subscriptionResource.GetVirtualMachineImagesEdgeZonesAsync(location, edgeZone, publisherName, offer, skus, expand: expand, top: top, orderby: orderby))
{
    Console.WriteLine($"Succeeded: {item}");
}
            
Console.WriteLine($"Succeeded");
