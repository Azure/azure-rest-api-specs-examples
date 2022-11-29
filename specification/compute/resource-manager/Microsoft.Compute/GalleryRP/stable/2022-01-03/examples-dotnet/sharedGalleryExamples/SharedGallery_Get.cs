using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/sharedGalleryExamples/SharedGallery_Get.json
// this example is just showing the usage of "SharedGalleries_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SharedGalleryResource created on azure
// for more information of creating SharedGalleryResource, please refer to the document of SharedGalleryResource
string subscriptionId = "{subscription-id}";
AzureLocation location = new AzureLocation("myLocation");
string galleryUniqueName = "galleryUniqueName";
ResourceIdentifier sharedGalleryResourceId = SharedGalleryResource.CreateResourceIdentifier(subscriptionId, location, galleryUniqueName);
SharedGalleryResource sharedGallery = client.GetSharedGalleryResource(sharedGalleryResourceId);

// invoke the operation
SharedGalleryResource result = await sharedGallery.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SharedGalleryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
