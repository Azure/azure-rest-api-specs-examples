const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about a gallery image version.
 *
 * @summary retrieves information about a gallery image version.
 * x-ms-original-file: 2025-03-03/galleryExamples/GalleryImageVersion_Get_WithValidationProfileAndReplicationStatus.json
 */
async function getAGalleryImageVersionWithValidationProfileAndReplicationStatus() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImageVersions.get(
    "myResourceGroup",
    "myGalleryName",
    "myGalleryImageName",
    "1.0.0",
    { expand: "ValidationProfile,ReplicationStatus" },
  );
  console.log(result);
}
