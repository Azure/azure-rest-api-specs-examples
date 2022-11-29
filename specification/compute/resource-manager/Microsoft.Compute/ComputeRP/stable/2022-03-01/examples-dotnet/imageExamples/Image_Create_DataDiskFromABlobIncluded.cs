using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/imageExamples/Image_Create_DataDiskFromABlobIncluded.json
// this example is just showing the usage of "Images_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DiskImageResource
DiskImageCollection collection = resourceGroupResource.GetDiskImages();

// invoke the operation
string imageName = "myImage";
DiskImageData data = new DiskImageData(new AzureLocation("West US"))
{
    StorageProfile = new ImageStorageProfile()
    {
        OSDisk = new ImageOSDisk(SupportedOperatingSystemType.Linux, OperatingSystemStateType.Generalized)
        {
            BlobUri = new Uri("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
        },
        DataDisks =
        {
        new ImageDataDisk(1)
        {
        BlobUri = new Uri("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
        }
        },
        ZoneResilient = false,
    },
};
ArmOperation<DiskImageResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, imageName, data);
DiskImageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiskImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
