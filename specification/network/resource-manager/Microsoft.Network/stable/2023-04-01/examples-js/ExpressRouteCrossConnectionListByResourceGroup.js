const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all the ExpressRouteCrossConnections in a resource group.
 *
 * @summary Retrieves all the ExpressRouteCrossConnections in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/ExpressRouteCrossConnectionListByResourceGroup.json
 */
async function expressRouteCrossConnectionListByResourceGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["NETWORK_RESOURCE_GROUP"] || "CrossConnection-SiliconValley";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteCrossConnections.listByResourceGroup(
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
