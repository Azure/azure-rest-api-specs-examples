const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAPeerAsn() {
  const subscriptionId = "subId";
  const peerAsnName = "peerAsnName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerAsns.get(peerAsnName);
  console.log(result);
}

getAPeerAsn().catch(console.error);
