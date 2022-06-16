const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Bastion Hosts in a resource group.
 *
 * @summary Lists all Bastion Hosts in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/BastionHostListByResourceGroup.json
 */
async function listAllBastionHostsForAGivenResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bastionHosts.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllBastionHostsForAGivenResourceGroup().catch(console.error);
