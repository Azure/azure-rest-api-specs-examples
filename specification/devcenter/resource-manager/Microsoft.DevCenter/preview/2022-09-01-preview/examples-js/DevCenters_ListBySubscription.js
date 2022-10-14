const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all devcenters in a subscription.
 *
 * @summary Lists all devcenters in a subscription.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/DevCenters_ListBySubscription.json
 */
async function devCentersListBySubscription() {
  const subscriptionId = "{subscriptionId}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.devCenters.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

devCentersListBySubscription().catch(console.error);
