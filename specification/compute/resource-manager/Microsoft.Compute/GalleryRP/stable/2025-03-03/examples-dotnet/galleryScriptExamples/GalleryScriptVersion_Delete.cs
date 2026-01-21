using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryScriptExamples/GalleryScriptVersion_Delete.json
// this example is just showing the usage of "GalleryScriptVersions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryScriptVersionResource created on azure
// for more information of creating GalleryScriptVersionResource, please refer to the document of GalleryScriptVersionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroupName";
string galleryName = "myGalleryName";
string galleryScriptName = "myGalleryScriptName";
string galleryScriptVersionName = "1.0.0";
ResourceIdentifier galleryScriptVersionResourceId = GalleryScriptVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, galleryScriptName, galleryScriptVersionName);
GalleryScriptVersionResource galleryScriptVersion = client.GetGalleryScriptVersionResource(galleryScriptVersionResourceId);

// invoke the operation
await galleryScriptVersion.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
