const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about a Shared Image Gallery.
 *
 * @summary Retrieves information about a Shared Image Gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryWithSelectPermissions.json
 */
async function getAGalleryWithSelectPermissions() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const select = "Permissions";
  const options = { select };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.get(resourceGroupName, galleryName, options);
  console.log(result);
}

getAGalleryWithSelectPermissions().catch(console.error);
