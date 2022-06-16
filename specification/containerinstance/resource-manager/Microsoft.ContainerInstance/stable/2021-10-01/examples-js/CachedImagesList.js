const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the list of cached images on specific OS type for a subscription in a region.
 *
 * @summary Get the list of cached images on specific OS type for a subscription in a region.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/CachedImagesList.json
 */
async function cachedImages() {
  const subscriptionId = "subid";
  const location = "westcentralus";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listCachedImages(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

cachedImages().catch(console.error);
