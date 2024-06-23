const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Azure Firewalls in a resource group.
 *
 * @summary Lists all Azure Firewalls in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/AzureFirewallListByResourceGroup.json
 */
async function listAllAzureFirewallsForAGivenResourceGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureFirewalls.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
