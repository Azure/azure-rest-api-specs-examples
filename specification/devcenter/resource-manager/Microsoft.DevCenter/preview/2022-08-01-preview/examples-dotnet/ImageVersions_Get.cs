using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ImageVersions_Get.json
// this example is just showing the usage of "ImageVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImageVersionResource created on azure
// for more information of creating ImageVersionResource, please refer to the document of ImageVersionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string galleryName = "DefaultDevGallery";
string imageName = "Win11";
string versionName = "{versionName}";
ResourceIdentifier imageVersionResourceId = ImageVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, galleryName, imageName, versionName);
ImageVersionResource imageVersion = client.GetImageVersionResource(imageVersionResourceId);

// invoke the operation
ImageVersionResource result = await imageVersion.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImageVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
