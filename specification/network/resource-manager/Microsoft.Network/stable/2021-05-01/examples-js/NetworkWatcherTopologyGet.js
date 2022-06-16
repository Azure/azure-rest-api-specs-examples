const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the current network topology by resource group.
 *
 * @summary Gets the current network topology by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherTopologyGet.json
 */
async function getTopology() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = { targetResourceGroupName: "rg2" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.getTopology(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

getTopology().catch(console.error);
