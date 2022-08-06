const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts the specified connection monitor.
 *
 * @summary Starts the specified connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkWatcherConnectionMonitorStart.json
 */
async function startConnectionMonitor() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.beginStartAndWait(
    resourceGroupName,
    networkWatcherName,
    connectionMonitorName
  );
  console.log(result);
}

startConnectionMonitor().catch(console.error);
