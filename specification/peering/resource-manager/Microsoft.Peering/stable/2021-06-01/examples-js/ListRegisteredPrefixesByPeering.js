const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all registered prefixes under the given subscription, resource group and peering.
 *
 * @summary Lists all registered prefixes under the given subscription, resource group and peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListRegisteredPrefixesByPeering.json
 */
async function listAllTheRegisteredPrefixesAssociatedWithThePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.registeredPrefixes.listByPeering(resourceGroupName, peeringName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllTheRegisteredPrefixesAssociatedWithThePeering().catch(console.error);
