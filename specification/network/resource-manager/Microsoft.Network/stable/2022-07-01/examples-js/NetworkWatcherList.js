const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all network watchers by resource group.
 *
 * @summary Gets all network watchers by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherList.json
 */
async function listNetworkWatchers() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkWatchers.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listNetworkWatchers().catch(console.error);
