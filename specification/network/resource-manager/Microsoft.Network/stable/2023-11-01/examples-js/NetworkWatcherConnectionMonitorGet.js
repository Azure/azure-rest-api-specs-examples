const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a connection monitor by name.
 *
 * @summary Gets a connection monitor by name.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/NetworkWatcherConnectionMonitorGet.json
 */
async function getConnectionMonitor() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.get(
    resourceGroupName,
    networkWatcherName,
    connectionMonitorName,
  );
  console.log(result);
}
