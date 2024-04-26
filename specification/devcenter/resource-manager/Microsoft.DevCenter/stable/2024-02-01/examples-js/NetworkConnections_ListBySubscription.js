const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists network connections in a subscription
 *
 * @summary Lists network connections in a subscription
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/NetworkConnections_ListBySubscription.json
 */
async function networkConnectionsListBySubscription() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkConnections.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
