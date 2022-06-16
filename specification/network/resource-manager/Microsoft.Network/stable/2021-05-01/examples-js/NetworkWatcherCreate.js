const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network watcher in the specified resource group.
 *
 * @summary Creates or updates a network watcher in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherCreate.json
 */
async function createNetworkWatcher() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = { location: "eastus" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.createOrUpdate(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

createNetworkWatcher().catch(console.error);
