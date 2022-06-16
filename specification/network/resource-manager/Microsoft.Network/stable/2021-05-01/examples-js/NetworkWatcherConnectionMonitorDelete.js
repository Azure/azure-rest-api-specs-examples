const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified connection monitor.
 *
 * @summary Deletes the specified connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherConnectionMonitorDelete.json
 */
async function deleteConnectionMonitor() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.beginDeleteAndWait(
    resourceGroupName,
    networkWatcherName,
    connectionMonitorName
  );
  console.log(result);
}

deleteConnectionMonitor().catch(console.error);
