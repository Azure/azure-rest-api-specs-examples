const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateAPrefixForThePeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const prefixName = "peeringServicePrefixName";
  const peeringServicePrefix = {
    peeringServicePrefixKey: "00000000-0000-0000-0000-000000000000",
    prefix: "192.168.1.0/24",
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.prefixes.createOrUpdate(
    resourceGroupName,
    peeringServiceName,
    prefixName,
    peeringServicePrefix
  );
  console.log(result);
}

createOrUpdateAPrefixForThePeeringService().catch(console.error);
