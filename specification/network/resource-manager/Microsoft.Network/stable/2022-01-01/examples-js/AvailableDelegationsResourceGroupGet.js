const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all of the available subnet delegations for this resource group in this region.
 *
 * @summary Gets all of the available subnet delegations for this resource group in this region.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/AvailableDelegationsResourceGroupGet.json
 */
async function getAvailableDelegationsInTheResourceGroup() {
  const subscriptionId = "subId";
  const location = "westcentralus";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableResourceGroupDelegations.list(
    location,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAvailableDelegationsInTheResourceGroup().catch(console.error);
