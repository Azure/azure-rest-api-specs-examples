using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/sharedGalleryExamples/SharedGalleryImage_Get.json
// this example is just showing the usage of "SharedGalleryImages_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SharedGalleryResource created on azure
// for more information of creating SharedGalleryResource, please refer to the document of SharedGalleryResource
string subscriptionId = "{subscription-id}";
AzureLocation location = new AzureLocation("myLocation");
string galleryUniqueName = "galleryUniqueName";
ResourceIdentifier sharedGalleryResourceId = SharedGalleryResource.CreateResourceIdentifier(subscriptionId, location, galleryUniqueName);
SharedGalleryResource sharedGallery = client.GetSharedGalleryResource(sharedGalleryResourceId);

// get the collection of this SharedGalleryImageResource
SharedGalleryImageCollection collection = sharedGallery.GetSharedGalleryImages();

// invoke the operation
string galleryImageName = "myGalleryImageName";
bool result = await collection.ExistsAsync(galleryImageName);

Console.WriteLine($"Succeeded: {result}");
