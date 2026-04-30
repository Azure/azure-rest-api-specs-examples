const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a gallery image definition.
 *
 * @summary update a gallery image definition.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImage_Update.json
 */
async function updateASimpleGalleryImage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImages.update(
    "myResourceGroup",
    "myGalleryName",
    "myGalleryImageName",
    {
      osType: "Windows",
      osState: "Generalized",
      hyperVGeneration: "V1",
      identifier: { publisher: "myPublisherName", offer: "myOfferName", sku: "mySkuName" },
    },
  );
  console.log(result);
}
