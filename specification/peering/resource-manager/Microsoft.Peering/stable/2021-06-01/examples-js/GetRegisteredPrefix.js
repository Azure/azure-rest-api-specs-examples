const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
