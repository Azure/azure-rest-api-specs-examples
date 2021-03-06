const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing registered ASN with the specified name under the given subscription, resource group and peering.
 *
 * @summary Gets an existing registered ASN with the specified name under the given subscription, resource group and peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/GetRegisteredAsn.json
 */
async function getARegisteredAsnAssociatedWithThePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const registeredAsnName = "registeredAsnName0";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.registeredAsns.get(resourceGroupName, peeringName, registeredAsnName);
  console.log(result);
}

getARegisteredAsnAssociatedWithThePeering().catch(console.error);
