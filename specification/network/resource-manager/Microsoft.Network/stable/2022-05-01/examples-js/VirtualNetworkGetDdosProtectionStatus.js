const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Ddos Protection Status of all IP Addresses under the Virtual Network
 *
 * @summary Gets the Ddos Protection Status of all IP Addresses under the Virtual Network
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGetDdosProtectionStatus.json
 */
async function getDdosProtectionStatusOfAVirtualNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "test-vnet";
  const top = 75;
  const options = {
    top,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworks.beginListDdosProtectionStatusAndWait(
    resourceGroupName,
    virtualNetworkName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getDdosProtectionStatusOfAVirtualNetwork().catch(console.error);
