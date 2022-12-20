const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Azure Firewalls in a resource group.
 *
 * @summary Lists all Azure Firewalls in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/AzureFirewallListByResourceGroup.json
 */
async function listAllAzureFirewallsForAGivenResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureFirewalls.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllAzureFirewallsForAGivenResourceGroup().catch(console.error);
