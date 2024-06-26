const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing registered prefix with the specified name under the given subscription, resource group and peering.
 *
 * @summary Gets an existing registered prefix with the specified name under the given subscription, resource group and peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/GetRegisteredPrefix.json
 */
async function getARegisteredPrefixAssociatedWithThePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const registeredPrefixName = "registeredPrefixName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.registeredPrefixes.get(
    resourceGroupName,
    peeringName,
    registeredPrefixName
  );
  console.log(result);
}

getARegisteredPrefixAssociatedWithThePeering().catch(console.error);
