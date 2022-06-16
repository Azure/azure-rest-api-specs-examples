const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List galleries under a subscription.
 *
 * @summary List galleries under a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/Gallery_ListBySubscription.json
 */
async function listGalleriesInASubscription() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleries.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listGalleriesInASubscription().catch(console.error);
