const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified virtual network tap.
 *
 * @summary Deletes the specified virtual network tap.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkTapDelete.json
 */
async function deleteVirtualNetworkTapResource() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const tapName = "test-vtap";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkTaps.beginDeleteAndWait(resourceGroupName, tapName);
  console.log(result);
}

deleteVirtualNetworkTapResource().catch(console.error);
