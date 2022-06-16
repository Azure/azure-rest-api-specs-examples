const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAPrefixAssociatedWithThePeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const prefixName = "peeringServicePrefixName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.prefixes.delete(resourceGroupName, peeringServiceName, prefixName);
  console.log(result);
}

deleteAPrefixAssociatedWithThePeeringService().catch(console.error);
