using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryApplication_Get.json
// this example is just showing the usage of "GalleryApplications_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryApplicationResource created on azure
// for more information of creating GalleryApplicationResource, please refer to the document of GalleryApplicationResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string galleryApplicationName = "myGalleryApplicationName";
ResourceIdentifier galleryApplicationResourceId = GalleryApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryApplicationName);
GalleryApplicationResource galleryApplication = client.GetGalleryApplicationResource(galleryApplicationResourceId);

// invoke the operation
GalleryApplicationResource result = await galleryApplication.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GalleryApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
