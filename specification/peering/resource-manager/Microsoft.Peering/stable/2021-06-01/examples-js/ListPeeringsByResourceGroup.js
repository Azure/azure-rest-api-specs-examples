const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the peerings under the given subscription and resource group.
 *
 * @summary Lists all of the peerings under the given subscription and resource group.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListPeeringsByResourceGroup.json
 */
async function listPeeringsInAResourceGroup() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peerings.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPeeringsInAResourceGroup().catch(console.error);
