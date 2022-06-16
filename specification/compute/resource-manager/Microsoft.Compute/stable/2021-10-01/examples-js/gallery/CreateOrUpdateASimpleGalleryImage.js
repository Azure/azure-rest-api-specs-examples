const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a gallery image definition.
 *
 * @summary Create or update a gallery image definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImage.json
 */
async function createOrUpdateASimpleGalleryImage() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImage = {
    hyperVGeneration: "V1",
    identifier: {
      offer: "myOfferName",
      publisher: "myPublisherName",
      sku: "mySkuName",
    },
    location: "West US",
    osState: "Generalized",
    osType: "Windows",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImages.beginCreateOrUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryImageName,
    galleryImage
  );
  console.log(result);
}

createOrUpdateASimpleGalleryImage().catch(console.error);
