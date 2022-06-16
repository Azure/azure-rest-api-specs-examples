const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
