const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified virtual network.
 *
 * @summary Deletes the specified virtual network.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkDelete.json
 */
async function deleteVirtualNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "test-vnet";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.beginDeleteAndWait(
    resourceGroupName,
    virtualNetworkName
  );
  console.log(result);
}

deleteVirtualNetwork().catch(console.error);
