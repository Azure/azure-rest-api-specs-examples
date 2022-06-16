const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all subnets in a virtual network.
 *
 * @summary Gets all subnets in a virtual network.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/SubnetList.json
 */
async function listSubnets() {
  const subscriptionId = "subid";
  const resourceGroupName = "subnet-test";
  const virtualNetworkName = "vnetname";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.subnets.list(resourceGroupName, virtualNetworkName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSubnets().catch(console.error);
