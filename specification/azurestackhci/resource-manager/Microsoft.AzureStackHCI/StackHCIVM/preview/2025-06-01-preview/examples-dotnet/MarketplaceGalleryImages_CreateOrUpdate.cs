using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/MarketplaceGalleryImages_CreateOrUpdate.json
// this example is just showing the usage of "MarketplaceGalleryImage_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HciVmMarketplaceGalleryImageResource
HciVmMarketplaceGalleryImageCollection collection = resourceGroupResource.GetHciVmMarketplaceGalleryImages();

// invoke the operation
string marketplaceGalleryImageName = "test-marketplace-gallery-image";
HciVmMarketplaceGalleryImageData data = new HciVmMarketplaceGalleryImageData(new AzureLocation("West US2"))
{
    Properties = new HciVmMarketplaceGalleryImageProperties(HciVmOSType.Windows)
    {
        ContainerId = new ResourceIdentifier("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
        CloudInitDataSource = CloudInitDataSource.Azure,
        HyperVGeneration = HciVmHyperVGeneration.V1,
        Identifier = new HciVmGalleryImageIdentifier("myPublisherName", "myOfferName", "mySkuName"),
        Version = new HciVmGalleryImageVersion
        {
            Name = "1.0.0",
        },
    },
    ExtendedLocation = new HciVmExtendedLocation
    {
        Name = "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
        Type = HciVmExtendedLocationType.CustomLocation,
    },
};
ArmOperation<HciVmMarketplaceGalleryImageResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, marketplaceGalleryImageName, data);
HciVmMarketplaceGalleryImageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciVmMarketplaceGalleryImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
