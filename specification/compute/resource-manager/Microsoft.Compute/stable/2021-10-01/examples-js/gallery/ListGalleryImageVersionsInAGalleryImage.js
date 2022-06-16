const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List gallery image versions in a gallery image definition.
 *
 * @summary List gallery image versions in a gallery image definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/ListGalleryImageVersionsInAGalleryImage.json
 */
async function listGalleryImageVersionsInAGalleryImageDefinition() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleryImageVersions.listByGalleryImage(
    resourceGroupName,
    galleryName,
    galleryImageName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listGalleryImageVersionsInAGalleryImageDefinition().catch(console.error);
