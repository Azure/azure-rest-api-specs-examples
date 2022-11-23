const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists network connections in a subscription
 *
 * @summary Lists network connections in a subscription
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/NetworkConnections_ListBySubscription.json
 */
async function networkConnectionsListBySubscription() {
  const subscriptionId = "{subscriptionId}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkConnections.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

networkConnectionsListBySubscription().catch(console.error);
