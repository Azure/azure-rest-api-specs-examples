using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/Gallery_Delete.json
// this example is just showing the usage of "Galleries_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryResource created on azure
// for more information of creating GalleryResource, please refer to the document of GalleryResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
ResourceIdentifier galleryResourceId = GalleryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName);
GalleryResource gallery = client.GetGalleryResource(galleryResourceId);

// invoke the operation
await gallery.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
