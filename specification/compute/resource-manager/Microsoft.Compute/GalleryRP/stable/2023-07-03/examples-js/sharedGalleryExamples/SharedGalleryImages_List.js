const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List shared gallery images by subscription id or tenant id.
 *
 * @summary List shared gallery images by subscription id or tenant id.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/sharedGalleryExamples/SharedGalleryImages_List.json
 */
async function listSharedGalleryImages() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
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
