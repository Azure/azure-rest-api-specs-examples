const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the next hop from the specified VM.
 *
 * @summary Gets the next hop from the specified VM.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/NetworkWatcherNextHopGet.json
 */
async function getNextHop() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const parameters = {
    destinationIPAddress: "10.0.0.10",
    sourceIPAddress: "10.0.0.5",
    targetNicResourceId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1",
    targetResourceId:
      "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginGetNextHopAndWait(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}
