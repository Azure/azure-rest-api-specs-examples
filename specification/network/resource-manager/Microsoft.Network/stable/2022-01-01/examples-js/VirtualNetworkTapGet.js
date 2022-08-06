const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified virtual network tap.
 *
 * @summary Gets information about the specified virtual network tap.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkTapGet.json
 */
async function getVirtualNetworkTap() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const tapName = "testvtap";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkTaps.get(resourceGroupName, tapName);
  console.log(result);
}

getVirtualNetworkTap().catch(console.error);
