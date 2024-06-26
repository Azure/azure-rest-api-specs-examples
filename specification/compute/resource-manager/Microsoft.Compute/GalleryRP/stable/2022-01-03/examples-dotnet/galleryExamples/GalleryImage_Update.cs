using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryImage_Update.json
// this example is just showing the usage of "GalleryImages_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this GalleryImageResource created on azure
// for more information of creating GalleryImageResource, please refer to the document of GalleryImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryImageName = "myGalleryImageName";
ResourceIdentifier galleryImageResourceId = GalleryImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryImageName);
GalleryImageResource galleryImage = client.GetGalleryImageResource(galleryImageResourceId);

// invoke the operation
GalleryImagePatch patch = new GalleryImagePatch()
{
    OSType = SupportedOperatingSystemType.Windows,
    OSState = OperatingSystemStateType.Generalized,
    HyperVGeneration = HyperVGeneration.V1,
    Identifier = new GalleryImageIdentifier("myPublisherName", "myOfferName", "mySkuName"),
};
ArmOperation<GalleryImageResource> lro = await galleryImage.UpdateAsync(WaitUntil.Completed, patch);
GalleryImageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GalleryImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
