const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new registered ASN with the specified name under the given subscription, resource group and peering.
 *
 * @summary Creates a new registered ASN with the specified name under the given subscription, resource group and peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreateRegisteredAsn.json
 */
async function createOrUpdateARegisteredAsnForThePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const registeredAsnName = "registeredAsnName";
  const registeredAsn = { asn: 65000 };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.registeredAsns.createOrUpdate(
    resourceGroupName,
    peeringName,
    registeredAsnName,
    registeredAsn
  );
  console.log(result);
}

createOrUpdateARegisteredAsnForThePeering().catch(console.error);
