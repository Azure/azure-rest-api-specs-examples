const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the current usages and limits in this location for the provided subscription.
 *
 * @summary Lists the current usages and limits in this location for the provided subscription.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Usages_ListByLocation.json
 */
async function listUsages() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listUsages().catch(console.error);
