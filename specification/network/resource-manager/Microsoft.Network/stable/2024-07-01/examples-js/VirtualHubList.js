const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all the VirtualHubs in a subscription.
 *
 * @summary Lists all the VirtualHubs in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualHubList.json
 */
async function virtualHubList() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualHubs.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
