const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a gallery image definition.
 *
 * @summary Update a gallery image definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryImage_UpdateFeatures.json
 */
async function updateAGalleryImageFeature() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImage = {
    allowUpdateImage: true,
    features: [
      {
        name: "SecurityType",
        startsAtVersion: "2.0.0",
        value: "TrustedLaunch",
      },
    ],
    hyperVGeneration: "V2",
    identifier: {
      offer: "myOfferName",
      publisher: "myPublisherName",
      sku: "mySkuName",
    },
    osState: "Generalized",
    osType: "Windows",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImages.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryImageName,
    galleryImage,
  );
  console.log(result);
}
