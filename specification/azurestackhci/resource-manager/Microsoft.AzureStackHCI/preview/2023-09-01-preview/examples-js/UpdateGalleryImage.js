const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a gallery image.
 *
 * @summary The operation to update a gallery image.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateGalleryImage.json
 */
async function updateGalleryImage() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const galleryImageName = "test-gallery-image";
  const galleryImages = {
    tags: { additionalProperties: "sample" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.galleryImagesOperations.beginUpdateAndWait(
    resourceGroupName,
    galleryImageName,
    galleryImages
  );
  console.log(result);
}
