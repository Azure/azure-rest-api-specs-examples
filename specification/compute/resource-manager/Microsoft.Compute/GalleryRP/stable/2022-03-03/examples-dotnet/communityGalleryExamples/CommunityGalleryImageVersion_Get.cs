using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/communityGalleryExamples/CommunityGalleryImageVersion_Get.json
// this example is just showing the usage of "CommunityGalleryImageVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommunityGalleryImageResource created on azure
// for more information of creating CommunityGalleryImageResource, please refer to the document of CommunityGalleryImageResource
string subscriptionId = "{subscription-id}";
AzureLocation location = new AzureLocation("myLocation");
string publicGalleryName = "publicGalleryName";
string galleryImageName = "myGalleryImageName";
ResourceIdentifier communityGalleryImageResourceId = CommunityGalleryImageResource.CreateResourceIdentifier(subscriptionId, location, publicGalleryName, galleryImageName);
CommunityGalleryImageResource communityGalleryImage = client.GetCommunityGalleryImageResource(communityGalleryImageResourceId);

// get the collection of this CommunityGalleryImageVersionResource
CommunityGalleryImageVersionCollection collection = communityGalleryImage.GetCommunityGalleryImageVersions();

// invoke the operation
string galleryImageVersionName = "myGalleryImageVersionName";
bool result = await collection.ExistsAsync(galleryImageVersionName);

Console.WriteLine($"Succeeded: {result}");
