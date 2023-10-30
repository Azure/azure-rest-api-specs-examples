const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete a marketplace gallery image.
 *
 * @summary The operation to delete a marketplace gallery image.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteMarketplaceGalleryImage.json
 */
async function deleteMarketplaceGalleryImage() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const marketplaceGalleryImageName = "test-marketplace-gallery-image";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.marketplaceGalleryImagesOperations.beginDeleteAndWait(
    resourceGroupName,
    marketplaceGalleryImageName
  );
  console.log(result);
}
