using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/GalleryImageVersion_Get_WithSnapshotsAsSource.json
// this example is just showing the usage of "GalleryImageVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryImageResource created on azure
// for more information of creating GalleryImageResource, please refer to the document of GalleryImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryImageName = "myGalleryImageName";
ResourceIdentifier galleryImageResourceId = GalleryImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryImageName);
GalleryImageResource galleryImage = client.GetGalleryImageResource(galleryImageResourceId);

// get the collection of this GalleryImageVersionResource
GalleryImageVersionCollection collection = galleryImage.GetGalleryImageVersions();

// invoke the operation
string galleryImageVersionName = "1.0.0";
bool result = await collection.ExistsAsync(galleryImageVersionName);

Console.WriteLine($"Succeeded: {result}");
