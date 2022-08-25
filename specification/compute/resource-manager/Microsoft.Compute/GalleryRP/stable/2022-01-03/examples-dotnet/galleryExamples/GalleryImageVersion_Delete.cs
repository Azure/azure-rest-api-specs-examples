using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryImageVersion_Delete.json
// this example is just showing the usage of "GalleryImageVersions_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

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

Console.WriteLine($"Succeeded");
