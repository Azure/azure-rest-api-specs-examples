const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function updatePeeringServiceTags() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const tags = { tags: { tag0: "value0", tag1: "value1" } };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.update(resourceGroupName, peeringServiceName, tags);
  console.log(result);
}

updatePeeringServiceTags().catch(console.error);
