const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update sharing profile of a gallery.
 *
 * @summary Update sharing profile of a gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/Gallery_ResetSharingProfile.json
 */
async function resetSharingProfileOfAGallery() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const galleryName = "myGalleryName";
  const sharingUpdate = { operationType: "Reset" };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.gallerySharingProfile.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    sharingUpdate,
  );
  console.log(result);
}
