const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a marketplace gallery image. Please note some properties can be set only during marketplace gallery image creation.
 *
 * @summary The operation to create or update a marketplace gallery image. Please note some properties can be set only during marketplace gallery image creation.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutMarketplaceGalleryImage.json
 */
async function putMarketplaceGalleryImage() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const marketplaceGalleryImageName = "test-marketplace-gallery-image";
  const marketplaceGalleryImages = {
    cloudInitDataSource: "Azure",
    containerId:
      "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container",
    extendedLocation: {
      name: "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
      type: "CustomLocation",
    },
    hyperVGeneration: "V1",
    identifier: {
      offer: "myOfferName",
      publisher: "myPublisherName",
      sku: "mySkuName",
    },
    location: "West US2",
    osType: "Windows",
    version: { name: "1.0.0" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.marketplaceGalleryImagesOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    marketplaceGalleryImageName,
    marketplaceGalleryImages
  );
  console.log(result);
}
