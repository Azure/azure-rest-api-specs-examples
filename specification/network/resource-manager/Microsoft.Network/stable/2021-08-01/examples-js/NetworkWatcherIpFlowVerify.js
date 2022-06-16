const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Verify IP flow from the specified VM to a location given the currently configured NSG rules.
 *
 * @summary Verify IP flow from the specified VM to a location given the currently configured NSG rules.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherIpFlowVerify.json
 */
async function ipFlowVerify() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const parameters = {
    direction: "Outbound",
    localIPAddress: "10.2.0.4",
    localPort: "80",
    remoteIPAddress: "121.10.1.1",
    remotePort: "80",
    targetResourceId:
      "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1",
    protocol: "TCP",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkWatchers.beginVerifyIPFlowAndWait(
    resourceGroupName,
    networkWatcherName,
    parameters
  );
  console.log(result);
}

ipFlowVerify().catch(console.error);
