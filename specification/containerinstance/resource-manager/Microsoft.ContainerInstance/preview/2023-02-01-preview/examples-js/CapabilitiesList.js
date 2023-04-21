const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the list of CPU/memory/GPU capabilities of a region.
 *
 * @summary Get the list of CPU/memory/GPU capabilities of a region.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2023-02-01-preview/examples/CapabilitiesList.json
 */
async function getCapabilities() {
  const subscriptionId = process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listCapabilities(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
