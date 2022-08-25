using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/imageExamples/Image_Update.json
// this example is just showing the usage of "Images_Update" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this DiskImageResource created on azure
// for more information of creating DiskImageResource, please refer to the document of DiskImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string imageName = "myImage";
ResourceIdentifier diskImageResourceId = DiskImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, imageName);
DiskImageResource diskImage = client.GetDiskImageResource(diskImageResourceId);
            
// invoke the operation
DiskImagePatch patch = new DiskImagePatch()
{
    SourceVirtualMachineId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
    HyperVGeneration = HyperVGeneration.V1,
    Tags =
                {
                ["department"] = "HR",
                },
};
ArmOperation<DiskImageResource> lro = await diskImage.UpdateAsync(WaitUntil.Completed, patch);
DiskImageResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiskImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
