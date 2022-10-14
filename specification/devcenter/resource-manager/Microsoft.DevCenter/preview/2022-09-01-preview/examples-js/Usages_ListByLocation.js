const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the current usages and limits in this location for the provided subscription.
 *
 * @summary Lists the current usages and limits in this location for the provided subscription.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Usages_ListByLocation.json
 */
async function listUsages() {
  const subscriptionId = "{subscriptionId}";
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
