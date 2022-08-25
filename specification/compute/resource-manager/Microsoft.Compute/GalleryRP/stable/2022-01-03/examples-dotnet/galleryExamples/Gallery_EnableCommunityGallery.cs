using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/Gallery_EnableCommunityGallery.json
// this example is just showing the usage of "GallerySharingProfile_Update" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this GalleryResource created on azure
// for more information of creating GalleryResource, please refer to the document of GalleryResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
ResourceIdentifier galleryResourceId = GalleryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName);
GalleryResource gallery = client.GetGalleryResource(galleryResourceId);
            
// invoke the operation
SharingUpdate sharingUpdate = new SharingUpdate(SharingUpdateOperationType.EnableCommunity);
ArmOperation<SharingUpdate> lro = await gallery.UpdateSharingProfileAsync(WaitUntil.Completed, sharingUpdate);
SharingUpdate result = lro.Value;
            
Console.WriteLine($"Succeeded: {result}");
