const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List shared gallery images by subscription id or tenant id.
 *
 * @summary List shared gallery images by subscription id or tenant id.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/sharedGallery/ListSharedGalleryImages.json
 */
async function getAGallery() {
  const subscriptionId = "{subscription-id}";
  const location = "myLocation";
  const galleryUniqueName = "galleryUniqueName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sharedGalleryImages.list(location, galleryUniqueName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAGallery().catch(console.error);
