const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update tags of the specified connection monitor.
 *
 * @summary Update tags of the specified connection monitor.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherConnectionMonitorUpdateTags.json
 */
async function updateConnectionMonitorTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const connectionMonitorName = "cm1";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectionMonitors.updateTags(
    resourceGroupName,
    networkWatcherName,
    connectionMonitorName,
    parameters
  );
  console.log(result);
}

updateConnectionMonitorTags().catch(console.error);
