const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the usage for a subscription
 *
 * @summary Get the usage for a subscription
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupUsage.json
 */
async function containerUsage() {
  const subscriptionId = "subid";
  const location = "westcentralus";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listUsage(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

containerUsage().catch(console.error);
