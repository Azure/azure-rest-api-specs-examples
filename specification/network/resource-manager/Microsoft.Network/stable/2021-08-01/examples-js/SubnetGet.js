const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified subnet by virtual network and resource group.
 *
 * @summary Gets the specified subnet by virtual network and resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/SubnetGet.json
 */
async function getSubnet() {
  const subscriptionId = "subid";
  const resourceGroupName = "subnet-test";
  const virtualNetworkName = "vnetname";
  const subnetName = "subnet1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.subnets.get(resourceGroupName, virtualNetworkName, subnetName);
  console.log(result);
}

getSubnet().catch(console.error);
