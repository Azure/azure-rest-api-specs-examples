using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryImageVersion_Delete.json
// this example is just showing the usage of "GalleryImageVersions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryImageVersionResource created on azure
// for more information of creating GalleryImageVersionResource, please refer to the document of GalleryImageVersionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryImageName = "myGalleryImageName";
string galleryImageVersionName = "1.0.0";
ResourceIdentifier galleryImageVersionResourceId = GalleryImageVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName);
GalleryImageVersionResource galleryImageVersion = client.GetGalleryImageVersionResource(galleryImageVersionResourceId);

// invoke the operation
await galleryImageVersion.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
