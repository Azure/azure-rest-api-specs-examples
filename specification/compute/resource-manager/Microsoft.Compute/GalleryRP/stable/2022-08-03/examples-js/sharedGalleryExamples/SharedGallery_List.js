const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List shared galleries by subscription id or tenant id.
 *
 * @summary List shared galleries by subscription id or tenant id.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/sharedGalleryExamples/SharedGallery_List.json
 */
async function listSharedGalleries() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const location = "myLocation";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sharedGalleries.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
