using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/MarketplaceGalleryImages_Get.json
// this example is just showing the usage of "MarketplaceGalleryImage_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciVmMarketplaceGalleryImageResource created on azure
// for more information of creating HciVmMarketplaceGalleryImageResource, please refer to the document of HciVmMarketplaceGalleryImageResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string marketplaceGalleryImageName = "test-marketplace-gallery-image";
ResourceIdentifier hciVmMarketplaceGalleryImageResourceId = HciVmMarketplaceGalleryImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, marketplaceGalleryImageName);
HciVmMarketplaceGalleryImageResource hciVmMarketplaceGalleryImage = client.GetHciVmMarketplaceGalleryImageResource(hciVmMarketplaceGalleryImageResourceId);

// invoke the operation
HciVmMarketplaceGalleryImageResource result = await hciVmMarketplaceGalleryImage.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciVmMarketplaceGalleryImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
