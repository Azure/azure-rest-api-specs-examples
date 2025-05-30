using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/communityGalleryExamples/CommunityGalleryImage_Get.json
// this example is just showing the usage of "CommunityGalleryImages_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

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
NullableResponse<CommunityGalleryImageResource> response = await collection.GetIfExistsAsync(galleryImageName);
CommunityGalleryImageResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CommunityGalleryImageData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
