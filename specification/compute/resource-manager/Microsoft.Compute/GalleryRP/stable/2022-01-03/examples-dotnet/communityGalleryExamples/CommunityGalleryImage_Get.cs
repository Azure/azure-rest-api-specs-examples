using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/communityGalleryExamples/CommunityGalleryImage_Get.json
// this example is just showing the usage of "CommunityGalleryImages_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this CommunityGalleryResource created on azure
// for more information of creating CommunityGalleryResource, please refer to the document of CommunityGalleryResource
string subscriptionId = "{subscription-id}";
AzureLocation location = new AzureLocation("myLocation");
string publicGalleryName = "publicGalleryName";
ResourceIdentifier communityGalleryResourceId = CommunityGalleryResource.CreateResourceIdentifier(subscriptionId, location, publicGalleryName);
CommunityGalleryResource communityGallery = client.GetCommunityGalleryResource(communityGalleryResourceId);

// get the collection of this CommunityGalleryImageResource
CommunityGalleryImageCollection collection = communityGallery.GetCommunityGalleryImages();

// invoke the operation
string galleryImageName = "myGalleryImageName";
bool result = await collection.ExistsAsync(galleryImageName);

Console.WriteLine($"Succeeded: {result}");
