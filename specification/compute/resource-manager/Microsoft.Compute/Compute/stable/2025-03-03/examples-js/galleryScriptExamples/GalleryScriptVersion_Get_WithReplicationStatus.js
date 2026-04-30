const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to custom ArmResourceRead operation template with CloudError as Error
 *
 * @summary custom ArmResourceRead operation template with CloudError as Error
 * x-ms-original-file: 2025-03-03/galleryScriptExamples/GalleryScriptVersion_Get_WithReplicationStatus.json
 */
async function getAGalleryScriptVersionWithReplicationStatus() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryScriptVersions.get(
    "myResourceGroupName",
    "myGalleryName",
    "myGalleryScriptName",
    "1.0.0",
  );
  console.log(result);
}
