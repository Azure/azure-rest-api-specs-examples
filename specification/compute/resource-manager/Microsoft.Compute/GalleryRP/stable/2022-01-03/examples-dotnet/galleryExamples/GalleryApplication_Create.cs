using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryApplication_Create.json
// this example is just showing the usage of "GalleryApplications_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this GalleryResource created on azure
// for more information of creating GalleryResource, please refer to the document of GalleryResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
ResourceIdentifier galleryResourceId = GalleryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName);
GalleryResource gallery = client.GetGalleryResource(galleryResourceId);

// get the collection of this GalleryApplicationResource
GalleryApplicationCollection collection = gallery.GetGalleryApplications();

// invoke the operation
string galleryApplicationName = "myGalleryApplicationName";
GalleryApplicationData data = new GalleryApplicationData(new AzureLocation("West US"))
{
    Description = "This is the gallery application description.",
    Eula = "This is the gallery application EULA.",
    PrivacyStatementUri = new Uri("myPrivacyStatementUri}"),
    ReleaseNoteUri = new Uri("myReleaseNoteUri"),
    SupportedOSType = SupportedOperatingSystemType.Windows,
};
ArmOperation<GalleryApplicationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, galleryApplicationName, data);
GalleryApplicationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GalleryApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
