const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List gallery Application Definitions in a gallery.
 *
 * @summary List gallery Application Definitions in a gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/ListGalleryApplicationsInAGallery.json
 */
async function listGalleryApplicationsInAGallery() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleryApplications.listByGallery(resourceGroupName, galleryName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listGalleryApplicationsInAGallery().catch(console.error);
