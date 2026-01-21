using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Delete.json
// this example is just showing the usage of "GalleryInVMAccessControlProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GalleryInVmAccessControlProfileResource created on azure
// for more information of creating GalleryInVmAccessControlProfileResource, please refer to the document of GalleryInVmAccessControlProfileResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string galleryName = "myGalleryName";
string inVmAccessControlProfileName = "myInVMAccessControlProfileName";
ResourceIdentifier galleryInVmAccessControlProfileResourceId = GalleryInVmAccessControlProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, galleryName, inVmAccessControlProfileName);
GalleryInVmAccessControlProfileResource galleryInVmAccessControlProfile = client.GetGalleryInVmAccessControlProfileResource(galleryInVmAccessControlProfileResourceId);

// invoke the operation
await galleryInVmAccessControlProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
