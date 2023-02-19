using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Galleries_Get.json
// this example is just showing the usage of "Galleries_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterResource created on azure
// for more information of creating DevCenterResource, please refer to the document of DevCenterResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
ResourceIdentifier devCenterResourceId = DevCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName);
DevCenterResource devCenter = client.GetDevCenterResource(devCenterResourceId);

// get the collection of this GalleryResource
GalleryCollection collection = devCenter.GetGalleries();

// invoke the operation
string galleryName = "{galleryName}";
bool result = await collection.ExistsAsync(galleryName);

Console.WriteLine($"Succeeded: {result}");
