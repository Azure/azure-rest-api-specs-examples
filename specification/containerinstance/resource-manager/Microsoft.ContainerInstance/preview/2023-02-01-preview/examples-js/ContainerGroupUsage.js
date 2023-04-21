const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the usage for a subscription
 *
 * @summary Get the usage for a subscription
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2023-02-01-preview/examples/ContainerGroupUsage.json
 */
async function containerUsage() {
  const subscriptionId = process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "subid";
  const location = "westcentralus";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listUsage(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
