const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all available service aliases for this resource group in this region.
 *
 * @summary Gets all available service aliases for this resource group in this region.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/AvailableServiceAliasesListByResourceGroup.json
 */
async function getAvailableServiceAliasesInTheResourceGroup() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const location = "westcentralus";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableServiceAliases.listByResourceGroup(
    resourceGroupName,
    location
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAvailableServiceAliasesInTheResourceGroup().catch(console.error);
