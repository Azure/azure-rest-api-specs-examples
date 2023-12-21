const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all flow log resources for the specified Network Watcher.
 *
 * @summary Lists all flow log resources for the specified Network Watcher.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/NetworkWatcherFlowLogList.json
 */
async function listConnectionMonitors() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.flowLogs.list(resourceGroupName, networkWatcherName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
