const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified network watcher resource.
 *
 * @summary Deletes the specified network watcher resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherDelete.json
 */
async function deleteNetworkWatcher() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginDeleteAndWait(
    resourceGroupName,
    networkWatcherName
  );
  console.log(result);
}

deleteNetworkWatcher().catch(console.error);
